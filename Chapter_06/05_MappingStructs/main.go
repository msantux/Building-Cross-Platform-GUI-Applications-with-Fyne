package main

import (
	"fmt"

	"fyne.io/fyne/v2/data/binding"
)

type person struct {
	Name string
	Age  int
}

func main() {
	src := person{Name: "John Doe", Age: 21}
	values := binding.BindStruct(&src)
	fmt.Println("Map size:", len(values.Keys()))
	name, _ := values.GetValue("Name")
	fmt.Println("Value for Name:", name)
	age, _ := values.GetValue("Age")
	fmt.Println("Value for Age:", age)
}
