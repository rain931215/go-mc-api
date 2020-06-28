package api

import (
	"github.com/cornelk/hashmap"
	"github.com/google/uuid"
	"sync"
)

type EntityList struct {
	hashMap *hashmap.HashMap
}

type BaseEntity struct {
	eID        int32
	eUUID      uuid.UUID
	eType      int32
	eX, eY, eZ float64
	//eYaw, ePitch                       float32
	//eVelocityX, eVelocityY, eVelocityZ int16
	sync.Mutex
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
	entity.Lock()
	id = entity.eID
	entity.Unlock()
	return
}
func (entity *BaseEntity) GetType() (eType int32) {
	if entity == nil {
		return 0
	}
	entity.Lock()
	eType = entity.eType
	entity.Unlock()
	return
}
func (entity *BaseEntity) GetUUID() (id uuid.UUID) {
	if entity == nil {
		return uuid.UUID{}
	}
	entity.Lock()
	id = entity.eUUID
	entity.Unlock()
	return
}
func (entity *BaseEntity) GetX() (x float64) {
	if entity == nil {
		return 0
	}
	entity.Lock()
	x = entity.eX
	entity.Unlock()
	return
}
func (entity *BaseEntity) GetY() (y float64) {
	if entity == nil {
		return 0
	}
	entity.Lock()
	y = entity.eY
	entity.Unlock()
	return
}
func (entity *BaseEntity) GetZ() (z float64) {
	if entity == nil {
		return 0
	}
	entity.Lock()
	z = entity.eZ
	entity.Unlock()
	return
}
