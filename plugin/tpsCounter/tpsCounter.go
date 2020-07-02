package tpscounter

import (
	"time"

	"github.com/rain931215/go-mc-api/api"
)

//TpsCounter _
type TpsCounter struct {
	cycleTime    int
	markRealTime int64 //Second
	packetCount  int
}

//New _
func New(c *api.Client, cycleTime int) *TpsCounter {
	p := new(TpsCounter)
	p.cycleTime = cycleTime
	p.markRealTime = time.Now().Unix()
	c.Event.AddEventHandler(p.onTimeUpdate, "timeUpdate")
	return p
}

func (p *TpsCounter) onTimeUpdate(age, timeOfDay int64) (bool, error) {
	p.packetCount++
	if p.packetCount > (p.cycleTime / 2) {
		p.markRealTime = p.markRealTime + ((time.Now().Unix() - p.markRealTime) / 2)
		p.packetCount = p.packetCount / 2
	}
	return false, nil
}

//GetTps _
func (p *TpsCounter) GetTps() float64 {
	return float64(p.packetCount*20) / float64(time.Now().Unix()-p.markRealTime)
}

//Sleep _
func (p *TpsCounter) Sleep(ms int) {
	time.Sleep(time.Millisecond * time.Duration(float64(ms)*(20/(p.GetTps()))))
}
