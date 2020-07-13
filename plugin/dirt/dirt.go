package dirt

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/rain931215/go-mc-api/api"
	"github.com/rain931215/go-mc-api/plugin/navigate"
)

type block struct {
	pos  blockPos
	cost uint
}
type blockPos struct {
	x, y, z int
}
type Dirt struct {
	c                  *api.Client
	navigate           *navigate.Navigate
	start              blockPos
	currentClaimCenter blockPos
	whiteList          []uint32
	blocks             []*block
}

func New(c *api.Client, n *navigate.Navigate) *Dirt {
	return &Dirt{c: c, navigate: n}
}
func (p *Dirt) Start() {
	p.whiteList = []uint32{8, 9, 10, 11, 12, 13}
	p.start.x, p.start.y, p.start.z = int(math.Floor(p.c.GetX())), int(math.Floor(p.c.GetY())), int(math.Floor(p.c.GetZ()))
	p.currentClaimCenter = p.start
	p.c.Move(math.Floor(p.c.GetX())+0.5, math.Floor(p.c.GetY()), math.Floor(p.c.GetZ())+0.5, false)
	p.c.Chat("/delallc")
	p.c.Chat("/claim")
	p.refrshBlockList()
	fmt.Println("Current Claim Pos: " + strconv.Itoa(p.currentClaimCenter.x) + " " + strconv.Itoa(p.currentClaimCenter.y) + " " + strconv.Itoa(p.currentClaimCenter.z))
	time.Sleep(200 * time.Millisecond)
	println(len(p.blocks))
	go func() {
		for {
			if len(p.blocks) < 1 {
				//rtp
				println("return")
				return
			}
			p.dig(p.blocks[0].pos.x, p.blocks[0].pos.y, p.blocks[0].pos.z)
			p.blocks = p.blocks[1:]
			time.Sleep(10 * time.Millisecond)
		}
	}()
}
func (p *Dirt) dig(x, y, z int) {
	dx := p.c.GetX() - (float64(x) + 0.5)
	dy := p.c.GetY() - (float64(y) + 0.5)
	dz := p.c.GetZ() - (float64(z) + 0.5)
	if (dx*dx)+(dy*dy)+(dz*dz) < 20 {
		time.Sleep(time.Millisecond * 70)
		p.c.ToggleFly(true)
		p.c.StartBreakBlock(x, y, z, 0)
		p.c.ToggleFly(false)
		return
	} else if p.move(x, y+1, z) {
		time.Sleep(time.Millisecond * 50)
		p.c.ToggleFly(true)
		p.c.StartBreakBlock(x, y, z, 0)
		p.c.ToggleFly(false)
	}
}
func (p *Dirt) move(x, y, z int) bool {
	if !p.checkClaimStatus(x, y, z) {
		p.setNewClaim(x, y, z)
		time.Sleep(time.Millisecond * 1000)
	}
	return p.navigate.MoveTo(float64(x)+0.5, float64(y), float64(z)+0.5)
}
func (p *Dirt) checkClaimStatus(x, y, z int) bool {
	if abs(x-p.currentClaimCenter.x) > 7 {
		return false
	}
	if abs(z-p.currentClaimCenter.z) > 7 {
		return false
	}
	return true
}
func (p *Dirt) getNewClaimCenterToBlock(x, y, z int) (newX, newY, newZ int) {
	newY = y
	if x > p.currentClaimCenter.x {
		newX = p.currentClaimCenter.x + 7
	} else if x < p.currentClaimCenter.x {
		newX = p.currentClaimCenter.x - 7
	} else {
		newX = p.currentClaimCenter.x
	}
	if z > p.currentClaimCenter.z {
		newZ = p.currentClaimCenter.z + 7
	} else if z < p.currentClaimCenter.z {
		newZ = p.currentClaimCenter.z - 7
	} else {
		newZ = p.currentClaimCenter.z
	}
	return
}
func (p *Dirt) setNewClaim(x, y, z int) {
	x, y, z = p.getNewClaimCenterToBlock(x, y+1, z)
	if p.c.World.GetBlockStatus(x, y, z) == 0 && p.c.World.GetBlockStatus(x, y+1, z) == 0 {
		p.navigate.MoveTo(float64(x)+0.5, float64(y), float64(z)+0.5)
		p.c.Chat("/delallc")
		p.c.Chat("/claim")
		p.currentClaimCenter = blockPos{x: x, y: y, z: z}
		p.refrshBlockList()
		fmt.Println("New Claim Center:", x, y, z)
		return
	}
	fmt.Println(x, y+1, z)
	p.setNewClaim(x, y+1, z)
}
func (p *Dirt) refrshBlockList() {
	p.blocks = make([]*block, 0)
	p.blocks = append(p.blocks, &block{pos: blockPos{x: p.currentClaimCenter.x, y: p.currentClaimCenter.y - 1, z: p.currentClaimCenter.z}, cost: 0})
	for x := p.currentClaimCenter.x - 12; x < p.currentClaimCenter.x+12; x++ {
		for z := p.currentClaimCenter.z - 12; z < p.currentClaimCenter.z+12; z++ {
			for y := p.currentClaimCenter.z + 5; y > 60; y-- {
				if p.checkBlock(uint32(p.c.World.GetBlockStatus(x, y, z))) {
					blockPos := blockPos{x: x, y: y, z: z}
					cost := abs(x-p.currentClaimCenter.x) + abs(z-p.currentClaimCenter.z) + abs(y-p.currentClaimCenter.y)
					blocK := &block{pos: blockPos, cost: uint(cost)}
					result := false
					for i := 0; i < len(p.blocks); i++ {
						if blocK.cost == p.blocks[i].cost || blocK.cost < p.blocks[i].cost {
							p.blocks = append(p.blocks[:i], append([]*block{blocK}, p.blocks[i:]...)...)
							result = true
							break
						}
					}
					if !result {
						p.blocks = append(p.blocks, blocK)
					}
				}
			}
		}
	}
}
func (p *Dirt) checkBlock(ID uint32) bool {
	var pass bool
	for i := 0; i < len(p.whiteList); i++ {
		if ID == p.whiteList[i] {
			pass = true
			return pass
		}
	}
	return pass
}
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
