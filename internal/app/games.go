package app

import (
	"math/rand"
	"strconv"
	"time"
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
	return "Answer 'yes' if given number is even, otherwise answer 'no'."
}

func (g *EvenGame) prepareQuestionAndAnswer() {
	n := int(rand.Int63n(1000))
	a := "no"
	if isEven(n) {
		a = "yes"
	}

	g.question = strconv.Itoa(n)
	g.answer = a
}

func isEven(n int) bool {
	return n%2 == 0
}
