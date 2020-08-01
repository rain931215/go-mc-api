package nbt

import (
	"strconv"

	"github.com/spf13/viper"
)

type CompoundTag struct {
	name     string
	values   []*Value
	writeMap []map[string]interface{}
}

func NewCompoundTag(name string) *CompoundTag {
	CompoundTag := new(CompoundTag)
	CompoundTag.name = name
	return CompoundTag
}

func (c *CompoundTag) AddValue(key string, newValue interface{}) {
	v := &Value{name: key, value: newValue}
	c.values = append(c.values, v)
}

func (c *CompoundTag) WriteToJson() {
	nbt := make([]map[string]interface{}, 0)
	c.writeMap = make([]map[string]interface{}, 0)
	file := viper.New()
	file.SetConfigName("viperTest")
	file.SetConfigType("json")

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
			num := []int{}
			for i := 0; i < len(v); i++ {
				num = append(num, int(v[i]))
			}
			c.writeValue(valueName, num, 7)
		case string:
			c.writeValue(valueName, v, 8)
		case *ListTag:
			list := c.values[i].value.(*ListTag)
			list.writeListTag()
			newMap := make(map[string]interface{})
			newMap["tagType"] = 9
			newMap["name"] = list.name
			newMap["value"] = list.writeSlice
			c.writeMap = append(c.writeMap, newMap)
		case *CompoundTag:
			tag := c.values[i].value.(*CompoundTag)
			tag.writeCompoundTag()
			newMap := make(map[string]interface{})
			newMap["tagType"] = 10
			newMap["name"] = tag.name
			newMap["value"] = tag.writeMap
			c.writeMap = append(c.writeMap, newMap)
		case []int:
			c.writeValue(valueName, v, 11)
		case []int64:
			strs := []string{}
			for i := 0; i < len(v); i++ {
				s := strconv.FormatInt(v[i], 10)
				strs = append(strs, s)
			}
			c.writeValue(valueName, strs, 12)
		default:
			println(v)
			return
		}
	}
	newMap := make(map[string]interface{})
	newMap["tagType"] = 10
	newMap["name"] = c.name
	newMap["value"] = c.writeMap
	nbt = append(nbt, newMap)
	file.Set("nbt", nbt)

	file.AddConfigPath(".")
	file.WriteConfig()
}

func (c *CompoundTag) writeCompoundTag() {
	c.writeMap = make([]map[string]interface{}, 0)
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
			c.writeMap = append(c.writeMap, newMap)
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

func (c *CompoundTag) writeValue(valueName string, value interface{}, Type int) {
	println(Type, valueName)
	newMap := make(map[string]interface{})
	c.writeMap = append(c.writeMap, newMap)
	newMap["tagType"] = Type
	newMap["name"] = valueName
	newMap["value"] = value
}
