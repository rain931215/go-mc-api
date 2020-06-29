package mcfalloutraid

import (
	"github.com/rain931215/go-mc-api/api"
	"time"
)

type McFalloutRaid struct {
	c *api.Client
}

func New(c *api.Client) (raid *McFalloutRaid) {
	raid = new(McFalloutRaid)
	raid.c = c
	go func() {
		for {
			if raid.c.Connected() {
				raid.c.Chat("/warp " + raid.c.Native.Name)
				time.Sleep(30 * time.Second)
				raid.c.Chat("/home a")
				time.Sleep(1 * time.Second)
				raid.c.Chat("/warp " + raid.c.Native.Name)
				time.Sleep(30 * time.Second)
				raid.c.Chat("/home b")
				time.Sleep(1 * time.Second)
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return
}
