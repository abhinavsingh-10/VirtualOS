package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
)

var count int = 1

func showTextEditor() {
	a := app.New()

	w := a.NewWindow("MY EDITER")
	w.Resize(fyne.NewSize(600, 600))

	//1st thing button k click hone pe new file open hona chaiye
	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("text Editer"),
		),
	)
	// add new file we will  create a button

	content.Add(widget.NewButton("Add New File", func() {
		content.Add(widget.NewLabel("New File" + strconv.Itoa(count)))
		count++

	}))

	// here we will whatevere we want
	//we will enter the text for multilines
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")

	input.Resize(fyne.NewSize(400, 400))

	//we will be creating the save button
	//we will save or open out file in local computer
	saveBtn := widget.NewButton("save text File", func() {
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)

				uc.Write(textData)

			}, w)

		saveFileDialog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")

		saveFileDialog.Show()

	})

	//file open

	openBtn := widget.NewButton("open text file", func() {

		openFileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				ReadData, _ := ioutil.ReadAll(r)

				ouput := fyne.NewStaticResource("New File", ReadData)

				viewData := widget.NewMultiLineEntry()

				viewData.SetText(string(ouput.StaticContent))

				w := fyne.CurrentApp().NewWindow(
					string(ouput.StaticName))

				w.SetContent(container.NewScroll(viewData))

				w.Resize(fyne.NewSize(400, 400))

				w.Show()

			}, w)

		openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))

		openFileDialog.Show()
	})

	w.SetContent(
		container.NewVBox(
			content,
			input,

			container.NewHBox(
				saveBtn,
				openBtn,
			),
		),
	)

	w.Show()

}
