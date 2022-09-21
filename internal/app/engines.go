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
	invalidGameErr     = errors.New("invalid game")
)

type Engine interface {
	welcome() error
	choose() (Game, error)
	play(game Game)
}

const roundsCount = 3

type BrainGamesEngine struct {
	username string
}

func (e *BrainGamesEngine) welcome() error {
	printer.PrintH1("Welcome to the Brain Games!")

	printer.PrintInfo("May I have your name?..")

	prompt := prompter.Prompt(func(input string) error {
		if len(input) < 2 {
			return invalidUsernameErr
		}

		return nil
	})

	username, _ := prompt.Run()
	if err := prompt.Validate(username); err != nil {
		return err
	}

	printer.PrintInfo(fmt.Sprintf("Hello, %s!", username))
	fmt.Println()

	e.username = username

	return nil
}

func (e *BrainGamesEngine) choose() (Game, error) {
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

	i, s, _ := prompt.Run()
	if s == "" {
		return nil, invalidGameErr
	}

	return Games[i], nil
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
