package app

type Game interface {
	getMission() string
	getQuestion() any
	getAnswer() any
	prepareQuestionAndAnswer()
}

type AbstractGame struct {
	mission  string
	question any
	answer   any
}

func (g *AbstractGame) getMission() string {
	return g.mission
}

func (g *AbstractGame) getQuestion() any {
	return g.question
}

func (g *AbstractGame) getAnswer() any {
	return g.answer
}
