package nbt

import (
	"strconv"
)

type ListTag struct {
	name       string
	values     []*Value
	writeSlice []interface{}
}

func NewListTag(name string) *ListTag {
	tag := new(ListTag)
	tag.name = name
	return tag
}

func (t *ListTag) AddValue(name string, value *Value) {
	t.values = append(t.values, value)
}

func (c *ListTag) writeListTag() {
	c.writeSlice = make([]interface{}, 0)
	for i := 0; i < len(c.values); i++ {
		valueName := c.values[i].name
		switch v := c.values[i].value.(type) {
		case byte:
			c.writeValue(valueName, v, 1)
		case int16:
			c.writeValue(valueName, v, 2)
		case int32:
			c.writeValue(valueName, v, 3)
		case int64:
			s := strconv.FormatInt(v, 10)
			c.writeValue(valueName, s, 4)
		case float32:
			c.writeValue(valueName, v, 5)
		case float64:
			c.writeValue(valueName, v, 6)
		case []byte:
			c.writeValue(valueName, v, 7)
		case string:
			c.writeValue(valueName, v, 8)
		case []struct{}:
			c.writeValue(valueName, v, 9)
		case *CompoundTag:
			tag := c.values[i].value.(*CompoundTag)
			tag.writeCompoundTag()
			newMap := make(map[string]interface{})
			newMap["tagType"] = 10
			newMap["name"] = tag.name
			newMap["value"] = tag.writeMap
			c.writeSlice = append(c.writeSlice, newMap)
		case []int:
			c.writeValue(valueName, v, 11)
		case []int64:
			strs := []string{}
			for i := 0; i < len(v); i++ {
				s := strconv.FormatInt(v[i], 10)
				strs = append(strs, s)
			}
			c.writeValue(valueName, strs, 12)
		}
	}
}

func (c *ListTag) writeValue(valueName string, value interface{}, Type int) {
	println(Type, valueName)
	newMap := make(map[string]interface{})
	c.writeSlice = append(c.writeSlice, newMap)
	newMap["tagType"] = Type
	newMap["name"] = valueName
	newMap["value"] = value
}
