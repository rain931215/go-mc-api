package nbt

type CompoundTag struct {
	Name   string
	Values []*Value
}

func NewCompoundTag(name string) *CompoundTag {
	CompoundTag := new(CompoundTag)
	CompoundTag.Name = name
	return CompoundTag
}

func (tag *CompoundTag) AddNewValue(name string, value interface{}) {
	v := &Value{Name: name, Value: value}
	tag.Values = append(tag.Values, v)
}
func (tag *CompoundTag) AddValue(value *Value) {
	tag.Values = append(tag.Values, value)
}
func (tag *CompoundTag) AddCompoundTag(CompoundTag *CompoundTag) {
	tag.AddValue(&Value{Name: CompoundTag.Name, Value: CompoundTag})
}
func (tag *CompoundTag) AddListTag(ListTag *ListTag) {
	tag.AddValue(&Value{Name: ListTag.Name, Value: ListTag})
}

func (tag *CompoundTag) toJSON() []interface{} {
	var result []interface{}
	for i := 0; i < len(tag.Values); i++ {
		newMap := make(map[string]interface{})
		t, value := tag.Values[i].getType()
		newMap["tagType"] = t
		newMap["name"] = tag.Values[i].Name
		newMap["value"] = value
		result = append(result, newMap)
	}
	return result
}
