package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Data List")

	strings := binding.NewStringList()
	fmt.Println("String list length:", strings.Length())

	strings.Append("astring")
	fmt.Println("String list length:", strings.Length())

	val, _ := strings.GetValue(0)
	fmt.Println("String at 0:", val)

	src := []string{"one"}
	strings2 := binding.BindStringList(&src)
	fmt.Println("String list 2 length:", strings2.Length())

	strings2.Append("astring")
	fmt.Println("String list 2 length:", strings2.Length())

	val2, _ := strings2.GetValue(0)
	fmt.Println("String at 0:", val2)

	ls := widget.NewListWithData(
		strings2,
		func() fyne.CanvasObject {
			return widget.NewLabel("placeholder")
		},
		func(item binding.DataItem, obj fyne.CanvasObject) {
			text := obj.(*widget.Label)
			text.Bind(item.(binding.String))
		})

	w.SetContent(container.NewVBox(ls))
	w.ShowAndRun()
}
