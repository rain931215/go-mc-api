package navigate

import (
	"log"
	"math"
	"time"

	"github.com/rain931215/go-mc-api/api"
)

//Navigate _
type Navigate struct {
	c *api.Client
}

//New _
func New(c *api.Client) *Navigate {
	p := new(Navigate)
	p.c = c
	return p
}

//Move _
func (p *Navigate) Move(x, y, z float64) {
	originalX := math.Floor(p.c.GetX()) + 0.5
	originalY := math.Floor(p.c.GetY())
	originalZ := math.Floor(p.c.GetZ()) + 0.5
	p.c.Move(originalX, originalY, originalZ, true)
	f := setNewPath(x, y, z, p.c)
	nodes := f.getNodes()

	for i := len(nodes) - 1; i > 0; i-- {

		dx := (originalX + float64(nodes[i].pos.x))
		dy := (originalY + float64(nodes[i].pos.y))
		dz := (originalZ + float64(nodes[i].pos.z))
		log.Println(dx, dy, dz)
		p.c.Move(dx, dy, dz, true)
		time.Sleep(300 * time.Millisecond)
	}
}
