package app

import (
	"errors"
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
	answer := noAnswer
	if isEven(n) {
		answer = yesAnswer
	}

	g.question = strconv.Itoa(n)
	g.answer = answer
}

func (g *EvenGame) askUserAnswer() string {
	_, userAnswer, _ := prompter.Select([]string{yesAnswer, noAnswer}).Run()

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

type CalcGame struct {
	AbstractGame
}

func (g *CalcGame) getGameName() string {
	return "calc"
}

func (g *CalcGame) getMission() string {
	return "What is the result of the expression?"
}

func (g *CalcGame) prepareQuestionAndAnswer() {
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

func (g *CalcGame) askUserAnswer() string {
	userAnswer, _ := prompter.Prompt(numberValidator).Run()

	return userAnswer
}

type GcdGame struct {
	AbstractGame
}

func (g *GcdGame) getGameName() string {
	return "gcd"
}

func (g *GcdGame) getMission() string {
	return "Find the greatest common divisor of given numbers."
}

func (g *GcdGame) prepareQuestionAndAnswer() {
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

func (g *GcdGame) askUserAnswer() string {
	userAnswer, _ := prompter.Prompt(numberValidator).Run()

	return userAnswer
}
