package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func showWeatherApp(w fyne.Window) {

	//a := app.New()
	//w := a.NewWindow("weather app")

	//w.Resize(fyne.NewSize(560, 560))

	//api
	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=mumbai,delhi&APPID=c185cd354cac868eb959efa2bf024f53")
	if err != nil {
		fmt.Print(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Print(err)
	}

	weather, err := UnmarshalWeather(body)
	if err != nil {
		fmt.Print(err)
	}

	img := canvas.NewImageFromFile("weather.png")
	img.FillMode = canvas.ImageFillOriginal

	//check

	combo := widget.NewSelect([]string{"Mumbai", "Delhi", "chennai"}, func(value string) {
		log.Println("Select City", value)
	})

	label1 := canvas.NewText("weather Details", color.White)
	//label1.TextStyle = fyne.TextStyle{Bold: true}

	label2 := canvas.NewText(fmt.Sprintf("Country %s", weather.Sys.Country), color.Black)
	label3 := canvas.NewText(fmt.Sprintf("Wind Speed %.2f", weather.Wind.Speed), color.Black)
	label4 := canvas.NewText(fmt.Sprintf("Temp %2f", weather.Main.Temp), color.Black)

	weatherContainer := container.NewVBox(

		label1,
		img,

		combo,
		label2,
		label3,
		label4,
		container.NewGridWithColumns(1),
	)

	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, weatherContainer))
	w.Show()

}

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`
	Weather    []WeatherElement `json:"weather"`
	Base       string           `json:"base"`
	Main       Main             `json:"main"`
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`
	Clouds     Clouds           `json:"clouds"`
	Dt         int64            `json:"dt"`
	Sys        Sys              `json:"sys"`
	Timezone   int64            `json:"timezone"`
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	Cod        int64            `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherElement struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}
