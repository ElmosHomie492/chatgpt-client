package main

import "fyne.io/fyne/v2/data/binding"

type QuestionAndAnswer struct {
	Question string
	Answer   string
}

func NewQuestionAndAnswerFromDataItem(item binding.DataItem) QuestionAndAnswer {
	v, _ := item.(binding.Untyped).Get()
	return v.(QuestionAndAnswer)
}
