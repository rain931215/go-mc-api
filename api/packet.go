package api

import (
	"bytes"
	"errors"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data"
	pk "github.com/Tnze/go-mc/net/packet"
)

func (c *Client) handlePacket(p *pk.Packet) error {
	if c.Event.chatHandlers != nil && len(c.Event.chatHandlers) >= 1 {
		for _, v := range c.Event.packetHandlers {
			if v == nil {
				continue
			}
			pass, err := v(*p)
			if err != nil {
				return errors.New("Packet event error" + err.Error())
			}
			if pass {
				return nil
			}
		}
	}
	switch p.ID {
	case data.ChatMessageClientbound:
		var (
			msg chat.Message
		)
		if msg.Decode(bytes.NewReader(p.Data)) == nil {
			if c.Event.chatHandlers == nil || len(c.Event.chatHandlers) < 1 {
				break
			}
			for _, v := range c.Event.chatHandlers {
				if v == nil {
					continue
				}
				pass, err := v(msg)
				if err != nil {
					return errors.New("Chat event error" + err.Error())
				}
				if pass {
					break
				}
			}
		}
		break
	case data.Title:
		var (
			r      = bytes.NewReader(p.Data)
			action pk.VarInt
			msg    chat.Message
		)
		if action.Decode(r) == nil && action == 0 {
			if msg.Decode(r) == nil {
				title := &chat.Message{Text: "[Title] "}
				title.Append(msg)
				if c.Event.titleHandlers == nil || len(c.Event.titleHandlers) < 1 {
					break
				}
				for _, v := range c.Event.titleHandlers {
					if v == nil {
						continue
					}
					pass, err := v(*title)
					if err != nil {
						return errors.New("Chat event error" + err.Error())
					}
					if pass {
						break
					}
				}
			}
		}
		break
	default:
		break
	}
	return nil
}
