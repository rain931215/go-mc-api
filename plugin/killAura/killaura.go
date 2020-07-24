package killaura

import (
	"github.com/rain931215/go-mc-api/api"
	tpscounter "github.com/rain931215/go-mc-api/plugin/tpsCounter"
)

/*
	Usage
	killaura := killaura.New(c, 600, 5)
	killaura.EntityType = []int32{23, 84, 87, 88, 90} //raid's mobs
	killaura.Start()

*/

//Killaura _
type Killaura struct {
	c                      *api.Client
	tpsCounter             *tpscounter.TpsCounter
	lock, stop             bool
	attackTimes, maxTarget int
	EntityType             []int32
	Delay                  int
}

//New plugin
func New(c *api.Client, Delay, maxTarget int) *Killaura {
	p := new(Killaura)
	p.c = c
	p.tpsCounter = tpscounter.New(c, 180)
	p.Delay = Delay
	p.maxTarget = maxTarget
	p.stop = true
	p.c.Event.AddEventHandler(p.onTimeUpdate, "time")
	return p
}

//Start killaura
func (p *Killaura) Start() {
	p.stop = false
}

//Stop killaura
func (p *Killaura) Stop() {
	p.attackTimes = 0
	p.stop = true
}

//SetDelay 設定攻擊間隔
func (p *Killaura) SetDelay(Delay int) {
	p.Delay = Delay
}

//SetMaxTarget 一次最多攻擊多少對象
func (p *Killaura) SetMaxTarget(maxTarget int) {
	p.maxTarget = maxTarget
}
func (p *Killaura) onTimeUpdate(age, timeOfDay int64) bool {
	if p.stop {
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
					p.tpsCounter.Sleep(p.Delay)
				}
				p.lock = false
			}()
		}
	}
	return false
}
func (p *Killaura) attack() {
	list := p.getAttackList()
	if len(list) > 0 {
		p.c.SwingArm(api.MainHand)
	}
	for i := 0; i < len(list); i++ {
		p.c.AttackEntity(list[i])
	}
}
func (p *Killaura) getAttackList() []int32 {
	result := []int32{}
	list := p.c.EntityList.GetAllEntities()
	for i := 0; i < len(list); i++ {
		if p.checkType(list[i]) {
			if list[i].GetSquaredDistanceToClient(p.c) < 25 {
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
