package api

import (
	"github.com/cornelk/hashmap"
)

type EntityList struct {
	hashMap *hashmap.HashMap
}

type BaseEntity struct {
	eID        int32
	eUUID      string
	eType      int32
	eX, eY, eZ float64
	//eYaw, ePitch                       float32
	//eVelocityX, eVelocityY, eVelocityZ int16
}

func NewEntityList() (list *EntityList) {
	list = new(EntityList)
	list.hashMap = &hashmap.HashMap{}
	return
}
func (list *EntityList) GetEntityByID(entityID int32) *BaseEntity {
	if list == nil || list.hashMap == nil {
		return nil
	}
	if element, ok := list.hashMap.Get(entityID); ok {
		if value, ok := element.(*BaseEntity); ok {
			return value
		}
	}
	return nil
}
func (list *EntityList) GetAllEntities() []*BaseEntity {
	if list == nil || list.hashMap == nil {
		return nil
	}
	var entitiesList []*BaseEntity
	for entity := range list.hashMap.Iter() {
		if value, ok := entity.Value.(*BaseEntity); ok {
			entitiesList = append(entitiesList, value)
		}
	}
	return entitiesList
}
func (list *EntityList) ClearAllEntities() {
	list.hashMap = &hashmap.HashMap{}
}
func (entity *BaseEntity) GetID() (id int32) {
	if entity == nil {
		return 0
	}
	id = entity.eID
	return
}
func (entity *BaseEntity) GetType() (eType int32) {
	if entity == nil {
		return 0
	}
	eType = entity.eType
	return
}
func (entity *BaseEntity) GetUUID() (id string) {
	if entity == nil {
		return ""
	}
	id = entity.eUUID
	return
}
func (entity *BaseEntity) GetX() (x float64) {
	if entity == nil {
		return 0
	}
	x = entity.eX
	return
}
func (entity *BaseEntity) GetY() (y float64) {
	if entity == nil {
		return 0
	}
	y = entity.eY
	return
}
func (entity *BaseEntity) GetZ() (z float64) {
	if entity == nil {
		return 0
	}
	z = entity.eZ
	return
}
func (entity *BaseEntity) GetSquaredDistanceToClient(c *Client) float64 {
	diffX := entity.eX - c.x
	diffY := entity.eY - c.y
	diffZ := entity.eZ - c.z
	return diffX*diffX + diffY*diffY + diffZ*diffZ
}
