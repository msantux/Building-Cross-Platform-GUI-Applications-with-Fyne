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

	visibleCount := 0
	for _, obj := range items {
		if !obj.Visible() {
			continue
		}

		visibleCount++
	}

	min := d.MinSize(items)
	extraX := (size.Width - min.Width) / float32(visibleCount)
	extraY := (size.Height - min.Height) / float32(visibleCount)

	for _, obj := range items {
		if !obj.Visible() {
			continue
		}

		size := obj.MinSize()
		size = size.Add(fyne.NewSize(extraX, extraY))

		obj.Move(topLeft)
		obj.Resize(size)
		topLeft = topLeft.Add(fyne.NewPos(size.Width, size.Height))
	}
}

func main() {
	a := app.New()

	w := a.NewWindow("Custom Layout")

	item1 := canvas.NewRectangle(color.Black)
	item2 := canvas.NewRectangle(color.Gray{128})
	item3 := canvas.NewRectangle(color.Black)
	myContainer := container.New(&diagonal{}, item1, item2, item3)

	w.SetContent(myContainer)
	w.Resize(fyne.NewSize(120, 120))
	w.ShowAndRun()
}
