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

const bufferPacketChannelSize int = 100

// 改寫的客戶端結構
type Client struct {
	Native     *bot.Client
	World      *world.World
	Inventory  *inventory
	Auth       *AuthInfo
	EntityList *EntityList
	*Position
	packetChannel struct {
		inChannel, outChannel chan *pk.Packet
		inStatusChannel       chan error
	}
	Event Events
	Status
}
type Status struct {
	connected bool
}

type AuthInfo struct {
	ID, UUID, AccessToken string
}

// 生成新的客戶端
func NewClient() *Client {
	client := new(Client)
	client.Native = bot.NewClient()
	client.World = &world.World{Chunks: make(map[world.ChunkLoc]*world.Chunk)}
	client.Inventory = NewInventory()
	client.Position = new(Position)
	client.Event = Events{}
	client.Auth = &AuthInfo{ID: "steve"}
	client.EntityList = NewEntityList()
	client.packetChannel.inChannel = make(chan *pk.Packet, bufferPacketChannelSize)
	client.packetChannel.outChannel = make(chan *pk.Packet, bufferPacketChannelSize)
	client.packetChannel.inStatusChannel = make(chan error, 1)
	go func() {
		for {
			if p := <-client.packetChannel.outChannel; p != nil {
				if client.Connected() {
					_ = client.Native.Conn().WritePacket(*p)
				}
			}
		}
	}()
	go func() {
		for {
			<-client.packetChannel.inStatusChannel
			client.Status.connected = true
			var incomeErr error
			for {
				incomeErr = nil
				p, err := client.Native.Conn().ReadPacket()
				if err != nil {
					incomeErr = err
					break
				}
				switch p.ID {
				case data.KeepAliveClientbound:
					var ID pk.Long
					if err := ID.Decode(bytes.NewReader(p.Data)); err != nil {
						incomeErr = err
						break
					} else {
						_ = client.Native.Conn().WritePacket(pk.Marshal(data.KeepAliveServerbound, ID))
					}
					break
				case data.DisconnectPlay:
					var (
						msg chat.Message
					)
					if msg.Decode(bytes.NewReader(p.Data)) == nil {
						if client.Event.disconnectHandlers == nil || len(client.Event.disconnectHandlers) < 1 {
							break
						}
						for _, v := range client.Event.disconnectHandlers {
							if v == nil {
								continue
							}
							pass, err := v(msg)
							if err != nil {
								incomeErr = err
								fmt.Println("Disconnect event error" + err.Error())
							}
							if pass {
								break
							}
						}
					}
					break
				default:
					err := client.handlePacket(&p)
					if err != nil {
						incomeErr = err
					}
					break
				}
				if p.ID == data.DisconnectPlay || incomeErr != nil {
					break
				}
			}
			client.Status.connected = false
			client.packetChannel.inStatusChannel <- incomeErr
		}
	}()
	return client
}

// 加入伺服器
func (c *Client) JoinServer(ip string, port int) error {
	return c.JoinServerWithDialer(ip, port, &net.Dialer{Timeout: 120 * time.Second})
}
func (c *Client) JoinServerWithDialer(ip string, port int, dialer *net.Dialer) error {
	c.Native.Name, c.Native.Auth.UUID, c.Native.AsTk = c.Auth.ID, c.Auth.UUID, c.Auth.AccessToken
	if port < 0 || port > 65535 {
		panic("try join server error: except port assigned")
	}
	return c.Native.JoinServerWithDialer(dialer, fmt.Sprintf("%s:%d", ip, port))
}
func (c *Client) HandleGame() error {
	c.packetChannel.inStatusChannel <- nil
	return <-c.packetChannel.inStatusChannel
}
func (c *Client) SendPacket(packet pk.Packet) {
	if c.packetChannel.outChannel != nil {
		c.packetChannel.outChannel <- &packet
	}
}
func (c *Client) Connected() bool {
	return c.Status.connected
}
