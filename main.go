package main

import (
	"chatgpt-client/chatgptAPI"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"os"
)

func main() {
	gptClient, err := chatgptAPI.Init()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	a := app.New()
	w := a.NewWindow("Hello")

	w.Resize(fyne.NewSize(700, 900))

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	//input.OnSubmitted("t")

	data := binding.NewUntypedList()
	info := widget.NewLabel("Please Wait...")
	infinite := widget.NewProgressBarInfinite()
	modalBox := container.NewVBox(container.NewCenter(info), infinite)
	d := dialog.NewCustom("Sending Request", "Cancel", modalBox, w)

	submitButton := widget.NewButtonWithIcon("Submit", theme.MailSendIcon(), func() {
		infinite.Start()
		d.Show()

		answer, err := gptClient.AskGPT(input.Text)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		err = data.Append(QuestionAndAnswer{input.Text, *answer})
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		input.SetText("")
		d.Hide()
		infinite.Stop()
	})

	submitButton.Disable()

	input.OnChanged = func(s string) {
		submitButton.Disable()

		if len(s) >= 3 {
			submitButton.Enable()
		}
	}

	w.SetContent(
		container.NewBorder(
			nil,
			container.NewBorder(
				nil,
				nil,
				nil,
				submitButton,
				input,
			),
			nil,
			nil,
			widget.NewListWithData(
				data,
				func() fyne.CanvasObject {
					return container.NewBorder(
						nil,
						nil,
						nil,
						widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
							notification := dialog.NewConfirm("Not Supported", "Do you still love us? This feature hasn't been implemented yet.", nil, w)
							notification.Show()
						}),
						container.NewBorder(
							widget.NewLabel(""),
							nil,
							nil,
							nil,
							widget.NewLabel(""),
						),
					)
				},
				func(di binding.DataItem, o fyne.CanvasObject) {
					ctr, _ := o.(*fyne.Container)
					innerCtr := ctr.Objects[0].(*fyne.Container)
					question := innerCtr.Objects[1].(*widget.Label)
					question.Wrapping = fyne.TextWrapWord
					answer := innerCtr.Objects[0].(*widget.Label)
					answer.Wrapping = fyne.TextWrapWord
					questionAndAnswer := NewQuestionAndAnswerFromDataItem(di)

					question.SetText("Q: " + questionAndAnswer.Question)
					answer.SetText("A: " + questionAndAnswer.Answer)
				}),
		),
	)

	w.ShowAndRun()
}
