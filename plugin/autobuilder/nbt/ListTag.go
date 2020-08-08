package nbt

import "errors"

type ListTag struct {
	Name    string
	TagType int
	Values  []*Value
}

func NewListTag(name string, tagType int) *ListTag { return &ListTag{Name: name, TagType: tagType} }

func (tag *ListTag) AddNewValue(value interface{}) {
	v := &Value{Name: "none", Value: value}
	tag.Values = append(tag.Values, v)
}
func (tag *ListTag) AddValue(value *Value) {
	tag.Values = append(tag.Values, value)
}
func (tag *ListTag) AddCompoundTag(CompoundTag *CompoundTag) {
	tag.AddValue(&Value{Name: CompoundTag.Name, Value: CompoundTag})
}
func (tag *ListTag) AddListTag(ListTag *ListTag) {
	tag.AddValue(&Value{Name: ListTag.Name, Value: ListTag})
}

func (tag *ListTag) toJSON() map[string]interface{} {
	result := make(map[string]interface{})
	list := make([]interface{}, 0)
	for i := 0; i < len(tag.Values); i++ {
		if tag.Values[i].getType() != tag.TagType {
			panic(errors.New("Wrong Type in TagList"))
		}
		list = append(list, tag.Values[i].Value)
	}
	if len(list) > 0 {
		result["tagListType"] = tag.TagType
		result["list"] = list
	}
	return result
}
