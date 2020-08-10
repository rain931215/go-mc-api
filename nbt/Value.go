package nbt

import (
	"fmt"
	"strconv"
)

const (
	TagEnd = iota
	TagByte
	TagShort     //int16
	TagInt       //int32
	TagLong      //int64
	TagFloat     //float32
	TagDouble    //float64
	TagByteArray //[]byte
	TagString    //[]string
	TagList
	TagCompound
	TagIntArray  //[]int32
	TagLongArray //[]int64
)

type Value struct {
	Name  string
	Value interface{}
}

func NewValue(name string, value interface{}) *Value {
	v := new(Value)
	v.Name = name
	v.Value = value
	return v
}

func (v *Value) getType() (int, interface{}) {
	switch k := v.Value.(type) {
	case byte:
		return 1, v.Value
	case int16:
		return 2, v.Value
	case int32:
		return 3, v.Value
	case int64:
		s := strconv.FormatInt(k, 10)
		return 4, s
	case float32:
		return 5, v.Value
	case float64:
		return 6, v.Value
	case []byte:
		num := []int{}
		for i := 0; i < len(k); i++ {
			num = append(num, int(k[i]))
		}
		return 7, num
	case string:
		return 8, v.Value
	case *ListTag:
		return 9, k.toJSON()
	case *CompoundTag:
		return 10, k.toJSON()
	case []int:
		return 11, v.Value
	case []int64:
		strs := []string{}
		for i := 0; i < len(k); i++ {
			s := strconv.FormatInt(k[i], 10)
			strs = append(strs, s)
		}
		return 12, strs
	default:
		fmt.Println("Unknown Type")
	}
	return 0, nil
}
