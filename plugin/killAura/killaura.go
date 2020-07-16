package killaura

import (
	"github.com/rain931215/go-mc-api/api"
	tpscounter "github.com/rain931215/go-mc-api/plugin/tpsCounter"
)

//Killaura _
type Killaura struct {
	c                      *api.Client
	tpsCounter             *tpscounter.TpsCounter
	lock, stop             bool
	attackTimes, maxTarget int
	EntityType             []int32
	Delay                  uint
}

//New _
func New(c *api.Client, maxTarget int) *Killaura {
	p := new(Killaura)
	p.c = c
	p.tpsCounter = tpscounter.New(c, 180)
	p.maxTarget = maxTarget
	p.stop = true
	p.c.Event.AddEventHandler(p.onTimeUpdate, "time")
	return p
}

//Start _
func (p *Killaura) Start() {
	p.stop = false
}

//Stop _
func (p *Killaura) Stop() {
	p.stop = true
}

func (p *Killaura) onTimeUpdate(age, timeOfDay int64) bool {
	if p.stop {
		p.attackTimes = 0
		return false
	}
	if p.attackTimes < 3 {
		p.attackTimes += 10
	} else {
		if !p.lock {
			p.lock = true
			go func() {
				for _ = p.attackTimes; p.attackTimes > 0; p.attackTimes-- {
					p.attack()
					p.tpsCounter.Sleep(600)
				}
				p.lock = false
			}()
		}
	}
	return false
}
func (p *Killaura) attack() {
	list := p.getAttackList()
	for i := 0; i < len(list); i++ {
		p.c.AttackEntity(list[i])
	}
}
func (p *Killaura) getAttackList() []int32 {
	result := make([]int32, 1)
	list := p.c.EntityList.GetAllEntities()
	for i := 0; i < len(list); i++ {
		if p.checkType(list[i]) {
			if list[i].GetSquaredDistanceToClient(p.c) < 36 {
				result = append(result, list[i].GetID())
				if len(result) >= p.maxTarget {
					return result
				}
			}
		}
	}
	return result
}
func (p *Killaura) checkType(e *api.BaseEntity) bool {
	t := e.GetType()
	for i := 0; i < len(p.EntityType); i++ {
		if t == p.EntityType[i] {
			return true
		}
	}
	return false
}
