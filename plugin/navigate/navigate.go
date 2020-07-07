package navigate

import (
	"math"
	"strconv"
	"time"

	"github.com/rain931215/go-mc-api/api"
	"github.com/rain931215/go-mc-api/plugin/mcfalloutcmd"
)

//Navigate _
type Navigate struct {
	c *api.Client
}

//New _
func New(cmdHandler *mcfalloutcmd.McfalloutCmd) *Navigate {
	p := new(Navigate)
	p.c = cmdHandler.Client
	cmdHandler.AddCmd("move", func(c *api.Client, sender string, text string, args []string) {
		if len(args) != 3 {
			return
		}
		x, err := strconv.ParseFloat(args[0], 64)
		y, err := strconv.ParseFloat(args[1], 64)
		z, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			return
		}
		p.Move(x, y, z)
	})

	cmdHandler.AddCmd("moveto", func(c *api.Client, sender string, text string, args []string) {
		if len(args) != 3 {
			return
		}
		x, err := strconv.ParseFloat(args[0], 64)
		y, err := strconv.ParseFloat(args[1], 64)
		z, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			return
		}
		go func() {
			p.MoveTo(x, y, z)
		}()
	})
	return p
}

//Move _
func (p *Navigate) Move(x, y, z float64) {
	p.MoveTo(p.c.GetX()+x, p.c.GetY()+y, p.c.GetZ()+z)
}

//MoveTo _
func (p *Navigate) MoveTo(x, y, z float64) {
	originalX := math.Floor(p.c.GetX()) + 0.5
	originalY := math.Floor(p.c.GetY())
	originalZ := math.Floor(p.c.GetZ()) + 0.5
	p.c.Move(originalX, originalY, originalZ, false)
	f := setNewPath(x, y, z, p.c)
	t := time.Now().UnixNano()
	nodes := sortNodes(f.getNodes())
	println((time.Now().UnixNano() - t) / 1000000)
	for i := len(nodes) - 1; i != 0; i-- {
		dx := originalX + float64(nodes[i].pos.x)
		dy := originalY + float64(nodes[i].pos.y)
		dz := originalZ + float64(nodes[i].pos.z)
		//log.Println(dx, dy, dz)
		p.c.Move(dx, dy, dz, false)
		time.Sleep(30 * time.Millisecond)
	}
}

func sortNodes(nodes []*node) []*node {
	var (
		result           []*node
		count, stepCount uint8
	)
	if len(nodes) < 2 {
		return nodes
	}
	result = append(result, nodes[0])
	result = append(result, nodes[1])
	for i := 1; i < len(nodes)-1; i++ {
		count = 0
		if nodes[i-1].pos.x != nodes[i+1].pos.x {
			count++
		}
		if nodes[i-1].pos.y != nodes[i+1].pos.y {
			count++
		}
		if nodes[i-1].pos.z != nodes[i+1].pos.z {
			count++
		}
		if count != 1 || stepCount > 8 {
			result = append(result, nodes[i])
			stepCount = 0
		} else {
			stepCount++
		}
	}
	result = append(result, nodes[len(nodes)-1])
	return result
}
