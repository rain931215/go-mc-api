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

func (v *Value) getType() int {
	switch k := v.Value.(type) {
	case byte:
		return 1
	case int16:
		return 2
	case int32:
		return 3
	case int64:
		s := strconv.FormatInt(k, 10)
		v.Value = s
		return 4
	case float32:
		return 5
	case float64:
		return 6
	case []byte:
		num := []int{}
		for i := 0; i < len(k); i++ {
			num = append(num, int(k[i]))
		}
		v.Value = num
		return 7
	case string:
		return 8
	case *ListTag:
		v.Value = k.toJSON()
		return 9
	case *CompoundTag:
		v.Value = k.toJSON()
		return 10
	case []int:
		return 11
	case []int64:
		strs := []string{}
		for i := 0; i < len(k); i++ {
			s := strconv.FormatInt(k[i], 10)
			strs = append(strs, s)
		}
		v.Value = strs
		return 12
	default:
		fmt.Println("Unknown Type")
	}
	return 0
}
