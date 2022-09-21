package app

import (
	"errors"
	"fmt"
	"github.com/InfluxOW/go-project-lvl1/internal/utils/fmt/printer"
	"github.com/InfluxOW/go-project-lvl1/internal/utils/fmt/prompter"
)

var (
	invalidUsernameErr = errors.New("invalid username")
)

type Engine interface {
	play(game Game)

	welcome()
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

	username, _ := prompt.Run()
	printer.PrintInfo(fmt.Sprintf("Hello, %s!", username))
	fmt.Println()

	e.username = username
}

func (e *BrainGamesEngine) play(game Game) {
	printer.PrintH2(game.getMission(), "Game mission:")

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
