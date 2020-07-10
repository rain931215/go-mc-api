package dirt

import (
	"fmt"
	"github.com/rain931215/go-mc-api/api"
	"math"
	"strconv"
	"time"
)

type BlockPos struct {
	x, y, z int
}
type Dirt struct {
	c                  *api.Client
	start              BlockPos
	currentClaimCenter BlockPos
}

func New(c *api.Client) *Dirt {
	return &Dirt{c: c}
}
func (dirt *Dirt) Start() {
	dirt.start.x, dirt.start.y, dirt.start.z = int(math.Floor(dirt.c.GetX())), int(math.Floor(dirt.c.GetY())), int(math.Floor(dirt.c.GetZ()))
	dirt.currentClaimCenter = dirt.start
	dirt.c.Move(math.Floor(dirt.c.GetX())+0.5, math.Floor(dirt.c.GetY()), math.Floor(dirt.c.GetZ())+0.5, false)
	fmt.Println("Current Claim Pos: " + strconv.Itoa(dirt.currentClaimCenter.x) + " " + strconv.Itoa(dirt.currentClaimCenter.y) + " " + strconv.Itoa(dirt.currentClaimCenter.z))
	time.Sleep(50 * time.Millisecond)
	if !dirt.checkClaimStatus(dirt.currentClaimCenter.x+8, dirt.currentClaimCenter.y, dirt.currentClaimCenter.z) {
		fmt.Println("Reset Claim")
		dirt.setNewClaim(dirt.getNewClaimCenterToBlock(dirt.currentClaimCenter.x+8, dirt.currentClaimCenter.y, dirt.currentClaimCenter.z))
	}
}
func (dirt *Dirt) checkClaimStatus(x, y, z int) bool {
	if Abs(x-dirt.currentClaimCenter.x) > 7 {
		return false
	}
	if Abs(z-dirt.currentClaimCenter.z) > 7 {
		return false
	}
	return true
}
func (dirt *Dirt) getNewClaimCenterToBlock(x, y, z int) (newX, newY, newZ int) {
	newY = y
	if x > dirt.currentClaimCenter.x {
		newX = dirt.currentClaimCenter.x + 7
	} else if x < dirt.currentClaimCenter.x {
		newX = dirt.currentClaimCenter.x - 7
	} else {
		newX = dirt.currentClaimCenter.x
	}
	if z > dirt.currentClaimCenter.z {
		newZ = dirt.currentClaimCenter.z + 7
	} else if z < dirt.currentClaimCenter.z {
		newZ = dirt.currentClaimCenter.z - 7
	} else {
		newZ = dirt.currentClaimCenter.z
	}
	return
}
func (dirt *Dirt) setNewClaim(x, y, z int) {
	fmt.Println("New Claim Center:", x, y, z)
}
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
