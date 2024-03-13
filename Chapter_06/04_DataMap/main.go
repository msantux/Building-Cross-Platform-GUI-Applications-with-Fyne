package main

import (
	"fmt"

	"fyne.io/fyne/v2/data/binding"
)

func main() {
	values := binding.NewUntypedMap()
	fmt.Println("Map size:", len(values.Keys()))
	_ = values.SetValue("akey", 5)
	fmt.Println("Map size:", len(values.Keys()))
	val, _ := values.GetValue("akey")
	fmt.Println("Value at akey:", val)

	src := map[string]interface{}{"akey": "before"}
	values2 := binding.BindUntypedMap(&src)
	fmt.Println("Map size:", len(values2.Keys()))
	_ = values2.SetValue("newkey", 5)
	fmt.Println("Map size:", len(values2.Keys()))
	val, _ = values2.GetValue("akey")
	fmt.Println("Value at akey:", val)

}
