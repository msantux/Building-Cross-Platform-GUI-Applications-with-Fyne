package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type diagonal struct{}

func (d *diagonal) MinSize(items []fyne.CanvasObject) fyne.Size {
	total := fyne.NewSize(0, 0)

	for _, obj := range items {
		if !obj.Visible() {
			continue
		}

		total = total.Add(obj.MinSize())
	}

	return total
}

func (d *diagonal) Layout(items []fyne.CanvasObject, size fyne.Size) {
	topLeft := fyne.NewPos(0, 0)
	for _, obj := range items {
		if !obj.Visible() {
			continue
		}

		size := obj.MinSize()

		obj.Move(topLeft)
		obj.Resize(size)
		topLeft = topLeft.Add(fyne.NewPos(size.Width, size.Height))
	}
}

func main() {
	a := app.New()

	w := a.NewWindow("Custom Layout")

	item1 := canvas.NewRectangle(color.Black)
	item1.SetMinSize(fyne.NewSize(35, 35))
	item2 := canvas.NewRectangle(color.Gray{128})
	item2.SetMinSize(fyne.NewSize(35, 35))
	item3 := canvas.NewRectangle(color.Black)
	item3.SetMinSize(fyne.NewSize(35, 35))
	myContainer := container.New(&diagonal{}, item1, item2, item3)

	w.SetContent(myContainer)
	w.ShowAndRun()
}
