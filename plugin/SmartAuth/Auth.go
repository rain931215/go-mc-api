package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Tnze/go-mc/yggdrasil"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Tokens struct {
	AccessToken string `json:"accessToken"`
	ClientToken string `json:"clientToken"`
}

var (
	syncLock     sync.Mutex
	accountsFile = viper.New()
	authProxy    *http.Client
)

func init() {
	syncLock.Lock()
	accountsFile.SetConfigName("accounts")
	accountsFile.SetConfigType("json")
	accountsFile.AddConfigPath(".")
	readAccountsFile()
	syncLock.Unlock()
}
func readAccountsFile() {
	if err := accountsFile.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			accountsFile.SetDefault("client-Token", uuid.New().String())
			accountsFile.SetDefault("accounts", map[string]string{})
			_ = accountsFile.SafeWriteConfig()
			_ = accountsFile.ReadInConfig()
		}
	}
}
func AuthWithEmail(email, password string, proxy *http.Client) (playerID, playerUUID, access string, authErr error) {
	if proxy != nil {
		authProxy = proxy
	}
	syncLock.Lock()
	defer syncLock.Unlock()
	readAccountsFile()
	clientToken := accountsFile.GetString("client-token")
	if clientToken == "" {
		clientToken = uuid.New().String()
		accountsFile.Set("client-token", clientToken)
		_ = accountsFile.WriteConfig()
	}
	if accounts := accountsFile.GetStringMapStringSlice("accounts"); accounts != nil {
		if account, ok := accounts[strings.Split(email, "@")[0]]; ok && len(account) == 3 {
			success, err := validateToken(clientToken, account[2])
			if err != nil {
				authErr = err
				return
			} else {
				if success {
					playerID = account[0]
					playerUUID = account[1]
					access = account[2]
					authErr = nil
					return
				} else {
					if clientToken, accessToken := refreshToken(clientToken, account[2]); clientToken != "" && accessToken != "" {
						if success2, err2 := validateToken(clientToken, accessToken); err2 == nil && success2 {
							playerID = account[0]
							playerUUID = account[1]
							access = accessToken
							authErr = nil
							account[2] = accessToken
							accounts[strings.Split(email, "@")[0]] = account
							accountsFile.Set("accounts", accounts)
							_ = accountsFile.WriteConfig()
							return
						}
					}
				}
			}
		}
		var err error
		playerID, playerUUID, access, err = loginNormal(email, password, clientToken)
		if err != nil {
			authErr = err
			return
		}
		authErr = nil
		account := make([]string, 3)
		account[0] = playerID
		account[1] = playerUUID
		account[2] = access
		accounts[strings.Split(email, "@")[0]] = account
		accountsFile.Set("accounts", accounts)
		_ = accountsFile.WriteConfig()
		return
	} else {
		authErr = errors.New("accounts.json valid error,please delete it and rerun the program")
		return
	}
}
func loginNormal(user, password, clientToken string) (id, uuid, access string, err error) {
	type agent struct {
		Name    string `json:"name"`
		Version int    `json:"version"`
	}
	type proof struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}
	type Profile struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	type AuthResp struct {
		Tokens
		AvailableProfiles []Profile `json:"availableProfiles"` // only present if the agent field was received
		SelectedProfile   Profile   `json:"selectedProfile"`   // only present if the agent field was received
		User              struct {
			ID         string `json:"id"`
			Properties []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			}
		} `json:"user"`
		*yggdrasil.Error
	}
	type authPayload struct {
		Agent agent `json:"agent"`
		proof
		ClientToken string `json:"clientToken,omitempty"`
		RequestUser bool   `json:"requestUser"`
	}
	pl := authPayload{
		Agent: agent{
			Name:    "Minecraft",
			Version: 1,
		},
		proof: proof{
			UserName: user,
			Password: password,
		},
		ClientToken: clientToken,
		RequestUser: true,
	}
	var ar AuthResp
	err = postAndParseResponse("authenticate", pl, &ar)
	if err != nil {
		return "", "", "", err
	}

	if ar.Error != nil {
		return "", "", "", *ar.Error
	}
	return ar.SelectedProfile.Name, ar.SelectedProfile.ID, ar.AccessToken, nil
}
func validateToken(clientToken, accessToken string) (bool, error) {
	pl := Tokens{
		AccessToken: accessToken,
		ClientToken: clientToken,
	}
	resp, err := postToAuthServer("validate", pl)
	if resp != nil && err == nil {
		switch resp.StatusCode {
		case 204:
			_ = resp.Body.Close()
			return true, nil
		case 403:
			_ = resp.Body.Close()
			return false, nil
		default:
			_ = resp.Body.Close()
			return false, errors.New("error when validating tokens")
		}
	}
	return false, nil
}
func refreshToken(clientToken, accessToken string) (string, string) {
	pl := Tokens{
		AccessToken: accessToken,
		ClientToken: clientToken,
	}
	resp, err := postToAuthServer("refresh", pl)
	if resp != nil && err == nil {
		switch resp.StatusCode {
		case 200:
			responseUnmarshal := Tokens{}
			_ = json.NewDecoder(resp.Body).Decode(&responseUnmarshal)
			_ = resp.Body.Close()
			return responseUnmarshal.ClientToken, responseUnmarshal.AccessToken
		default:
			_ = resp.Body.Close()
			return "", ""
		}
	}
	if resp != nil {
		_ = resp.Body.Close()
	}
	return "", ""
}
func postAndParseResponse(endpoint string, payload interface{}, resp interface{}) error {
	rowResp, err := postToAuthServer(endpoint, payload)
	if err != nil {
		return fmt.Errorf("request fail: %v", err)
	}
	defer rowResp.Body.Close()
	err = json.NewDecoder(rowResp.Body).Decode(resp)
	if err != nil {
		return fmt.Errorf("parse resp fail: %v", err)
	}
	return nil
}
func postToAuthServer(endPoint string, payload interface{}) (resp *http.Response, err error) {
	marshal, err := json.Marshal(payload)
	if err != nil {
		return &http.Response{}, err
	}
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	if authProxy != nil {
		client = authProxy
	}
	request, err := http.NewRequest(
		http.MethodPost,
		"https://authserver.mojang.com/"+endPoint,
		bytes.NewReader(marshal))
	if err != nil {
		return &http.Response{}, err
	}
	request.Header.Set("User-agent", "go-mc")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(request)
	return
}
