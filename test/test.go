package main

import (
	"io/ioutil"

	"github.com/midnightfreddie/nbt2json"
	"github.com/rain931215/go-mc-api/plugin/autobuilder/nbt"
)

func main() {

	n := nbt.NewCompoundTag("test")
	/*
		n.AddValue("ByteTest", byte(0x01))
		n.AddValue("ShortTest", int16(123))
		n.AddValue("IntTest", int32(1234567))
		n.AddValue("LongTest", int64(91151548))
		n.AddValue("FloatTest", float32(3.14))
		n.AddValue("DoubleTest", float64(3.14159))
		n.AddValue("ByteArrayTest", []byte{1, 127, 3})
		n.AddValue("StringTest", "Hello")
		n.AddValue("CompoundTagTest", nbt.NewCompoundTag("CompoundTagTest"))
		n.AddValue("IntArrayTest", []int{1, 2, 3})
		n.AddValue("LongArrayTest", []int64{4, 5, 6})
	*/
	list := nbt.NewListTag("testList")
	v := nbt.NewValue("String", "Hello")
	list.AddValue("test", v)
	list.AddValue("test", v)
	list.AddValue("test", v)
	list.AddValue("test", v)
	n.AddValue("ListTest", list)

	n.WriteToJson()

	data, err := ioutil.ReadFile("viperTest.json")
	checkerr(err)

	out, err := nbt2json.Json2Nbt(data)
	checkerr(err)

	err = ioutil.WriteFile("test.nbt", out, 0644)
	checkerr(err)

	data, err = ioutil.ReadFile("test.nbt")
	checkerr(err)

	out, err = nbt2json.Nbt2Json(data, "test")
	checkerr(err)

	err = ioutil.WriteFile("test.json", out, 0644)
	checkerr(err)

}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
