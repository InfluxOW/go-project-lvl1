package app

import (
	"github.com/InfluxOW/go-project-lvl1/internal/utils/fmt/printer"
	osUtils "github.com/InfluxOW/go-project-lvl1/internal/utils/os"
	"math/rand"
	"os"
)

func Play(random bool) {
	engine := BrainGamesEngine{}
	err := engine.welcome()
	handleErr(err)

	var game Game
	if random {
		game = Games[rand.Int63n(int64(len(Games)))]
	} else {
		g, err := engine.choose()
		handleErr(err)

		game = g
	}
	engine.play(game)
}

func handleErr(err error) {
	if err != nil {
		printer.PrintlnFailure(err.Error())
		os.Exit(osUtils.ExitCodeError)
	}
}
