package app

import (
	"github.com/InfluxOW/go-project-lvl1/internal/utils/fmt/printer"
	osUtils "github.com/InfluxOW/go-project-lvl1/internal/utils/os"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

var games = []game{&evenGame{}, &gcdGame{}, &calcGame{}, &progressionGame{}, &primeGame{}, &rootGame{}}

func Play(random bool) {
	gracefullyHandleShutdown()

	e := chooseEngine()
	e.welcome()

	g := chooseGame(random, e)
	e.play(g)

	os.Exit(osUtils.ExitCodeSuccess)
}

func gracefullyHandleShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGSTOP, syscall.SIGKILL, syscall.SIGTERM)
	go func() {
		<-c
		printer.PrintlnInfo("See you later!")
		os.Exit(osUtils.ExitCodeSuccess)
	}()
}

func chooseEngine() engine {
	e := brainGamesEngine{}

	return &e
}

func chooseGame(random bool, e engine) game {
	var g game
	if random {
		g = games[rand.Int63n(int64(len(games)))]
	} else {
		g = e.choose()
	}

	return g
}
