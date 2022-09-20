package app

type Engine interface {
	start(game Game)

	welcome()
}
