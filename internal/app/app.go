package app

func Play(random bool) {
	engine := BrainGamesEngine{}
	engine.welcome()

	game := EvenGame{}
	engine.play(&game)
}
