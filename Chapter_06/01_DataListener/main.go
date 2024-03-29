package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/data/binding"
)

func main() {
	val := binding.NewString()

	callback := binding.NewDataListener(func() {
		str, _ := val.Get()
		fmt.Println("String changed to:", str)
	})
	val.AddListener(callback)

	time.Sleep(time.Millisecond * 100)

	val.Set("new data")
	time.Sleep(time.Millisecond * 100)
}
