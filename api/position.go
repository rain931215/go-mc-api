package api

import "sync"

type Direction int

const (
	Bottom Direction = iota
	Top
	North
	South
	West
	East
)

type Position struct {
	Vec3d
	Vec2f
	onGround bool
	sync.Mutex
}
type Vec3d struct {
	x, y, z float64
}
type Vec2f struct {
	yaw, pitch float32
}

func (pos *Position) GetX() (x float64) {
	if pos == nil {
		return 0
	}
	pos.Lock()
	x = pos.x
	pos.Unlock()
	return
}
func (pos *Position) SetX(x float64) {
	if pos == nil {
		return
	}
	pos.Lock()
	pos.x = x
	pos.Unlock()
}
func (pos *Position) GetY() (y float64) {
	if pos == nil {
		return 0
	}
	pos.Lock()
	y = pos.y
	pos.Unlock()
	return
}
func (pos *Position) SetY(y float64) {
	if pos == nil {
		return
	}
	pos.Lock()
	pos.y = y
	pos.Unlock()
}
func (pos *Position) GetZ() (z float64) {
	if pos == nil {
		return 0
	}
	pos.Lock()
	z = pos.z
	pos.Unlock()
	return
}
func (pos *Position) SetZ(z float64) {
	if pos == nil {
		return
	}
	pos.Lock()
	pos.z = z
	pos.Unlock()
}
func (pos *Position) GetYaw() (yaw float32) {
	if pos == nil {
		return 0
	}
	pos.Lock()
	yaw = pos.yaw
	pos.Unlock()
	return
}
func (pos *Position) SetYaw(yaw float32) {
	if pos == nil {
		return
	}
	pos.Lock()
	pos.yaw = yaw
	pos.Unlock()
}
func (pos *Position) GetPitch() (pitch float32) {
	if pos == nil {
		return 0
	}
	pos.Lock()
	pitch = pos.pitch
	pos.Unlock()
	return
}
func (pos *Position) SetPitch(pitch float32) {
	if pos == nil {
		return
	}
	pos.Lock()
	pos.pitch = pitch
	pos.Unlock()
}
func (pos *Position) GetOnGround() (onGround bool) {
	if pos == nil {
		return false
	}
	pos.Lock()
	onGround = pos.onGround
	pos.Unlock()
	return
}
func (pos *Position) SetOnGround(onGround bool) {
	if pos == nil {
		return
	}
	pos.Lock()
	pos.onGround = onGround
	pos.Unlock()
}
