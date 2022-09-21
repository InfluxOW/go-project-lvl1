package app

import (
	"github.com/InfluxOW/go-project-lvl1/internal/utils/fmt/printer"
	osUtils "github.com/InfluxOW/go-project-lvl1/internal/utils/os"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func Play(random bool) {
	gracefullyHandleShutdown()

	engine := BrainGamesEngine{}
	engine.welcome()

	var game Game
	if random {
		game = Games[rand.Int63n(int64(len(Games)))]
	} else {
		game = engine.choose()
	}

	engine.play(game)

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
