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

type Engine interface {
	welcome()
	choose() Game
	play(game Game)
}

const roundsCount = 3

type BrainGamesEngine struct {
	username string
}

func (e *BrainGamesEngine) welcome() {
	printer.PrintH1("Welcome to the Brain Games!")

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

func (e *BrainGamesEngine) choose() Game {
	templates := &promptui.SelectTemplates{
		Label:    "Available Games:",
		Active:   fmt.Sprintf("%s {{ .GetName | green }}", promptui.IconSelect),
		Inactive: "  {{ .GetName | red }}",
		Selected: fmt.Sprintf(`{{ "%s" | green }} {{ .GetName | faint }}`, promptui.IconGood),
		Details: `
--------- Game ----------
{{ "Mission:" | faint }}	{{ .GetMission }}`,
	}

	prompt := prompter.Select("Game", Games, templates, true)

	i, _ := prompter.RunSelect(prompt)

	return Games[i]
}

func (e *BrainGamesEngine) play(game Game) {
	printer.PrintH2(game.GetMission(), "Game mission:")

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
