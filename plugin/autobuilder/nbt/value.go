package nbt

type Value struct {
	name  string
	value interface{}
}

func NewValue(name string, value interface{}) *Value {
	v := new(Value)
	v.name = name
	v.value = value
	return v
}
