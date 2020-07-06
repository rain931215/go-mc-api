package api

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data"
	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/google/uuid"
	"github.com/rain931215/go-mc-api/api/world"
)

func (c *Client) handlePacket(p *pk.Packet) error {
	if p == nil {
		return nil
	}
	go func() {
		//TODO (Async Events)
		if c.Event.packetHandlers != nil && len(c.Event.packetHandlers) >= 1 {
			for _, v := range c.Event.packetHandlers {
				if v == nil {
					continue
				}
				pass, err := v(p)
				if err != nil || pass {
					break
				}
			}
		}
	}()
	switch p.ID {
	case data.ChatMessageClientbound:
		go c.handleChatPacket(p)
		break
	case data.Title:
		go c.handleTitlePacket(p)
		break
	case data.BlockChange:
		go c.handleBlockChangePacket(p)
		break
	case data.MultiBlockChange:
		go c.handleMultiBlockChangePacket(p)
		break
	case data.PlayerPositionAndLookClientbound:
		go c.handleMoveAndRotationPacket(p)
		break
	case data.ChunkData:
		go c.handleLoadChunkPacket(p)
		break
	case data.SetSlot:
		go c.handleSetSlotPacket(p)
		break
	case data.TimeUpdate:
		go c.handleTimeUpdatePacket(p)
		break
	case data.SpawnMob:
		go c.handleSpawnMobPacket(p)
		break
	case data.EntityRelativeMove, data.EntityLookAndRelativeMove:
		go c.handleEntityLocationUpdatePacket(p)
		break
	case data.EntityTeleport:
		go c.handleEntityTeLePortPacket(p)
		break
	case data.DestroyEntities:
		go c.handleRemoveEntityPacket(p)
		break
	case data.UnloadChunk:
		go c.handleUnlockChunk(p)
		break
	}
	return nil
}
func (c *Client) handleSetSlotPacket(p *pk.Packet) error {
	var (
		windowID pk.Byte
		slot     pk.Short
		slotData entity.Slot
	)
	if err := ScanFields(p, &windowID, &slot, &slotData); err != nil && !errors.Is(err, nbt.ErrEND) {
		return err
	}
	if c.Inventory != nil && windowID == 0 { // 更新背包
		c.Inventory.lock.Lock()
		c.Inventory.itemStacks[slot] = ItemStack{id: uint32(slotData.ItemID), count: int(slotData.Count), nbt: nil} //TODO(Need improve nbt)
		c.Inventory.lock.Unlock()
	}
	if c.Event.setSlotHandlers == nil || len(c.Event.setSlotHandlers) < 1 { // 如果沒有任何handler的話就跳過解析
		return nil
	}
	for _, v := range c.Event.setSlotHandlers {
		if v == nil {
			continue
		}
		pass, err := v(int8(windowID), int16(slot), slotData)
		if err != nil {
			return errors.New("Set Slot event error" + err.Error())
		}
		if pass {
			break
		}
	}
	return nil
}
func (c *Client) handleTimeUpdatePacket(p *pk.Packet) error {
	if c.Event.timeUpdateHandlers == nil || len(c.Event.timeUpdateHandlers) < 1 { // 如果沒有任何handler的話就跳過解析
		return nil
	}
	var age, timeOfDay pk.Long
	if err := ScanFields(p, &age, &timeOfDay); err != nil {
		return err
	}
	for _, v := range c.Event.timeUpdateHandlers {
		if v == nil {
			continue
		}
		pass, err := v(int64(age), int64(timeOfDay))
		if err != nil {
			return errors.New("Time Update event error" + err.Error())
		}
		if pass {
			break
		}
	}
	return nil
}
func (c *Client) handleChatPacket(p *pk.Packet) error {
	if c.Event.chatHandlers == nil || len(c.Event.chatHandlers) < 1 { // 如果沒有任何handler的話就跳過解析
		return nil
	}
	var msg chat.Message
	if err := ScanFields(p, &msg); err == nil {
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
	return nil
}
func (c *Client) handleTitlePacket(p *pk.Packet) error {
	if c.Event.titleHandlers == nil || len(c.Event.titleHandlers) < 1 { // 如果沒有任何handler的話就跳過解析
		return nil
	}
	var (
		action pk.VarInt
		msg    chat.Message
	)
	if ScanFields(p, &action, &msg) == nil && action == 0 {
		title := &chat.Message{Text: "[Title] "}
		title.Append(msg)
		for _, v := range c.Event.titleHandlers {
			if v == nil {
				continue
			}
			pass, err := v(*title)
			if err != nil {
				return errors.New("Title Message event error" + err.Error())
			}
			if pass {
				break
			}
		}
	}
	return nil
}
func (c *Client) handleBlockChangePacket(p *pk.Packet) error {
	if c.World == nil {
		return nil
	}
	var (
		pos pk.Position
		id  pk.VarInt
	)
	if err := ScanFields(p, &pos, &id); err != nil {
		return err
	}
	c.World.ChunkMapLock.Lock()
	if chunk := c.World.Chunks[world.ChunkLoc{X: pos.X >> 4, Z: pos.Z >> 4}]; chunk != nil {
		if v := chunk.Sections[pos.Y/16]; v != nil {
			v.SetBlock(world.SectionOffset(pos.X&15, pos.Y&15, pos.Z&15), world.BlockStatus(id))
		}
	}
	c.World.ChunkMapLock.Unlock()
	return nil
}
func (c *Client) handleMultiBlockChangePacket(p *pk.Packet) error {
	if c.World == nil {
		return nil
	}
	var (
		r      = bufio.NewReader(bytes.NewReader(p.Data))
		cX, cY pk.Int
		count  pk.VarInt
	)
	if err := cX.Decode(r); err != nil {
		return err
	}
	if err := cY.Decode(r); err == nil {
		return err
	}
	if err := count.Decode(r); err == nil {
		return err
	}
	c.World.ChunkMapLock.Lock()
	if chunk := c.World.Chunks[world.ChunkLoc{X: int(cX), Z: int(cY)}]; chunk != nil {
		for i := 0; i < int(count); i++ {
			if xz, err := r.ReadByte(); err == nil {
				if y, err := r.ReadByte(); err == nil {
					var blockID pk.VarInt
					if blockID.Decode(r) == nil {
						x, z := xz>>4, xz&0x0F
						if v := chunk.Sections[y/16]; v != nil {
							v.SetBlock(world.SectionOffset(int(x), int(y%16), int(z)), world.BlockStatus(blockID))
						}
					}
				}
			}
		}
	}
	c.World.ChunkMapLock.Unlock()
	return nil
}
func (c *Client) handleMoveAndRotationPacket(p *pk.Packet) error {
	var (
		x, y, z    pk.Double
		yaw, pitch pk.Float
		flags      pk.Byte
		TpID       pk.VarInt
	)
	if err := ScanFields(p, &x, &y, &z, &yaw, &pitch, &flags, &TpID); err != nil {
		return err
	}
	c.SendPacket(pk.Marshal(
		data.TeleportConfirm,
		TpID,
	))
	if flags&0x01 == 0 {
		c.SetX(float64(x))
	} else {
		c.SetX(c.GetX() + float64(x))
	}
	if flags&0x02 == 0 {
		c.SetY(float64(y))
	} else {
		c.SetY(c.GetY() + float64(y))
	}
	if flags&0x04 == 0 {
		c.SetZ(float64(z))
	} else {
		c.SetZ(c.GetZ() + float64(z))
	}
	if flags&0x08 == 0 {
		c.SetYaw(float32(yaw))
	} else {
		c.SetYaw(c.GetYaw() + float32(yaw))
	}
	if flags&0x10 == 0 {
		c.SetPitch(float32(pitch))
	} else {
		c.SetPitch(c.GetPitch() + float32(pitch))
	}
	return nil
}
func (c *Client) handleSpawnMobPacket(p *pk.Packet) error {
	if c.EntityList == nil || c.EntityList.hashMap == nil {
		return nil
	}
	var (
		eID     pk.VarInt
		eUUID   pk.UUID
		eType   pk.VarInt
		x, y, z pk.Double
	)
	if err := ScanFields(p, &eID, &eUUID, &eType, &x, &y, &z); err != nil {
		return err
	}
	newEntity := new(BaseEntity)
	newEntity.eID = int32(eID)
	newEntity.eType = int32(eType)
	newEntity.eUUID = uuid.UUID(eUUID)
	newEntity.eX = float64(x)
	newEntity.eY = float64(y)
	newEntity.eZ = float64(z)
	c.EntityList.hashMap.Set(int32(eID), newEntity)
	return nil
}
func (c *Client) handleEntityLocationUpdatePacket(p *pk.Packet) error {
	if c.EntityList == nil || c.EntityList.hashMap == nil {
		return nil
	}
	var (
		eID     pk.VarInt
		x, y, z pk.Short
	)
	if err := ScanFields(p, &eID, &x, &y, &z); err != nil {
		return err
	}
	if element, ok := c.EntityList.hashMap.Get(int32(eID)); ok {
		if value, ok := element.(*BaseEntity); ok {
			value.eX = (float64(x)/128 + value.eX*32) / 32
			value.eY = (float64(y)/128 + value.eY*32) / 32
			value.eZ = (float64(z)/128 + value.eZ*32) / 32
		}
	}
	return nil
}
func (c *Client) handleEntityTeLePortPacket(p *pk.Packet) error {
	if c.EntityList == nil || c.EntityList.hashMap == nil {
		return nil
	}
	var (
		eID     pk.VarInt
		x, y, z pk.Double
	)
	if err := ScanFields(p, &eID, &x, &y, &z); err != nil {
		return err
	}
	if element, ok := c.EntityList.hashMap.Get(int32(eID)); ok {
		if value, ok := element.(*BaseEntity); ok {
			value.eX = float64(x)
			value.eY = float64(y)
			value.eZ = float64(z)
		}
	}
	return nil
}
func (c *Client) handleRemoveEntityPacket(p *pk.Packet) error {
	if c.EntityList == nil || c.EntityList.hashMap == nil {
		return nil
	}
	var (
		r     = bufio.NewReader(bytes.NewReader(p.Data))
		count pk.VarInt
	)
	if err := count.Decode(r); err != nil {
		return err
	}
	for i := 0; i < int(count); i++ {
		var entityID pk.VarInt
		if entityID.Decode(r) == nil {
			c.EntityList.hashMap.Del(int32(entityID))
		}
	}
	return nil
}
func (c *Client) handleUnlockChunk(p *pk.Packet) error {
	if c.World == nil {
		return nil
	}
	var cX, cZ pk.Int
	err := ScanFields(p, &cX, &cZ)
	if err == nil {
		c.World.ChunkMapLock.Lock()
		if _, ok := c.World.Chunks[world.ChunkLoc{X: int(cX), Z: int(cZ)}]; ok {
			c.World.Chunks[world.ChunkLoc{X: int(cX), Z: int(cZ)}] = nil
		}
		c.World.ChunkMapLock.Unlock()
	}
	return nil
}
func (c *Client) handleLoadChunkPacket(p *pk.Packet) error {
	if c.World == nil {
		return nil
	}
	var (
		X, Z           pk.Int
		FullChunk      pk.Boolean
		PrimaryBitMask pk.VarInt
		Heightmaps     struct{}
		Biomes         = biomesData{fullChunk: (*bool)(&FullChunk)}
		Data           chunkData
	)
	if err := p.Scan(&X, &Z, &FullChunk, &PrimaryBitMask, pk.NBT{V: &Heightmaps}, &Biomes, &Data); err != nil {
		return err
	}
	chunk, err := world.DecodeChunkColumn(int32(PrimaryBitMask), Data)
	if err != nil {
		return fmt.Errorf("decode chunk column fail: %w", err)
	}
	c.World.LoadChunk(int(X), int(Z), chunk)
	return nil
}

type biomesData struct {
	fullChunk *bool
	data      [1024]int32
}

func (b *biomesData) Decode(r pk.DecodeReader) error {
	if b.fullChunk == nil || !*b.fullChunk {
		return nil
	}
	for i := range b.data {
		err := (*pk.Int)(&b.data[i]).Decode(r)
		if err != nil {
			return err
		}
	}
	return nil
}

type chunkData []byte

// Decode implement net.packet.FieldDecoder
func (c *chunkData) Decode(r pk.DecodeReader) error {
	var Size pk.VarInt
	if err := Size.Decode(r); err != nil {
		return err
	}
	*c = make([]byte, Size)
	if _, err := r.Read(*c); err != nil {
		return err
	}
	return nil
}
func ScanFields(p *pk.Packet, fields ...pk.FieldDecoder) error {
	r := bufio.NewReader(bytes.NewReader(p.Data))
	for _, v := range fields {
		err := v.Decode(r)
		if err != nil {
			return err
		}
	}
	return nil
}
