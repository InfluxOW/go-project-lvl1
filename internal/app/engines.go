package app

import (
	"errors"
	"fmt"
	"github.com/InfluxOW/go-project-lvl1/internal/utils/fmt/printer"
	"github.com/InfluxOW/go-project-lvl1/internal/utils/fmt/prompter"
	"github.com/manifoldco/promptui"
)

var (
	invalidUsernameErr = errors.New("invalid username")
)

type engine interface {
	welcome()
	choose() game
	play(game game)
}

const roundsCount = 3

type brainGamesEngine struct {
	username string
}

func (e *brainGamesEngine) welcome() {
	printer.PrintH1("Welcome to the Brain games!")

	printer.PrintInfo("May I have your name?..")

	prompt := prompter.Prompt(func(input string) error {
		if len(input) < 2 {
			return invalidUsernameErr
		}

		return nil
	})

	username := prompter.RunPrompt(prompt)

	printer.PrintInfo(fmt.Sprintf("Hello, %s!", username))
	fmt.Println()

	e.username = username
}

func (e *brainGamesEngine) choose() game {
	templates := &promptui.SelectTemplates{
		Label:    "Available games:",
		Active:   fmt.Sprintf("%s {{ .GetName | green }}", promptui.IconSelect),
		Inactive: "  {{ .GetName | red }}",
		Selected: fmt.Sprintf(`{{ "%s" | green }} {{ .GetName | faint }}`, promptui.IconGood),
		Details: `
--------- Game ----------
{{ "Mission:" | faint }}	{{ .GetMission }}`,
	}

	prompt := prompter.Select("Game", games, templates, true)

	i, _ := prompter.RunSelect(prompt)

	return games[i]
}

func (e *brainGamesEngine) play(game game) {
	printer.PrintH2(game.GetMission(), "Game Mission:")

	for rc := 1; rc <= roundsCount; rc++ {
		game.prepareQuestionAndAnswer()

		printer.PrintH3(game.getQuestion(), "Question:")

		userAnswer := game.askUserAnswer()
		correctAnswer := game.getAnswer()

		if userAnswer == correctAnswer {
			printer.PrintSuccess("Correct!")
			fmt.Println()

			continue
		}

		printer.PrintFailure(fmt.Sprintf("'%s' is wrong answer ;(. Correct answer was '%s'", userAnswer, correctAnswer))
		fmt.Println()
		printer.PrintlnFailure(fmt.Sprintf("Let's try again, %s!", e.username))

		return
	}

	printer.PrintlnSuccess(fmt.Sprintf("Congratulations, %s!", e.username))
}
