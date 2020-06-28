package api

import (
	"sync"
)

type inventory struct {
	itemStacks [46]ItemStack // 背包總共有46格
	lock       sync.Mutex    // 同步鎖
}
type ItemStack struct {
	id    uint32                 // 物品ID
	count int                    //物品數量
	nbt   map[string]interface{} //物品NBT標籤
}

func NewInventory() *inventory {
	return &inventory{itemStacks: [46]ItemStack{}}
}
func (inv *inventory) GetSlotItemStack(slot int) (item *ItemStack) {
	// 格數定義可以在 https://wiki.vg/Inventory#Player_Inventory 找到
	if slot < 0 || slot > 45 {
		return &ItemStack{}
	}
	if inv == nil {
		return &ItemStack{}
	}
	inv.lock.Lock()
	item = &inv.itemStacks[slot]
	inv.lock.Unlock()
	return
}
func (stack *ItemStack) GetID() uint32 {
	if stack == nil {
		return 0
	}
	return stack.id
}
func (stack *ItemStack) SetID(id uint32) {
	if stack == nil {
		return
	}
	stack.id = id
}
func (stack *ItemStack) GetCount() int {
	if stack == nil {
		return 0
	}
	return stack.count
}
func (stack *ItemStack) GetNBT() map[string]interface{} {
	if stack == nil {
		return map[string]interface{}{}
	}
	return stack.nbt
}
