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
}

func New(c *api.Client, n *navigate.Navigate) *Dirt {
	return &Dirt{c: c, navigate: n}
}
func (p *Dirt) Start() {
	p.start.x, p.start.y, p.start.z = int(math.Floor(p.c.GetX())), int(math.Floor(p.c.GetY())), int(math.Floor(p.c.GetZ()))
	p.currentClaimCenter = p.start
	p.c.Move(math.Floor(p.c.GetX())+0.5, math.Floor(p.c.GetY()), math.Floor(p.c.GetZ())+0.5, false)
	p.c.Chat("/claim")
	fmt.Println("Current Claim Pos: " + strconv.Itoa(p.currentClaimCenter.x) + " " + strconv.Itoa(p.currentClaimCenter.y) + " " + strconv.Itoa(p.currentClaimCenter.z))
	time.Sleep(200 * time.Millisecond)
	p.dig(p.currentClaimCenter.x+8, p.currentClaimCenter.y-1, p.currentClaimCenter.z)
}
func (p *Dirt) dig(x, y, z int) {
	p.move(x, y+1, z)
	time.Sleep(time.Millisecond * 70)
	p.c.ToggleFly(true)
	p.c.StartBreakBlock(x, y, z, 0)
	p.c.ToggleFly(false)
}
func (p *Dirt) move(x, y, z int) {
	if !p.checkClaimStatus(x, y, z) {
		p.setNewClaim(x, y, z)
	}
	p.navigate.MoveTo(float64(x)+0.5, float64(y), float64(z)+0.5)
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
	x, y, z = p.getNewClaimCenterToBlock(x, y, z)
	p.navigate.MoveTo(float64(x)+0.5, float64(y), float64(z)+0.5)
	p.c.Chat("/delc")
	p.c.Chat("/claim")
	fmt.Println("New Claim Center:", x, y, z)
	time.Sleep(time.Millisecond * 1000)
}
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
