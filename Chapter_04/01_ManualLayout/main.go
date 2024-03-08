package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()

	square := canvas.NewRectangle(color.Black)
	circle := canvas.NewCircle(color.Transparent)
	circle.StrokeColor = &color.Gray{128}
	circle.StrokeWidth = 5
	box := container.NewWithoutLayout(square, circle)
	square.Move(fyne.NewPos(10, 10))
	square.Resize(fyne.NewSize(90, 90))
	circle.Move(fyne.NewPos(70, 70))
	circle.Resize(fyne.NewSize(120, 120))
	box.Resize(fyne.NewSize(200, 200))
	box.Refresh()

	w := a.NewWindow("Manual Layout")
	w.SetContent(box)
	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}
