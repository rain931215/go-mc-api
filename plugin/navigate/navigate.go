package navigate

import (
	"fmt"
	"math"
	"strconv"
	"sync"
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
func (p *Navigate) Move(x, y, z float64) bool {
	return p.MoveTo(p.c.GetX()+x, p.c.GetY()+y, p.c.GetZ()+z)
}

//MoveTo _
func (p *Navigate) MoveTo(x, y, z float64) bool {
	originalX := math.Floor(p.c.GetX()) + 0.5
	originalY := math.Floor(p.c.GetY())
	originalZ := math.Floor(p.c.GetZ()) + 0.5
	p.c.Move(originalX, originalY, originalZ, false)

	finder1 := setNewPath(p.c, p.c.GetX(), p.c.GetY(), p.c.GetZ(), x, y, z)
	finder2 := setNewPath(p.c, x, y, z, p.c.GetX(), p.c.GetY(), p.c.GetZ())
	var (
		successNodes []*node
		wait         sync.WaitGroup
		reverse      bool
	)
	t := time.Now().UnixNano()
	wait.Add(2)
	go func() {
		pass, nodes := finder1.getNodes()
		finder2.stop()
		if pass {
			fmt.Println("正向搜尋完畢 方塊:" + strconv.Itoa(len(nodes)))
			successNodes = nodes
		} else {
			finder2.stop()
		}
		wait.Done()
	}()
	go func() {
		pass, nodes := finder2.getNodes()
		finder1.stop()
		if pass {
			fmt.Println("反向搜尋完畢 方塊:" + strconv.Itoa(len(nodes)))
			for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
				nodes[i], nodes[j] = nodes[j], nodes[i]
			}
			successNodes = nodes
			reverse = true
		}
		wait.Done()
	}()
	wait.Wait()
	println((time.Now().UnixNano() - t) / 1000000)
	if len(successNodes) > 0 {
		fmt.Println("成功尋找到路徑 長度:" + strconv.Itoa(len(successNodes)) + "方塊")
		nodes := sortNodes(successNodes)
		for i := len(nodes) - 1; i != 0; i-- {
			var dx, dy, dz float64
			//log.Println(dx, dy, dz)
			if reverse {
				dx = x + float64(nodes[i].pos.x)
				dy = y + float64(nodes[i].pos.y)
				dz = z + float64(nodes[i].pos.z)
			} else {
				dx = originalX + float64(nodes[i].pos.x)
				dy = originalY + float64(nodes[i].pos.y)
				dz = originalZ + float64(nodes[i].pos.z)
			}
			p.c.Move(dx, dy, dz, false)
			time.Sleep(30 * time.Millisecond)
		}
		return true
	} else {
		fmt.Println("找不到路徑")
		return false
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
