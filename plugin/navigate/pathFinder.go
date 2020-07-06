package navigate

import (
	"fmt"
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
	var nodes = make([]*node, 1)
	if f.c.World.GetBlockStatus(f.endPointX, f.endPointY, f.endPointZ) != 0 || f.c.World.GetBlockStatus(f.endPointX, f.endPointY+1, f.endPointZ) != 0 {
		println("wrong")
		return nodes
	}
	tempCount := 0
	for {
		tempCount++
		fmt.Println(tempCount)

		if len(f.FList) < 1 {
			println("wrong")
			return nodes
		}

		f.node = f.FList[0]
		f.FList = f.FList[1:]
		if f.node.pos == f.endPos {
			println("finish")
			nodes = append(nodes, f.node)
			nodes = f.node.returnNodes(nodes)
			return nodes
		}

		delete(f.openNodeList, f.node.pos)
		f.closeNodeList[f.node.pos] = f.node
		for x := -1; x < 2; x += 2 {
			pos := pos{x: x + f.node.pos.x, y: f.node.pos.y, z: f.node.pos.z}
			f.openNewNode(pos)
		}
		for y := -1; y < 2; y += 2 {
			pos := pos{x: f.node.pos.x, y: y + f.node.pos.y, z: f.node.pos.z}
			y := pos.y + f.startPointY
			if y < -2 || y > 255 {
				continue
			}
			f.openNewNode(pos)
		}
		for z := -1; z < 2; z += 2 {
			pos := pos{x: f.node.pos.x, y: f.node.pos.y, z: z + f.node.pos.z}
			f.openNewNode(pos)
		}
	}
}

func (f *pathFinder) openNewNode(p pos) {
	if f.nodeRule(p) {
		node := newNode(p, f.node)
		node.f = node.cost + node.getGuessCost(f.endPos)
		f.flistRefrsh(node)
		f.openNodeList[p] = node
	}
}

func (f *pathFinder) flistRefrsh(node *node) {
	for i := 0; i < len(f.FList); i++ {
		if node.f == f.FList[i].f || node.f < f.FList[i].f {
			f.FList = append(append(f.FList[:i], node), f.FList[i:]...)
			break
		}
	}
	f.FList = append(f.FList, node)
}

func (f *pathFinder) nodeRule(p pos) bool {
	var pass bool
	if _, ok := f.closeNodeList[p]; ok == true {
		return false
	}
	if v, ok := f.openNodeList[p]; ok == true {
		if v.getGuessCost(f.endPos) < f.node.getGuessCost(f.endPos) {
			v.lastNode = f.node
			v.setCost()
			v.f = v.cost + v.getGuessCost(f.endPos)
			f.flistRefrsh(v)
		}
		return false
	}
	x := f.startPointX + p.x
	y := f.startPointY + p.y
	z := f.startPointZ + p.z

	//println(x, y, z, f.c.World.GetBlockStatus(x, y, z))
	feetBlock := f.c.World.GetBlockStatus(x, y, z)
	headBlock := f.c.World.GetBlockStatus(x, y+1, z)
	if (feetBlock == 0 || feetBlock == 9130) && (headBlock == 0 || headBlock == 9130) {
		pass = true
	}
	return pass
}

func min(l []*node) (min *node) {
	min = l[0]
	for _, v := range l {
		if v.f < min.f {
			min = v
		}
	}
	return
}
func (f *pathFinder) getBlock(pos1 pos) uint32 {
	return uint32(f.c.World.GetBlockStatus(pos1.x, pos1.y, pos1.z))
}
