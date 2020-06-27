package api

import (
	"bytes"
	"fmt"
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/data"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/rain931215/go-mc-api/api/world"
	"net"
	"time"
)

const bufferPacketChannelSize int = 300

// 改寫的客戶端結構
type Client struct {
	Native    *bot.Client
	World     *world.World
	Inventory *inventory
	Auth      *AuthInfo
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
	client.packetChannel.inChannel = make(chan *pk.Packet, bufferPacketChannelSize)
	client.packetChannel.outChannel = make(chan *pk.Packet, bufferPacketChannelSize)
	client.packetChannel.inStatusChannel = make(chan error, 1)
	go func() {
		for {
			if client == nil {
				return
			}
			<-client.packetChannel.inStatusChannel
			client.Status.connected = true
			var incomeErr error
			for {
				p, err := client.Native.Conn().ReadPacket()
				if err != nil {
					incomeErr = err
					break
				}
				if p.ID == data.KeepAliveClientbound {
					var ID pk.Long
					if err := ID.Decode(bytes.NewReader(p.Data)); err != nil {
						incomeErr = err
						continue
					} else {
						_ = client.Native.SendPacket(pk.Marshal(data.KeepAliveServerbound, ID))
					}
					continue
				}
				if err = client.handlePacket(&p); err != nil {
					incomeErr = err
					break
				}
			}
			client.Status.connected = false
			client.packetChannel.inStatusChannel <- incomeErr
		}
	}()
	go func() {
		for {
			p := <-client.packetChannel.outChannel
			if client == nil {
				return
			}
			if p != nil {
				_ = client.Native.Conn().WritePacket(*p)
			}
		}
	}()
	return client
}

// 加入伺服器
func (c *Client) JoinServer(ip string, port int) error {
	return c.JoinServerWithDialer(ip, port, &net.Dialer{Timeout: 30 * time.Second})
}
func (c *Client) JoinServerWithDialer(ip string, port int, dialer *net.Dialer) error {
	c.Native.Name, c.Native.Auth.UUID, c.Native.AsTk = c.Auth.ID, c.Auth.UUID, c.Auth.AccessToken
	if port < 0 || port > 65535 {
		panic("try join server error:port is not in range 0~65535")
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
