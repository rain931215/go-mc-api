package navigate

import (
	"math"

	"github.com/rain931215/go-mc-api/api"
)

type pathFinder struct {
	c                                                                  *api.Client
	startAbsolutePos, endAbsolutePos, startRelativePos, endRelativePos pos
	openNodeList, closeNodeList                                        map[pos]*node
	//dx, dz                                                     uint16
	count    uint16
	node     *node
	FList    []*node
	stopFlag bool
}

func setNewPath(c *api.Client, startPosX, startPosY, startPosZ, endPosX, endPosY, endPosZ float64) *pathFinder {
	f := new(pathFinder)
	f.c = c
	f.startAbsolutePos = pos{x: int(math.Floor(startPosX)), y: int(math.Floor(startPosY)), z: int(math.Floor(startPosZ))}
	f.startRelativePos = pos{x: 0, y: 0, z: 0}
	f.endAbsolutePos = pos{x: int(math.Floor(endPosX)), y: int(math.Floor(endPosY)), z: int(math.Floor(endPosZ))}
	f.endRelativePos = pos{x: f.endAbsolutePos.x - f.startAbsolutePos.x, y: f.endAbsolutePos.y - f.startAbsolutePos.y, z: f.endAbsolutePos.z - f.startAbsolutePos.z}
	//f.dx = simpleAbs(f.endRelativePos.x-f.startRelativePos.x) * 2
	//f.dz = simpleAbs(f.endRelativePos.z-f.startRelativePos.z) * 2
	f.openNodeList = make(map[pos]*node)
	f.closeNodeList = make(map[pos]*node)
	pos := pos{x: 0, y: 0, z: 0}
	firstNode := newNode(pos, new(node))
	f.openNodeList[pos] = firstNode
	f.FList = append(f.FList, firstNode)
	return f
}

func (f *pathFinder) getNodes() (bool, []*node) {
	f.stopFlag = false
	tempCount := 0
	for {
		if f.stopFlag || len(f.FList) < 1 {
			//println("No Path")
			return false, nil
		}
		f.node = f.FList[0]
		f.FList = f.FList[1:]
		if f.node.pos == f.endRelativePos {
			//println("Found Path")
			return true, f.node.returnNodes([]*node{f.node, f.node})
		}
		tempCount++
		if tempCount > 20000 {
			return false, nil
		}
		delete(f.openNodeList, f.node.pos)
		f.closeNodeList[f.node.pos] = f.node
		for offSet := -1; offSet < 2; offSet += 2 {
			f.openNewNode(pos{x: f.node.pos.x + offSet, y: f.node.pos.y, z: f.node.pos.z})
			f.openNewNode(pos{x: f.node.pos.x, y: f.node.pos.y, z: f.node.pos.z + offSet})
			pos := pos{x: f.node.pos.x, y: offSet + f.node.pos.y, z: f.node.pos.z}
			if y := pos.y + f.startAbsolutePos.y; y < -2 || y > 257 {
				continue
			}
			f.openNewNode(pos)
		}
	}
}

func (f *pathFinder) stop() {
	f.stopFlag = true
}

// 插入新節點
func (f *pathFinder) openNewNode(p pos) {
	if f.nodeRule(p) { // 如果節點需要計算
		node := newNode(p, f.node)                               // 產生新節點
		node.f = node.cost + node.getGuessCost(f.endRelativePos) // 取得成本
		f.fListInsert(node)                                      // 插入節點
		f.openNodeList[p] = node                                 // 加入節點
	}
}

// 從小到大排序
func (f *pathFinder) fListInsert(nodeToInsert *node) {
	for i := 0; i < len(f.FList); i++ {
		if nodeToInsert.f == f.FList[i].f || nodeToInsert.f < f.FList[i].f {
			f.FList = append(f.FList[:i], append([]*node{nodeToInsert}, f.FList[i:]...)...)
			return
		}
	}
	f.FList = append(f.FList, nodeToInsert)
}

// 清除已存在的node
func (f *pathFinder) clearNode(nodeToClear *node) {
	for i := 0; i < len(f.FList); i++ {
		if nodeToClear.pos == f.FList[i].pos {
			f.FList = append(f.FList[:i], f.FList[i+1:]...)
			return
		}
	}
}

// 節點判斷
func (f *pathFinder) nodeRule(p pos) bool {
	// 判斷節點是否已經算過
	if _, ok := f.closeNodeList[p]; ok {
		return false
	}

	if v, ok := f.openNodeList[p]; ok {
		if v.getGuessCost(f.endRelativePos) < f.node.getGuessCost(f.endRelativePos) {
			v.lastNode = f.node
			v.setCost()
			v.f = v.cost + v.getGuessCost(f.endRelativePos)
			f.clearNode(v)
			f.fListInsert(v)
		}
		return false
	}

	// 得出絕對座標
	x := f.startAbsolutePos.x + p.x
	y := f.startAbsolutePos.y + p.y
	z := f.startAbsolutePos.z + p.z
	// 取得方塊
	feetBlock := f.c.World.GetBlockStatus(x, y, z)
	headBlock := f.c.World.GetBlockStatus(x, y+1, z)
	// 0 = 空氣 9130 = 洞穴空氣
	if (feetBlock == 0 || feetBlock == 9130) && (headBlock == 0 || headBlock == 9130) {
		return true
	}
	return false
}

func (f *pathFinder) getBlock(pos1 pos) uint32 {
	return uint32(f.c.World.GetBlockStatus(pos1.x, pos1.y, pos1.z))
}
