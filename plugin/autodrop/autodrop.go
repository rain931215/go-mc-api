package autodrop

import (
	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/rain931215/go-mc-api/api"
)

/* Usage
autodrop := autodrop.New(c)
whiteList := []int32{836} //836為不死圖騰
for i := 9; i < 36; i++ {
	autodrop.OpenSlot[i] = true
	autodrop.SetWhiteListBySlot(int16(i), whiteList)
}
autodrop.Start()
*/

//Autodrop _
type Autodrop struct {
	c         *api.Client
	start     bool
	OpenSlot  [46]bool
	whiteList [46][]int32
}

//New _
func New(c *api.Client) *Autodrop {
	p := new(Autodrop)
	p.c = c
	p.c.Event.AddEventHandler(p.onSetSlotEvent, "onSetSlot")
	return p
}

func (p *Autodrop) Start() {
	p.start = true
}
func (p *Autodrop) Stop() {
	p.start = false
}
func (p *Autodrop) SetAllSlotOpen() {
	for i := 0; i < len(p.OpenSlot); i++ {
		p.OpenSlot[i] = true
	}
}
func (p *Autodrop) SetAllSlotClose() {
	for i := 0; i < len(p.OpenSlot); i++ {
		p.OpenSlot[i] = false
	}
}
func (p *Autodrop) SetWhiteListBySlot(slot int16, items []int32) {
	if slot < 0 || slot > 46 {
		return
	}
	p.whiteList[slot] = items
}

func (p *Autodrop) onSetSlotEvent(id int8, slot int16, data entity.Slot) bool {
	if id != 0 && p.start {
		return false
	}
	if data.ItemID == 0 {
		return false
	}
	if slot < 0 || slot > 46 {
		return false
	}
	if p.OpenSlot[slot] {
		for i := 0; i < len(p.whiteList[slot]); i++ {
			if data.ItemID == p.whiteList[slot][i] {
				return false
			}
		}
		//println("throw", slot, data.ItemID)
		p.c.ClickWindow(0, slot, 1, 4)
	}
	return false
}
