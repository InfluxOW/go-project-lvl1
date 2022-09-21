package prompter

import (
	"github.com/manifoldco/promptui"
)

const (
	label = "Answer"
)

func Select(items interface{}) *promptui.Select {
	return &promptui.Select{
		Label:    label,
		Items:    items,
		HideHelp: true,
	}
}

func Prompt(validate func(input string) error) *promptui.Prompt {
	return &promptui.Prompt{
		Label:    label,
		Validate: validate,
	}
}
