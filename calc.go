package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"github.com/Knetic/govaluate"
)

func showCalc() {

	output := ""
	input := widget.NewLabel(output)
	ishistory := false
	historystr := " "
	history := widget.NewLabel(historystr)
	var historyarr []string
	historyBtn := widget.NewButton("history", func() {

		if ishistory {

		} else {
			for i := len(historyarr) - 1; i >= 0; i-- {
				historystr = historystr + historyarr[i]
				historystr += "\n"

			}

		}
		ishistory = !ishistory
		history.SetText(historystr)
	})

	backBtn := widget.NewButton("back", func() {

		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}

	})

	clearBtn := widget.NewButton("clear", func() {

		output = ""
		input.SetText(output)
	})

	openBtn := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)
	})

	closeBtn := widget.NewButton(")", func() {
		output = output + ")"
		input.SetText(output)
	})

	divideBtn := widget.NewButton("/", func() {
		output = output + "/"
		input.SetText(output)
	})

	sevenBtn := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)
	})

	eightBtn := widget.NewButton("8", func() {
		output = output + "8"
		input.SetText(output)
	})

	nineBtn := widget.NewButton("9", func() {
		output = output + "9"
		input.SetText(output)
	})

	multiplyeBtn := widget.NewButton("*", func() {
		output = output + "*"
		input.SetText(output)
	})

	fourBtn := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})

	fiveBtn := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)
	})

	sixBtn := widget.NewButton("6", func() {
		output = output + "6"
		input.SetText(output)
	})

	minusBtn := widget.NewButton("-", func() {
		output = output + "-"
		input.SetText(output)
	})

	oneBtn := widget.NewButton("1", func() {
		output = output + "1"
		input.SetText(output)
	})

	twoBtn := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)
	})

	threeBtn := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})

	pluseBtn := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})

	zeroBtn := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)
	})

	dotBtn := widget.NewButton(".", func() {
		output = output + "."
		input.SetText(output)
	})

	equalBtn := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strToAppend := output + "=" + ans
				historyarr = append(historyarr, strToAppend)
				output = ans

			} else {
				output = "error"

			}
		} else {
			output = "error"
		}

		input.SetText(output)
		fmt.Print(historyarr)

	})
	equalBtn.Importance = widget.HighImportance

	calcContainer := container.NewVBox(
		container.NewVBox(
			input,
			history,
			container.NewGridWithColumns(1,
				container.NewGridWithColumns(2,
					historyBtn,
					backBtn,
				),
				container.NewGridWithColumns(4,
					clearBtn,
					openBtn,
					closeBtn,
					divideBtn),

				container.NewGridWithColumns(4,
					nineBtn,
					eightBtn,
					sevenBtn,
					multiplyeBtn),

				container.NewGridWithColumns(4,
					fourBtn,
					fiveBtn,
					sixBtn,
					minusBtn),

				container.NewGridWithColumns(4,
					oneBtn,
					twoBtn,
					threeBtn,
					pluseBtn),

				container.NewGridWithColumns(2,
					container.NewGridWithColumns(2,
						zeroBtn,
						dotBtn),
					equalBtn,
				),
			),
		))

	w := myApp.NewWindow("Calc")

	w.Resize(fyne.NewSize(500, 280))

	w.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, calcContainer),
	)

	w.Show()

}
