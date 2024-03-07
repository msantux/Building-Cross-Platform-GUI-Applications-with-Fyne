package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	w.SetContent(canvas.NewText("This works!", color.Black))
	w.ShowAndRun()
}
