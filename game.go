package bowlingkata

// Game keeps the score of a bowling game.
type Game struct {
	currentFrame  Frame
	previousFrame Frame
	score         int
}

// Score returns the total score of a game.
func (game *Game) Score() int {
	return game.score
}

// Roll registers a new roll to the game.
func (game *Game) Roll(pinsDown int) {
	game.currentFrame.Roll(pinsDown)

	if game.currentFrame.IsComplete() {
		game.score += game.previousFrame.Bonus(&game.currentFrame)
		game.score += game.currentFrame.Score()
		game.previousFrame = game.currentFrame
		game.currentFrame = *new(Frame)
	}
}
