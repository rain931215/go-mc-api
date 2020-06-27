package mcfalloutcmd

import (
	"fmt"
	"strings"

	"github.com/Tnze/go-mc/chat"
	"github.com/fsnotify/fsnotify"
	"github.com/rain931215/go-mc-api/api"
	"github.com/spf13/viper"
)

//Func is the type of command's method
type Func = func(c *api.Client, Text string)

//Command is contained command's name and command's method
type Command struct {
	name   string
	method Func
}

var (
	client    *api.Client
	whiteList []string
	cmdList   []*Command
)

// Load Plugin
func Load(c *api.Client) {
	client = c
	c.Event.AddEventHandler(main, "chat")

	//Load whiteList
	viper.SetConfigName("whiteList")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./plugin/mcfalloutCmd")
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	whiteList = viper.GetStringSlice("admin")
	//熱插拔
	viper.OnConfigChange(func(e fsnotify.Event) {
		whiteList = viper.GetStringSlice("admin")
		fmt.Println("White List Change")
	})
	//Load defaultCommand
	AddCmd("say", say)
	AddCmd("addadmin", addAdmin)
}

func main(msg chat.Message) (bool, error) {
	var text string = msg.ClearString()
	for id := 0; id < len(whiteList); id++ {
		if strings.Index(text, "[收到私訊 "+whiteList[id]) == 0 {
			text = strings.TrimPrefix(text, "[收到私訊 "+whiteList[id]+"] : ")
			for i := 0; i < len(cmdList); i++ {
				if strings.Index(text, cmdList[i].name) == 0 {
					text = strings.TrimPrefix(text, cmdList[i].name+" ")
					cmdList[i].method(client, text)
					return false, nil
				}
			}
			return false, nil
		}
	}
	return false, nil
}

// AddCmd _
func AddCmd(name string, command Func) {
	newCommand := new(Command)
	newCommand.name = name
	newCommand.method = command
	cmdList = append(cmdList, newCommand)
}
