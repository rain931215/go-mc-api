package nbt

import (
	"errors"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Nbt struct {
	nbt    nbt
	values []interface{}
}

type nbt struct {
	Nbt []interface{}
}

func NewNbt() *Nbt { return &Nbt{} }

func (n *Nbt) AddNewValue(name string, value interface{}) {
	v := &Value{Name: name, Value: value}
	n.values = append(n.values, v)
}
func (n *Nbt) AddValue(value *Value) {
	n.values = append(n.values, value)
}
func (n *Nbt) AddCompoundTag(CompoundTag *CompoundTag) {
	n.AddNewValue(CompoundTag.Name, CompoundTag)
}
func (n *Nbt) AddListTag(ListTag *ListTag) {
	n.AddNewValue(ListTag.Name, ListTag)
}
func (nbt *Nbt) ToJson() ([]byte, error) {
	for i := 0; i < len(nbt.values); i++ {
		nbt.nbt.writeValue(nbt.values[i])
	}
	return json.Marshal(nbt.nbt)
}

func (nbt *nbt) writeValue(v interface{}) {
	switch v := v.(type) {
	case *Value:
		newMap := make(map[string]interface{})
		t, value := v.getType()
		newMap["tagType"] = t
		newMap["name"] = v.Name
		newMap["value"] = value
		nbt.Nbt = append(nbt.Nbt, newMap)
		break
	case *CompoundTag:
		newMap := make(map[string]interface{})
		newMap["tagType"] = 10
		newMap["name"] = v.Name
		newMap["value"] = v.toJSON()
		nbt.Nbt = append(nbt.Nbt, newMap)
		break
	case *ListTag:
		newMap := make(map[string]interface{})
		newMap["tagType"] = 9
		newMap["name"] = v.Name
		newMap["value"] = v.toJSON()
		nbt.Nbt = append(nbt.Nbt, newMap)
		break
	default:
		panic(errors.New("Unknown type"))
	}
}

func (nbt *nbt) toJSON() ([]byte, error) {
	return json.Marshal(nbt)
}
