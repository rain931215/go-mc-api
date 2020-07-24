package api

import (
	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/chat"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/rain931215/go-mc-api/api/world"
)

type Events struct {
	blockChangeHandlers []func(x, y, z int, id world.BlockStatus) bool
	packetHandlers      []func(p *pk.Packet) bool
	setSlotHandlers     []func(id int8, slot int16, data entity.Slot) bool
	chatHandlers        []func(msg chat.Message) bool
	disconnectHandlers  []func(msg chat.Message) bool
	timeUpdateHandlers  []func(age, timeOfDay int64) bool
	dieHandlers         []func() bool
}

func (e *Events) AddEventHandler(handler interface{}, handlerType string) {
	switch handler.(type) {
	case func(x, y, z int, id world.BlockStatus) bool:
		e.blockChangeHandlers = append(e.blockChangeHandlers, handler.(func(x, y, z int, id world.BlockStatus) bool))
		break
	case func(age, timeOfDay int64) bool:
		e.timeUpdateHandlers = append(e.timeUpdateHandlers, handler.(func(age, timeOfDay int64) bool))
		break
	case func(id int8, slot int16, data entity.Slot) bool:
		e.setSlotHandlers = append(e.setSlotHandlers, handler.(func(id int8, slot int16, data entity.Slot) bool))
		break
	case func(p *pk.Packet) bool:
		e.packetHandlers = append(e.packetHandlers, handler.(func(p *pk.Packet) bool))
		break
	case func(msg chat.Message) bool:
		switch handlerType {
		case "chat":
			e.chatHandlers = append(e.chatHandlers, handler.(func(msg chat.Message) bool))
			break
		case "disconnect":
			e.disconnectHandlers = append(e.disconnectHandlers, handler.(func(msg chat.Message) bool))
			break
		default:
			panic("Unknown handler on type [func(msg chat.Message) bool]")
		}
		break
	case func() bool:
		switch handlerType {
		case "die":
			e.dieHandlers = append(e.dieHandlers, handler.(func() bool))
			break
		default:
			panic("Unknown handler on type [func() bool]")
		}
		break
	default:
		panic("Unknown handler type")
	}
}
