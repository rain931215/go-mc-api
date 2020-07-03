package api

import (
	"github.com/Tnze/go-mc/data"
	pk "github.com/Tnze/go-mc/net/packet"
)

type Hand int32

const (
	MainHand Hand = iota
	OffHand
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
func (c *Client) StartBreakBlock(x, y, z int, direction Direction) {
	c.SendPacket(pk.Marshal(
		data.PlayerDigging,
		pk.VarInt(0),
		pk.Position{X: x, Y: y, Z: z},
		pk.Byte(direction),
	))
}
func (c *Client) CancelBreakBlock(x, y, z int, direction Direction) {
	c.SendPacket(pk.Marshal(
		data.PlayerDigging,
		pk.VarInt(1),
		pk.Position{X: x, Y: y, Z: z},
		pk.Byte(direction),
	))
}
func (c *Client) FinishBreakBlock(x, y, z int, direction Direction) {
	c.SendPacket(pk.Marshal(
		data.PlayerDigging,
		pk.VarInt(2),
		pk.Position{X: x, Y: y, Z: z},
		pk.Byte(direction),
	))
}
func (c *Client) AttackEntity(id int32) {
	c.SendPacket(pk.Marshal(
		data.UseEntity,
		pk.VarInt(id),
		pk.VarInt(0),
	))
}
func (c *Client) SwitchHotBar(slot int16) {
	// 接受0~8的格數
	if slot < 0 || slot > 8 {
		panic("switch hot bar error: unknown slot")
	}
	c.SendPacket(pk.Marshal(
		data.HeldItemChangeServerbound,
		pk.Short(slot),
	))
}
func (c *Client) CloseWindow(id uint8) {
	c.SendPacket(pk.Marshal(
		data.CloseWindowServerbound,
		pk.UnsignedByte(id),
	))
}
func (c *Client) PlaceBlock(hand Hand, x, y, z int, face Direction, cursorX, cursorY, cursorZ float32, insideBlock bool) {
	c.SendPacket(pk.Marshal(
		data.PlayerBlockPlacement,
		pk.VarInt(hand),
		pk.Position{X: x, Y: y, Z: z},
		pk.VarInt(face),
		pk.Float(cursorX),
		pk.Float(cursorY),
		pk.Float(cursorZ),
		pk.Boolean(insideBlock),
	))
}
func (c *Client) SwingArm(hand Hand) {
	c.SendPacket(pk.Marshal(
		data.AnimationServerbound,
		pk.VarInt(hand),
	))
}
