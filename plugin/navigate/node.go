package navigate

type pos struct {
	x, y, z int
}

type node struct {
	pos      pos
	lastNode *node
	cost     uint16
	f        uint16
}

func newNode(pos pos, lastNode *node) *node {
	n := new(node)
	n.pos = pos
	n.lastNode = lastNode
	n.setCost()
	return n
}

func (n *node) setCost() {
	n.cost = n.lastNode.cost + n.getCost()
}

func (n *node) getCost() uint16 {
	/*
		var count uint8
		if n.lastNode.pos.x != n.pos.x {
			count++
		}
		if n.lastNode.pos.y != n.pos.y {
			count++
		}
		if n.lastNode.pos.z != n.pos.z {
			count++
		}
		switch count {
		case 1:
			return 10 // 最小單位
		case 2:
			return 14 // (1^2+1^2)開根號近似值
		case 3:
			return 17 // ((1^2+1^2)^2)+(1^2)開根號近似值
		}
		return 0
	*/
	return 10
}

func (n *node) getGuessCost(end pos) uint16 {
	return 350 * (simpleAbs(end.x-n.pos.x) + simpleAbs(end.y-n.pos.y) + simpleAbs(end.z-n.pos.z))
}

func (n *node) returnNodes(nodes []*node) []*node {
	if n.lastNode != nil {
		nodes = append(nodes, n.lastNode)
		nodes = n.lastNode.returnNodes(nodes)
	} else {
		return nodes
	}
	return nodes
}

func simpleAbs(n int) uint16 {
	if n > 0 {
		return uint16(n)
	}
	return uint16(-n)
}
