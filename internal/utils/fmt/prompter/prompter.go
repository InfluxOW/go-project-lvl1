package prompter

import (
	"github.com/manifoldco/promptui"
)

const (
	defaultLabel = "Answer"
)

func SimpleSelect(items interface{}) *promptui.Select {
	return Select(defaultLabel, items, nil, false)
}

func Select(label string, items interface{}, templates *promptui.SelectTemplates, hideSelected bool) *promptui.Select {
	return &promptui.Select{
		Label:        label,
		Items:        items,
		Templates:    templates,
		HideHelp:     true,
		HideSelected: hideSelected,
	}
}

func Prompt(validate func(input string) error) *promptui.Prompt {
	return &promptui.Prompt{
		Label:    defaultLabel,
		Validate: validate,
	}
}
