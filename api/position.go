package api

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
	x = pos.x
	return
}
func (pos *Position) SetX(x float64) {
	if pos == nil {
		return
	}
	pos.x = x
}
func (pos *Position) GetY() (y float64) {
	if pos == nil {
		return 0
	}
	y = pos.y
	return
}
func (pos *Position) SetY(y float64) {
	if pos == nil {
		return
	}
	pos.y = y
}
func (pos *Position) GetZ() (z float64) {
	if pos == nil {
		return 0
	}
	z = pos.z
	return
}
func (pos *Position) SetZ(z float64) {
	if pos == nil {
		return
	}
	pos.z = z
}
func (pos *Position) GetYaw() (yaw float32) {
	if pos == nil {
		return 0
	}
	yaw = pos.yaw
	return
}
func (pos *Position) SetYaw(yaw float32) {
	if pos == nil {
		return
	}
	pos.yaw = yaw
}
func (pos *Position) GetPitch() (pitch float32) {
	if pos == nil {
		return 0
	}
	pitch = pos.pitch
	return
}
func (pos *Position) SetPitch(pitch float32) {
	if pos == nil {
		return
	}
	pos.pitch = pitch
}
func (pos *Position) GetOnGround() (onGround bool) {
	if pos == nil {
		return false
	}
	onGround = pos.onGround
	return
}
func (pos *Position) SetOnGround(onGround bool) {
	if pos == nil {
		return
	}
	pos.onGround = onGround
}
