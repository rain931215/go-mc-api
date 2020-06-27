package api

import (
	"github.com/Tnze/go-mc/chat"
	pk "github.com/Tnze/go-mc/net/packet"
)

type Events struct {
	packetHandlers     []func(p pk.Packet) (bool, error)
	chatHandlers       []func(msg chat.Message) (bool, error)
	titleHandlers      []func(msg chat.Message) (bool, error)
	disconnectHandlers []func(msg chat.Message) (bool, error)
	dieHandlers        []func() error
}

func (e *Events) AddEventHandler(handler interface{}, handlerType string) {
	switch handler.(type) {
	case func(p pk.Packet) (bool, error):
		e.packetHandlers = append(e.packetHandlers, handler.(func(p pk.Packet) (bool, error)))
	case func(msg chat.Message) (bool, error):
		switch handlerType {
		case "chat":
			e.chatHandlers = append(e.chatHandlers, handler.(func(msg chat.Message) (bool, error)))
			break
		case "title":
			e.titleHandlers = append(e.titleHandlers, handler.(func(msg chat.Message) (bool, error)))
			break
		case "disconnect":
			e.disconnectHandlers = append(e.disconnectHandlers, handler.(func(msg chat.Message) (bool, error)))
			break
		default:
			panic("Unknown handler on type [func(msg chat.Message) error]")
		}
		break
	case func() error:
		switch handlerType {
		case "die":
			e.dieHandlers = append(e.dieHandlers, handler.(func() error))
			break
		default:
			panic("Unknown handler on type [func() error]")
		}
		break
	default:
		panic("Unknown handler type")
	}
}
