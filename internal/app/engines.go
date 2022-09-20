package app

import "fmt"

type Engine interface {
	play(game Game)

	welcome()
}

const roundsCount = 3

type BrainGamesEngine struct {
	username string
}

func (e *BrainGamesEngine) welcome() {
	fmt.Println("Welcome to the Brain Games!")
	fmt.Println("May I have your name?")
	var username string
	fmt.Scan(&username)
	fmt.Println(fmt.Sprintf("Hello, %s!", username))
	fmt.Println()

	e.username = username
}

func (e *BrainGamesEngine) play(game Game) {
	fmt.Println(fmt.Sprintf("(!!!) Game mission: %s", game.getMission()))
	fmt.Println()

	for rc := 1; rc <= roundsCount; rc++ {
		game.prepareQuestionAndAnswer()

		fmt.Println(fmt.Sprintf("Question: %s", game.getQuestion()))
		fmt.Println("Your answer...")
		var userAnswer string
		fmt.Scan(&userAnswer)
		correctAnswer := game.getAnswer()

		if userAnswer == correctAnswer {
			fmt.Println("Correct!")
		} else {
			fmt.Println(fmt.Sprintf("'%s' is wrong answer ;(. Correct answer was '%s'", userAnswer, correctAnswer))
			fmt.Println(fmt.Sprintf("Let's try again, %s!", e.username))

			return
		}

		fmt.Println()
	}

	fmt.Println(fmt.Sprintf("Congratulations, %s!", e.username))
}
