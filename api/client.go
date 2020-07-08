package api

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/enriquebris/goconcurrentqueue"
	"github.com/rain931215/go-mc-api/api/world"
	"net"
	"time"
)

const bufferPacketChannelSize int = 300

// 改寫的客戶端結構
type Client struct {
	Native     *bot.Client
	World      *world.World
	Inventory  *inventory
	Auth       *AuthInfo
	EntityList *EntityList
	*Position
	packetOutStream *goconcurrentqueue.FixedFIFO
	inStatusChannel chan error
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
	client.packetOutStream = goconcurrentqueue.NewFixedFIFO(bufferPacketChannelSize)
	//client.packetChannel.outChannel = make(chan *pk.Packet, bufferPacketChannelSize)
	client.inStatusChannel = make(chan error, 1)
	go func() {
		for {
			if obj, err := client.packetOutStream.DequeueOrWaitForNextElement(); err == nil {
				if p, ok := obj.(*pk.Packet); ok && p != nil {
					if client == nil || client.Native == nil || client.Native.Conn() == nil {
						continue
					}
					_ = client.Native.SendPacket(*p)
				}
			}
		}
	}()
	go func() {
		var (
			incomeErr error
		)
		for {
			<-client.inStatusChannel
			client.connected = true // 設定連線狀態
			for {
				incomeErr = nil
				p, err := client.Native.Conn().ReadPacket()
				if err != nil {
					incomeErr = err
					break
				}
				twoBreak := false
				switch p.ID {
				case 0x1b: // 0x1b = Disconnect (play) https://wiki.vg/Protocol#Disconnect_.28play.29
					var msg chat.Message
					if msg.Decode(bufio.NewReader(bytes.NewReader(p.Data))) == nil {
						//TODO (Async Events)
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
					twoBreak = true
					break
				case data.KeepAliveClientbound:
					var ID pk.Long
					if err := ID.Decode(bufio.NewReader(bytes.NewReader(p.Data))); err == nil {
						go func() {
							_ = client.Native.SendPacket(pk.Marshal(data.KeepAliveServerbound, ID))
							time.Sleep(5 * time.Second)
							_ = client.Native.SendPacket(pk.Marshal(data.KeepAliveServerbound, ID))
						}()
					}
					break
				default:
					if err := client.handlePacket(&p); err != nil {
						incomeErr = err
						twoBreak = true
					}
					break
				}
				if twoBreak {
					break
				}
			}
			client.connected = false // 設定連線狀態
			client.inStatusChannel <- incomeErr
		}
	}()
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
	c.inStatusChannel <- nil
	return <-c.inStatusChannel
}
func (c *Client) SendPacket(packet pk.Packet) {
	if c.packetOutStream == nil {
		return
	} else if err := c.packetOutStream.Enqueue(&packet); err != nil {
		fmt.Println(fmt.Sprintf("Enqueue packet error: %v", err))
	}
}
func (c *Client) Connected() bool {
	return c.connected
}
