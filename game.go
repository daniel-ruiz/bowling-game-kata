package game

type game struct {
	score int
}

func New() game {
	game := game{}
	return game
}
