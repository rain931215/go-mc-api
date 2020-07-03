package dirt

import (
	"fmt"
	"github.com/rain931215/go-mc-api/api"
	"math"
	"time"
)

type BlockPos struct {
	x, y, z int
}
type Dirt struct {
	c                      *api.Client
	closedSet              []BlockPos
	startX, startY, startZ int
}

func New(c *api.Client) *Dirt {
	return &Dirt{c: c}
}
func (dirt *Dirt) Start() {
	dirt.c.Move(math.Floor(dirt.c.GetX())+0.5, math.Floor(dirt.c.GetY()), math.Floor(dirt.c.GetZ())+0.5, false)
	time.Sleep(50 * time.Millisecond)
	dirt.startX, dirt.startY, dirt.startZ = int(math.Floor(dirt.c.GetX())), int(math.Floor(dirt.c.GetY())), int(math.Floor(dirt.c.GetZ()))
	dirt.recursiveSearch(dirt.c.GetX(), dirt.c.GetY(), dirt.c.GetZ(), -1, 300)
}
func (dirt *Dirt) recursiveSearch(x, y, z float64, from api.Direction, level int) {
	if y < 0 || y >= 255 {
		return
	}
	if int(math.Floor(x)) < dirt.startX-7 || int(math.Floor(x)) > dirt.startX+7 || int(math.Floor(y)) < dirt.startY-7 || int(math.Floor(y)) > dirt.startY+7 || int(math.Floor(z)) < dirt.startZ-7 || int(math.Floor(z)) > dirt.startZ+7 {
		return
	}
	moveDelay := 10 * time.Millisecond
	dirt.closedSet = append(dirt.closedSet, BlockPos{x: int(math.Floor(x)), y: int(math.Floor(y)), z: int(math.Floor(z))})
	fmt.Println("目前座標 ", x, y, z)
	dirt.c.Move(x, y, z, false)
	time.Sleep(moveDelay)
	var (
		newFeetBlock, newHeadBlock uint32
		blockPos                   BlockPos
	)
	// 向北
	{
		blockPos = BlockPos{x: int(math.Floor(x)), y: int(math.Floor(y)), z: int(math.Floor(z)) - 1}
		if from != api.North && !dirt.getInClosedSet(blockPos) {
			newFeetBlock = dirt.checkBlockStatus(blockPos)
			if newFeetBlock == 0 || newFeetBlock == 10 || newFeetBlock == 1341 {
				if newFeetBlock == 10 || newFeetBlock == 1341 {
					dirt.c.ToggleFly(true)
					dirt.c.StartBreakBlock(blockPos.x, blockPos.y, blockPos.z, api.Bottom)
					dirt.c.FinishBreakBlock(blockPos.x, blockPos.y, blockPos.z, api.Bottom)
					dirt.c.ToggleFly(false)
				}
				blockPos2 := BlockPos{x: int(math.Floor(x)), y: int(math.Floor(y + 1)), z: int(math.Floor(z)) - 1}
				newHeadBlock = dirt.checkBlockStatus(blockPos2)
				if newHeadBlock == 0 || newHeadBlock == 10 || newHeadBlock == 1341 {
					if newHeadBlock == 10 || newHeadBlock == 1341 {
						dirt.c.ToggleFly(true)
						dirt.c.StartBreakBlock(blockPos2.x, blockPos2.y, blockPos2.z, api.South)
						dirt.c.FinishBreakBlock(blockPos2.x, blockPos2.y, blockPos2.z, api.South)
						dirt.c.ToggleFly(false)
					}
					dirt.recursiveSearch(x, y, z-1, api.Top, level-1)
					//dirt.c.Move(math.Floor(x)+0.5, math.Floor(y), math.Floor(z)+0.5, false)
					//time.Sleep(moveDelay)
				}
			}
		}
	}
	// 向南
	{
		blockPos = BlockPos{x: int(math.Floor(x)), y: int(math.Floor(y)), z: int(math.Floor(z)) + 1}
		if from != api.North && !dirt.getInClosedSet(blockPos) {
			newFeetBlock = dirt.checkBlockStatus(blockPos)
			if newFeetBlock == 0 || newFeetBlock == 10 || newFeetBlock == 1341 {
				if newFeetBlock == 10 || newFeetBlock == 1341 {
					dirt.c.ToggleFly(true)
					dirt.c.StartBreakBlock(blockPos.x, blockPos.y, blockPos.z, api.Bottom)
					dirt.c.FinishBreakBlock(blockPos.x, blockPos.y, blockPos.z, api.Bottom)
					dirt.c.ToggleFly(false)
				}
				blockPos2 := BlockPos{x: int(math.Floor(x)), y: int(math.Floor(y + 1)), z: int(math.Floor(z)) + 1}
				newHeadBlock = dirt.checkBlockStatus(blockPos2)
				if newHeadBlock == 0 || newHeadBlock == 10 || newHeadBlock == 1341 {
					if newHeadBlock == 10 || newHeadBlock == 1341 {
						dirt.c.ToggleFly(true)
						dirt.c.StartBreakBlock(blockPos2.x, blockPos2.y, blockPos2.z, api.South)
						dirt.c.FinishBreakBlock(blockPos2.x, blockPos2.y, blockPos2.z, api.South)
						dirt.c.ToggleFly(false)
					}
					dirt.recursiveSearch(x, y, z+1, api.Top, level-1)
					//dirt.c.Move(math.Floor(x)+0.5, math.Floor(y), math.Floor(z)+0.5, false)
					//time.Sleep(moveDelay)
				}
			}
		}
	}
	// 向西
	{
		blockPos = BlockPos{x: int(math.Floor(x)) - 1, y: int(math.Floor(y)), z: int(math.Floor(z))}
		if from != api.North && !dirt.getInClosedSet(blockPos) {
			newFeetBlock = dirt.checkBlockStatus(blockPos)
			if newFeetBlock == 0 || newFeetBlock == 10 || newFeetBlock == 1341 {
				if newFeetBlock == 10 || newFeetBlock == 1341 {
					dirt.c.ToggleFly(true)
					dirt.c.StartBreakBlock(blockPos.x, blockPos.y, blockPos.z, api.Bottom)
					dirt.c.FinishBreakBlock(blockPos.x, blockPos.y, blockPos.z, api.Bottom)
					dirt.c.ToggleFly(false)
				}
				blockPos2 := BlockPos{x: int(math.Floor(x)) - 1, y: int(math.Floor(y + 1)), z: int(math.Floor(z))}
				newHeadBlock = dirt.checkBlockStatus(blockPos2)
				if newHeadBlock == 0 || newHeadBlock == 10 || newHeadBlock == 1341 {
					if newHeadBlock == 10 || newHeadBlock == 1341 {
						dirt.c.ToggleFly(true)
						dirt.c.StartBreakBlock(blockPos2.x, blockPos2.y, blockPos2.z, api.South)
						dirt.c.FinishBreakBlock(blockPos2.x, blockPos2.y, blockPos2.z, api.South)
						dirt.c.ToggleFly(false)
					}
					dirt.recursiveSearch(x-1, y, z, api.Top, level-1)
					//dirt.c.Move(math.Floor(x)+0.5, math.Floor(y), math.Floor(z)+0.5, false)
					//time.Sleep(moveDelay)
				}
			}
		}
	}
	// 向東
	{
		blockPos = BlockPos{x: int(math.Floor(x)) + 1, y: int(math.Floor(y)), z: int(math.Floor(z))}
		if from != api.North && !dirt.getInClosedSet(blockPos) {
			newFeetBlock = dirt.checkBlockStatus(blockPos)
			if newFeetBlock == 0 || newFeetBlock == 10 || newFeetBlock == 1341 {
				if newFeetBlock == 10 || newFeetBlock == 1341 {
					dirt.c.ToggleFly(true)
					dirt.c.StartBreakBlock(blockPos.x, blockPos.y, blockPos.z, api.Bottom)
					dirt.c.FinishBreakBlock(blockPos.x, blockPos.y, blockPos.z, api.Bottom)
					dirt.c.ToggleFly(false)
				}
				blockPos2 := BlockPos{x: int(math.Floor(x)) + 1, y: int(math.Floor(y + 1)), z: int(math.Floor(z))}
				newHeadBlock = dirt.checkBlockStatus(blockPos2)
				if newHeadBlock == 0 || newHeadBlock == 10 || newHeadBlock == 1341 {
					if newHeadBlock == 10 || newHeadBlock == 1341 {
						dirt.c.ToggleFly(true)
						dirt.c.StartBreakBlock(blockPos2.x, blockPos2.y, blockPos2.z, api.South)
						dirt.c.FinishBreakBlock(blockPos2.x, blockPos2.y, blockPos2.z, api.South)
						dirt.c.ToggleFly(false)
					}
					dirt.recursiveSearch(x+1, y, z, api.Top, level-1)
					//dirt.c.Move(math.Floor(x)+0.5, math.Floor(y), math.Floor(z)+0.5, false)
					//time.Sleep(moveDelay)
				}
			}
		}
	}
	// 向上
	{
		blockPos = BlockPos{x: int(math.Floor(x)), y: int(math.Floor(y)) + 1, z: int(math.Floor(z))}
		if from != api.Top && !dirt.getInClosedSet(blockPos) {
			newHeadBlock = dirt.checkBlockStatus(blockPos)
			if newHeadBlock == 0 || newHeadBlock == 10 || newHeadBlock == 1341 {
				if newHeadBlock == 10 || newHeadBlock == 1341 {
					dirt.c.ToggleFly(true)
					dirt.c.StartBreakBlock(blockPos.x, blockPos.y, blockPos.z, api.Top)
					dirt.c.FinishBreakBlock(blockPos.x, blockPos.y, blockPos.z, api.Top)
					dirt.c.ToggleFly(false)
				}
				dirt.recursiveSearch(x, y+1, z, api.Bottom, level-1)
				dirt.c.Move(math.Floor(x)+0.5, math.Floor(y), math.Floor(z)+0.5, false)
				time.Sleep(moveDelay)
			}
		}
	}
	// 向下
	{
		blockPos = BlockPos{x: int(math.Floor(x)), y: int(math.Floor(y)) - 1, z: int(math.Floor(z))}
		if from != api.Bottom && !dirt.getInClosedSet(blockPos) {
			newFeetBlock = dirt.checkBlockStatus(blockPos)
			if newFeetBlock == 0 || newFeetBlock == 10 || newFeetBlock == 1341 {
				if newFeetBlock == 10 || newFeetBlock == 1341 {
					dirt.c.ToggleFly(true)
					dirt.c.StartBreakBlock(blockPos.x, blockPos.y, blockPos.z, api.Bottom)
					dirt.c.FinishBreakBlock(blockPos.x, blockPos.y, blockPos.z, api.Bottom)
					dirt.c.ToggleFly(false)
				}
				dirt.recursiveSearch(x, y-1, z, api.Top, level-1)
				dirt.c.Move(math.Floor(x)+0.5, math.Floor(y), math.Floor(z)+0.5, false)
				time.Sleep(moveDelay)
			}
		}
	}
}
func (dirt *Dirt) checkBlockStatus(pos BlockPos) uint32 {
	return uint32(dirt.c.World.GetBlockStatus(pos.x, pos.y, pos.z))
}
func (dirt *Dirt) getInClosedSet(pos BlockPos) bool {
	if dirt.closedSet != nil {
		for _, v := range dirt.closedSet {
			if v.x == pos.x && v.y == pos.y && v.z == pos.z {
				return true
			}
		}
	}
	return false
}
