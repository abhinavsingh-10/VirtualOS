package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne"

	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
)

func showGalleryApp() {
	// a := app.New()
	// w := a.NewWindow("image")
	//w.Resize(fyne.NewSize(800, 600))
	//var image :=(.image)
	root_src := "C:\\Users\\Abhinav Singh\\Pictures\\Screenshots"
	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}
	var picsArr []string
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() == false {
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "png" || extension == "jpeg" {
				picsArr = append(picsArr, root_src+"\\"+file.Name())

			}

		}
	}

	tabs := container.NewAppTabs(
		container.NewTabItem("image", canvas.NewImageFromFile(picsArr[0])),
		//container.NewTabItem("2", canvas.NewImageFromFile(picsArr[1])),
	)
	for i := 1; i < len(picsArr); i++ {
		tabs.Append(container.NewTabItem("image", canvas.NewImageFromFile(picsArr[i])))
	}
	tabs.SetTabLocation(container.TabLocationLeading)
	myWindow.SetContent(tabs)

	myWindow := myApp.NewWindow("image")

	myWindow.Resize(fyne.NewSize(500, 280))

	myWindow.SetContent(
		container.NewBorder(DeskBtn, nil, nil, nil, tabs),
	)

	myWindow.Show()

}
