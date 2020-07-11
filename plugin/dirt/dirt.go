package dirt

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/rain931215/go-mc-api/api"
	"github.com/rain931215/go-mc-api/plugin/navigate"
)

type blockPos struct {
	x, y, z int
}
type Dirt struct {
	c                  *api.Client
	navigate           *navigate.Navigate
	start              blockPos
	currentClaimCenter blockPos
	whiteList          []uint32
	blocks             []*blockPos
	blockMap           map[blockPos]bool
}

func New(c *api.Client, n *navigate.Navigate) *Dirt {
	return &Dirt{c: c, navigate: n}
}
func (p *Dirt) Start() {
	p.whiteList = []uint32{8, 9, 10, 11, 12, 13}
	p.blockMap = make(map[blockPos]bool)
	p.start.x, p.start.y, p.start.z = int(math.Floor(p.c.GetX())), int(math.Floor(p.c.GetY())), int(math.Floor(p.c.GetZ()))
	p.currentClaimCenter = p.start
	p.c.Move(math.Floor(p.c.GetX())+0.5, math.Floor(p.c.GetY()), math.Floor(p.c.GetZ())+0.5, false)
	p.c.Chat("/claim")
	fmt.Println("Current Claim Pos: " + strconv.Itoa(p.currentClaimCenter.x) + " " + strconv.Itoa(p.currentClaimCenter.y) + " " + strconv.Itoa(p.currentClaimCenter.z))
	time.Sleep(200 * time.Millisecond)
	go func() {
		p.dig(p.start.x, p.start.y-1, p.start.z)
		for {
			if len(p.blocks) < 1 {
				//rtp
				//dig feet block
				println("return")
				return
			}
			p.dig(p.blocks[0].x, p.blocks[0].y, p.blocks[0].z)
			p.blocks = p.blocks[1:]
			time.Sleep(10 * time.Millisecond)
		}
	}()
}
func (p *Dirt) dig(x, y, z int) {
	if p.move(x, y+1, z) {
		for ox := -1; ox < 2; ox++ {
			for oy := -1; oy < 2; oy++ {
				for oz := -1; oz < 2; oz++ {
					if ox+oy+oz == 0 {
						continue
					}
					pos := &blockPos{x: x + ox, y: y + oy, z: z + oz}
					if _, ok := p.blockMap[*pos]; ok {
						continue
					}
					if p.checkBlock(uint32(p.c.World.GetBlockStatus(x+oz, y+oy, z+oz))) {
						println("new block")
						p.blocks = append(p.blocks, pos)
						p.blockMap[*pos] = true
					}
				}
			}
		}
		time.Sleep(time.Millisecond * 80)
		p.c.ToggleFly(true)
		p.c.StartBreakBlock(x, y, z, 0)
		p.c.ToggleFly(false)
	}
}
func (p *Dirt) move(x, y, z int) bool {
	if !p.checkClaimStatus(x, y+2, z) {
		p.setNewClaim(x, y+2, z)
		time.Sleep(time.Millisecond * 1000)
	}
	return p.navigate.MoveTo(float64(x)+0.5, float64(y)+2, float64(z)+0.5)
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
		p.c.Chat("/delc")
		p.c.Chat("/claim")
		//p.c.Chat("delc&claim")
		p.currentClaimCenter = blockPos{x: x, y: y, z: z}
		fmt.Println("New Claim Center:", x, y, z)
		return
	}
	fmt.Println(x, y+1, z)
	p.setNewClaim(x, y+1, z)
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
