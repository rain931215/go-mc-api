package api

import (
	"github.com/Tnze/go-mc/data"
	pk "github.com/Tnze/go-mc/net/packet"
)

func (c *Client) Chat(msg string) {
	c.SendPacket(pk.Marshal(data.ChatMessageServerbound, pk.String(msg)))
}
func (c *Client) ToggleFly(enable bool) {
	b := pk.Byte(4)
	if enable {
		b = pk.Byte(2)
	}
	c.SendPacket(pk.Marshal(
		data.PlayerAbilitiesServerbound,
		b,
		pk.Float(1),
		pk.Float(1),
	))
}
func (c *Client) Move(x, y, z float64, onGround bool) {
	c.SetX(x)
	c.SetY(y)
	c.SetZ(z)
	c.SetOnGround(onGround)
	c.SendPacket(pk.Marshal(
		data.PlayerPosition,
		pk.Double(x),
		pk.Double(y),
		pk.Double(z),
		pk.Boolean(onGround),
	))
}
func (c *Client) Rotation(yaw, pitch float32, onGround bool) {
	c.SetYaw(yaw)
	c.SetPitch(pitch)
	c.SetOnGround(onGround)
	c.SendPacket(pk.Marshal(
		data.PlayerLook,
		pk.Float(yaw),
		pk.Float(pitch),
		pk.Boolean(onGround),
	))
}
func (c *Client) MoveAndRotation(x, y, z float64, yaw, pitch float32, onGround bool) {
	c.SetX(x)
	c.SetY(y)
	c.SetZ(z)
	c.SetYaw(yaw)
	c.SetPitch(pitch)
	c.SetOnGround(onGround)
	c.SendPacket(pk.Marshal(
		data.PlayerPositionAndLookServerbound,
		pk.Double(x),
		pk.Double(y),
		pk.Double(z),
		pk.Float(yaw),
		pk.Float(pitch),
		pk.Boolean(onGround),
	))
}
