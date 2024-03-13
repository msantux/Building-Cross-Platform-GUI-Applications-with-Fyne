package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
)

func main() {
	a := app.NewWithID("com.example.preferences")

	key := "demokey"
	a.Preferences().SetString(key, "somevalue")
	val := a.Preferences().String(key)
	fmt.Println("Value is:", val)

	key2 := "anotherkey"
	val2 := a.Preferences().String(key2)
	fmt.Println("Value is:", val2)
	val2 = a.Preferences().StringWithFallback(key2, "missing")
	fmt.Println("Value is:", val2)

	fmt.Println("Removing")
	a.Preferences().RemoveValue(key)
	a.Preferences().RemoveValue(key2)
	val = a.Preferences().String(key)
	fmt.Println("Value is:", val)

	data := binding.BindPreferenceString("demokey", a.Preferences())
	data.Set("somevalue")
	val, _ = data.Get()
	fmt.Println("Bound value:", val)
	a.Preferences().RemoveValue("demokey")
	val, _ = data.Get()
	fmt.Println("Bound value:", val)
}
