package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type greeter struct {
	greeting       *widget.Label
	name           *widget.Entry
	updateGreeting *widget.Button
}

func (g *greeter) makeUI() fyne.CanvasObject {
	g.greeting = widget.NewLabel("Hello World!")
	g.name = widget.NewEntry()
	g.name.PlaceHolder = "Enter name"
	g.updateGreeting = widget.NewButton("Welcom", g.setGreeting)

	return container.NewVBox(g.greeting, g.name, g.updateGreeting)
}

func (g *greeter) setGreeting() {
	text := fmt.Sprintf("Hello %s!", g.name.Text)
	g.greeting.SetText(text)
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello!")

	g := &greeter{}
	w.SetContent(g.makeUI())
	w.ShowAndRun()
}
