package app

import (
	"fmt"
	"github.com/InfluxOW/go-project-lvl1/internal/utils/fmt/prompter"
	"math/rand"
	"strconv"
	"time"
)

const (
	yesAnswer = "yes"
	noAnswer  = "no"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Game interface {
	getGameName() string
	getMission() string
	getQuestion() string
	getAnswer() string
	prepareQuestionAndAnswer()
	askUserAnswer() string
}

type AbstractGame struct {
	question string
	answer   string
}

func (g *AbstractGame) getQuestion() string {
	return g.question
}

func (g *AbstractGame) getAnswer() string {
	return g.answer
}

type EvenGame struct {
	AbstractGame
}

func (g *EvenGame) getGameName() string {
	return "even"
}

func (g *EvenGame) getMission() string {
	return fmt.Sprintf("Answer '%s' if given number is even, otherwise answer '%s'.", yesAnswer, noAnswer)
}

func (g *EvenGame) prepareQuestionAndAnswer() {
	n := int(rand.Int63n(1000))
	a := noAnswer
	if isEven(n) {
		a = yesAnswer
	}

	g.question = strconv.Itoa(n)
	g.answer = a
}

func (g *EvenGame) askUserAnswer() string {
	_, userAnswer, _ := prompter.Select([]string{yesAnswer, noAnswer}).Run()

	return userAnswer
}

func isEven(n int) bool {
	return n%2 == 0
}
