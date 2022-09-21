package app

import (
	"errors"
	"fmt"
	"github.com/InfluxOW/go-project-lvl1/internal/utils/fmt/prompter"
	"math/rand"
	"strconv"
	"time"
)

var games = []game{&evenGame{}, &gcdGame{}, &calcGame{}}

const (
	yesAnswer = "yes"
	noAnswer  = "no"
)

var (
	numberValidator = func(input string) error {
		if _, err := strconv.Atoi(input); err != nil {
			return invalidNumberErr
		}

		return nil
	}
)

var (
	invalidNumberErr = errors.New("invalid number")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type game interface {
	GetName() string
	GetMission() string
	getQuestion() string
	getAnswer() string
	prepareQuestionAndAnswer()
	askUserAnswer() string
}

type abstractGame struct {
	question string
	answer   string
}

func (g *abstractGame) getQuestion() string {
	return g.question
}

func (g *abstractGame) getAnswer() string {
	return g.answer
}

type evenGame struct {
	abstractGame
}

func (g *evenGame) GetName() string {
	return "even"
}

func (g *evenGame) GetMission() string {
	return fmt.Sprintf("Answer '%s' if given number is even, otherwise answer '%s'.", yesAnswer, noAnswer)
}

func (g *evenGame) prepareQuestionAndAnswer() {
	n := int(rand.Int63n(1000))
	answer := noAnswer
	if isEven(n) {
		answer = yesAnswer
	}

	g.question = strconv.Itoa(n)
	g.answer = answer
}

func (g *evenGame) askUserAnswer() string {
	prompt := prompter.SimpleSelect([]string{yesAnswer, noAnswer})
	_, userAnswer := prompter.RunSelect(prompt)

	return userAnswer
}

func isEven(n int) bool {
	return n%2 == 0
}

type operation string

const (
	addition       operation = "+"
	subtraction    operation = "-"
	multiplication operation = "*"
)

var operations = []operation{addition, subtraction, multiplication}

type calcGame struct {
	abstractGame
}

func (g *calcGame) GetName() string {
	return "calc"
}

func (g *calcGame) GetMission() string {
	return "What is the result of the expression?"
}

func (g *calcGame) prepareQuestionAndAnswer() {
	a := int(rand.Int63n(50))
	b := int(rand.Int63n(50))
	op := operations[rand.Int63n(int64(len(operations)))]

	var answer int
	switch op {
	case addition:
		answer = a + b
	case subtraction:
		answer = a - b
	case multiplication:
		answer = a * b
	}

	g.question = fmt.Sprintf("%d %s %d", a, op, b)
	g.answer = strconv.Itoa(answer)
}

func (g *calcGame) askUserAnswer() string {
	prompt := prompter.Prompt(numberValidator)

	return prompter.RunPrompt(prompt)
}

type gcdGame struct {
	abstractGame
}

func (g *gcdGame) GetName() string {
	return "gcd"
}

func (g *gcdGame) GetMission() string {
	return "Find the greatest common divisor of given numbers."
}

func (g *gcdGame) prepareQuestionAndAnswer() {
	a := int(rand.Int63n(100))
	b := int(rand.Int63n(100))

	g.question = fmt.Sprintf("%d, %d", a, b)
	g.answer = strconv.Itoa(gcd(a, b))
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func (g *gcdGame) askUserAnswer() string {
	prompt := prompter.Prompt(numberValidator)

	return prompter.RunPrompt(prompt)
}
