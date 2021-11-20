package main

// import fyne
import (
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showcolorApp() {
	// New app
	a := app.New()
	// Title and New window
	w := a.NewWindow("Random Color Generator")
	// Resize window
	w.Resize(fyne.NewSize(400, 400))
	// New Rectange
	// Color for new Rectange
	colorx := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	rect1 := canvas.NewRectangle(colorx)
	rect1.SetMinSize(fyne.NewSize(200, 200))
	// Btn for color change
	btn1 := widget.NewButton("Random color", func() {
		// UI is done.. Now Logic
		// unit8 is necessary to convert int to uint8
		rect1.FillColor = color.NRGBA{R: uint8(rand.Intn(255)),
			G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 255}
		rect1.Refresh() // refresh screen
	})
	// Random RED Changer
	btnRed := widget.NewButton("Random RED", func() {
		// UI is done.. Now Logic
		// unit8 is necessary to convert int to uint8
		rect1.FillColor = color.NRGBA{R: uint8(rand.Intn(255)),
			G: 0, B: 0, A: 255}
		rect1.Refresh() // refresh screen
	})
	// Random GREEN COLOR
	btnGreen := widget.NewButton("Random Green", func() {
		// UI is done.. Now Logic
		// unit8 is necessary to convert int to uint8
		rect1.FillColor = color.NRGBA{R: 0,
			G: uint8(rand.Intn(255)), B: 0, A: 255}
		rect1.Refresh() // refresh screen
	})
	// Random Blue COLOR
	btnBlue := widget.NewButton("Random Blue", func() {
		// UI is done.. Now Logic
		// unit8 is necessary to convert int to uint8
		rect1.FillColor = color.NRGBA{R: 0,
			G: 0, B: uint8(rand.Intn(255)), A: 255}
		rect1.Refresh() // refresh screen
	})
	w.SetContent(
		container.NewVBox(
			rect1,
			btn1,
			btnRed,
			btnGreen,
			btnBlue,
		),
	)
	w.Show()
}
