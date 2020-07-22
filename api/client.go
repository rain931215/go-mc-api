package api

import (
	"bytes"
	"fmt"
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/rain931215/go-mc-api/api/world"
	"net"
	"time"
)

//const bufferPacketChannelSize int = 100

// 改寫的客戶端結構
type Client struct {
	Native     *bot.Client
	World      *world.World
	Inventory  *inventory
	Auth       *AuthInfo
	EntityList *EntityList
	*Position
	packetInStream  chan *pk.Packet
	packetOutStream chan *pk.Packet
	Event           Events
	connected       bool
}

type AuthInfo struct {
	ID, UUID, AccessToken string
}

// 生成新的客戶端
func NewClient() (client *Client) {
	client = new(Client)
	client.Native = bot.NewClient()
	client.World = &world.World{Chunks: make(map[world.ChunkLoc]*world.Chunk)}
	client.Inventory = NewInventory()
	client.Position = new(Position)
	client.Event = Events{}
	client.Auth = &AuthInfo{ID: "steve"}
	client.EntityList = NewEntityList()
	client.packetInStream = make(chan *pk.Packet, 1024)
	client.packetOutStream = make(chan *pk.Packet, 1024)
	go func(pChannel <-chan *pk.Packet) {
		for v := range pChannel {
			if v == nil {
				continue
			}
			_ = client.Native.SendPacket(*v)
		}
	}(client.packetOutStream)
	go func(pChannel <-chan *pk.Packet) {
		for v := range pChannel {
			if v == nil {
				continue
			}
			currentTime := time.Now().UnixNano()
			if v.ID == 0x1b {
				var msg chat.Message
				if v.Scan(&msg) == nil {
					if len(client.Event.disconnectHandlers) < 1 {
						continue
					}
					for i := 0; i < len(client.Event.disconnectHandlers); i++ {
						v := client.Event.disconnectHandlers[i]
						if v == nil {
							continue
						}
						if v(msg) {
							client.Event.disconnectHandlers = append(client.Event.disconnectHandlers[:i], client.Event.disconnectHandlers[i+1:]...)
							i--
						}
					}
				}
			} else {
				_ = client.handlePacket(v)
			}
			if diff := time.Now().UnixNano() - currentTime; diff > 30000000 { // 大於30ms就輸出時間
				fmt.Println(fmt.Sprintf("封包超過正常時間:%v.%v毫秒", diff/1000000, diff%1000000))
			}
		}
	}(client.packetInStream)
	return
}

// 加入伺服器
func (c *Client) JoinServer(ip string, port int) error {
	return c.JoinServerWithDialer(ip, port, &net.Dialer{Timeout: 30 * time.Second})
}
func (c *Client) JoinServerWithDialer(ip string, port int, dialer *net.Dialer) error {
	c.Native.Name, c.Native.Auth.UUID, c.Native.AsTk = c.Auth.ID, c.Auth.UUID, c.Auth.AccessToken
	if port < 0 || port > 65535 {
		panic("try join server error: except port assigned")
	}
	return c.Native.JoinServerWithDialer(dialer, fmt.Sprintf("%s:%d", ip, port))
}
func (c *Client) HandleGame() error {
	c.connected = true
	defer func() { c.connected = false }()
	for {
		if c == nil || c.Native == nil || c.Native.Conn() == nil {
			return nil
		}
		p, err := c.Native.Conn().ReadPacket()
		if err != nil {
			return err
		}
		if p.ID == data.KeepAliveClientbound {
			var ID pk.Long
			if ID.Decode(bytes.NewReader(p.Data)) == nil {
				_ = c.Native.SendPacket(pk.Marshal(data.KeepAliveServerbound, ID))
			}
			continue
		}
		c.packetInStream <- &p
	}
}
func (c *Client) SendPacket(packet pk.Packet) {
	if c.packetOutStream == nil {
		return
	}
	c.packetOutStream <- &packet
}
func (c *Client) Connected() bool {
	return c.connected
}
