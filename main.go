package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

var myApp fyne.App = app.New()

var myWindow fyne.Window = myApp.NewWindow(" my oS ")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var btn6 fyne.Widget

var img fyne.CanvasObject
var DeskBtn fyne.Widget

var panelContent *fyne.Container

func main() {

	myApp.Settings().SetTheme(theme.LightTheme())
	myWindow := myApp.NewWindow("os")

	img = canvas.NewImageFromFile("C:\\Users\\Abhinav Singh\\Desktop\\Fyne\\image.png")

	btn1 = widget.NewButtonWithIcon("Weather App", theme.InfoIcon(), func() {
		showWeatherApp(myWindow)
	})

	btn2 = widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(), func() {
		showCalc()
	})

	btn3 = widget.NewButtonWithIcon("Gallery App", theme.StorageIcon(), func() {
		showGalleryApp()
	})

	btn4 = widget.NewButtonWithIcon("Text App", theme.FileTextIcon(), func() {
		showTextEditor()
	})

	btn5 = widget.NewButtonWithIcon("Music App", theme.MediaPlayIcon(), func() {
		showmusicApp()
	})

	btn6 = widget.NewButtonWithIcon("Color ", theme.ColorPaletteIcon(), func() {
		showcolorApp()
	})

	DeskBtn = widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))

	})

	panelContent = container.NewVBox((container.NewGridWithColumns(7, DeskBtn, btn1, btn2, btn3, btn4, btn5, btn6)))

	myWindow.Resize(fyne.NewSize(1280, 720))
	myWindow.CenterOnScreen()

	myWindow.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, img),
	)

	myWindow.ShowAndRun()

}
