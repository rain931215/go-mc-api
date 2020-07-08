package api

import "github.com/rain931215/go-mc-api/data"

type inventory struct {
	itemStacks [46]ItemStack // 背包總共有46格
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
	item = &inv.itemStacks[slot]
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
func (inv *inventory) GetItem(itemName string) (int, []int) {
	var (
		amount int
		slots  []int
	)
	if inv == nil {
		return amount, slots
	}
	for i := 9; i < 45; i++ {
		if data.ItemNameByID[inv.itemStacks[i].id] == itemName {
			slots = append(slots, i)
			amount += inv.itemStacks[i].count
		}
	}

	return amount, slots
}
