package bowlingkata

// Game keeps the score of a bowling game.
type Game struct {
	score int
}

// Score returns the total score of a game.
func (game *Game) Score() int {
	return game.score
}

// Roll registers a new roll to the game.
func (game *Game) Roll(pinsDown int) {
	game.score += pinsDown
}
