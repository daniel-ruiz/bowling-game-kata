package bowlingkata

type game struct {
	score int
}

// New initializes the game struct.
func New() *game {
	game := &game{}
	return game
}

func (game *game) Score() int {
	return game.score
}

func (game *game) Roll(pinsDown int) {
	game.score += pinsDown
}
