package navigate

import (
	"math"

	"github.com/rain931215/go-mc-api/api"
)

type pathFinder struct {
	c                                     *api.Client
	startPointX, startPointY, startPointZ int
	endPointX, endPointY, endPointZ       int
	startPos, endPos                      pos
	openNodeList, closeNodeList           map[pos]*node
	dx, dz                                uint16
	count                                 uint16
	node                                  *node
	FList                                 []*node
}

func setNewPath(x, y, z float64, c *api.Client) *pathFinder {
	f := new(pathFinder)
	f.c = c
	f.startPointX = int(math.Floor(f.c.GetX()))
	f.startPointY = int(math.Floor(f.c.GetY()))
	f.startPointZ = int(math.Floor(f.c.GetZ()))
	f.startPos = pos{x: 0, y: 0, z: 0}
	f.endPointX = int(math.Floor(x))
	f.endPointY = int(math.Floor(y))
	f.endPointZ = int(math.Floor(z))
	f.endPos = pos{x: f.endPointX - f.startPointX, y: f.endPointY - f.startPointY, z: f.endPointZ - f.startPointZ}
	f.dx = simpleAbs(f.endPos.x-f.startPos.x) * 2
	f.dz = simpleAbs(f.endPos.z-f.startPos.z) * 2
	f.openNodeList = make(map[pos]*node)
	f.closeNodeList = make(map[pos]*node)
	pos := pos{x: 0, y: 0, z: 0}
	firstNode := newNode(pos, new(node))
	f.openNodeList[pos] = firstNode
	f.FList = append(f.FList, firstNode)
	return f
}

func (f *pathFinder) getNodes() []*node {
	if f.c.World.GetBlockStatus(f.endPointX, f.endPointY, f.endPointZ) != 0 || f.c.World.GetBlockStatus(f.endPointX, f.endPointY+1, f.endPointZ) != 0 {
		println("wrong")
		return make([]*node, 1)
	}
	tempCount := 0
	for {
		if len(f.FList) < 1 {
			println("wrong")
			return make([]*node, 1)
		}

		f.node = f.FList[0]
		f.FList = f.FList[1:]
		if f.node.pos == f.endPos {
			println("finish")
			return f.node.returnNodes([]*node{f.node, f.node})
		}
		tempCount++
		if tempCount > 200000 {
			return make([]*node, 1)
		}
		delete(f.openNodeList, f.node.pos)
		f.closeNodeList[f.node.pos] = f.node

		for offSet := -1; offSet < 2; offSet += 2 {
			f.openNewNode(pos{x: f.node.pos.x + offSet, y: f.node.pos.y, z: f.node.pos.z})
			f.openNewNode(pos{x: f.node.pos.x, y: f.node.pos.y, z: f.node.pos.z + offSet})
			pos := pos{x: f.node.pos.x, y: offSet + f.node.pos.y, z: f.node.pos.z}
			if y := pos.y + f.startPointY; y < -2 || y > 257 {
				continue
			}
			f.openNewNode(pos)
		}
	}
}

// 插入新節點
func (f *pathFinder) openNewNode(p pos) {
	if f.nodeRule(p) { // 如果節點需要計算
		node := newNode(p, f.node)                       // 產生新節點
		node.f = node.cost + node.getGuessCost(f.endPos) // 取得成本
		f.fListInsert(node)                              // 插入節點
		f.openNodeList[p] = node                         // 加入節點
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
		if v.getGuessCost(f.endPos) < f.node.getGuessCost(f.endPos) {
			v.lastNode = f.node
			v.setCost()
			v.f = v.cost + v.getGuessCost(f.endPos)
			f.clearNode(v)
			f.fListInsert(v)
		}
		return false
	}

	// 得出絕對座標
	x := f.startPointX + p.x
	y := f.startPointY + p.y
	z := f.startPointZ + p.z
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
