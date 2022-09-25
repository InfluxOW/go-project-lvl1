package app

import (
	"errors"
	"fmt"
	"github.com/InfluxOW/go-project-lvl1/internal/utils/fmt/prompter"
	"math"
	"math/rand"
	"strconv"
	"strings"
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

type progressionGame struct {
	abstractGame
}

func (g *progressionGame) GetName() string {
	return "progression"
}

func (g *progressionGame) GetMission() string {
	return "What number is missing in the progression?"
}

func (g *progressionGame) prepareQuestionAndAnswer() {
	const progressionLength = 10

	step := int(rand.Int63n(100))
	progressionStart := int(rand.Int63n(100))
	progression := generateProgression(step, progressionStart, progressionLength)
	hiddenElementIndex := int(rand.Int63n(progressionLength))

	g.question = getProgressionWithHiddenElementString(progression, hiddenElementIndex)
	g.answer = strconv.Itoa(progression[hiddenElementIndex])
}

func getProgressionWithHiddenElementString(progression []int, hiddenElementIndex int) string {
	var prStr []string
	for i, n := range progression {
		var str string
		if i == hiddenElementIndex {
			str = "?"
		} else {
			str = strconv.Itoa(n)
		}

		prStr = append(prStr, str)
	}

	return strings.Join(prStr, ", ")
}

func generateProgression(step, start, length int) []int {
	var pr []int
	for i := 0; i <= length; i++ {
		pr = append(pr, start+step*i)
	}

	return pr
}

func (g *progressionGame) askUserAnswer() string {
	prompt := prompter.Prompt(numberValidator)

	return prompter.RunPrompt(prompt)
}

type primeGame struct {
	abstractGame
}

func (g *primeGame) GetName() string {
	return "prime"
}

func (g *primeGame) GetMission() string {
	return fmt.Sprintf("Answer '%s' if given number is even, otherwise answer '%s'.", yesAnswer, noAnswer)
}

func (g *primeGame) prepareQuestionAndAnswer() {
	n := int(rand.Int63n(100))
	a := noAnswer
	if isPrime(n) {
		a = yesAnswer
	}

	g.question = strconv.Itoa(n)
	g.answer = a
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for start := 2; start <= n; start++ {
		if n%start == 0 {
			return false
		}
	}

	return true
}

func (g *primeGame) askUserAnswer() string {
	prompt := prompter.SimpleSelect([]string{yesAnswer, noAnswer})
	_, userAnswer := prompter.RunSelect(prompt)

	return userAnswer
}

type rootGame struct {
	abstractGame
}

func (g *rootGame) GetName() string {
	return "root"
}

func (g *rootGame) GetMission() string {
	return "Find an integer whose square is closest to the given one."
}

func (g *rootGame) prepareQuestionAndAnswer() {
	n := int(rand.Int63n(1000))

	g.question = strconv.Itoa(n)
	g.answer = strconv.Itoa(int(math.Sqrt(float64(n))))
}

func (g *rootGame) askUserAnswer() string {
	prompt := prompter.Prompt(numberValidator)

	return prompter.RunPrompt(prompt)
}
