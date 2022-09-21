package prompter

import (
	"github.com/manifoldco/promptui"
	"syscall"
	"time"
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

func RunSelect(p *promptui.Select) (int, string) {
	i, s, err := p.Run()

	handleInterrupt(err)

	return i, s
}

func Prompt(validate func(input string) error) *promptui.Prompt {
	return &promptui.Prompt{
		Label:    defaultLabel,
		Validate: validate,
	}
}

func RunPrompt(p *promptui.Prompt) string {
	s, err := p.Run()

	handleInterrupt(err)

	return s
}

func handleInterrupt(err error) {
	if err == promptui.ErrInterrupt || err == promptui.ErrEOF {
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)

		time.Sleep(time.Second)
	}
}
