package bowlingkata

// Game keeps the score of a bowling game.
type Game struct {
	currentFrame        Frame
	previousFrame       Frame
	beforePreviousFrame Frame
	score               int
}

// Score returns the total score of a game.
func (game *Game) Score() int {
	return game.score
}

// Roll registers a new roll to the game.
func (game *Game) Roll(pinsDown int) {
	game.currentFrame.Roll(pinsDown)

	if game.currentFrame.IsComplete() {
		game.score += game.beforePreviousFrame.Bonus(
			game.previousFrame,
			game.currentFrame,
		)
		game.score += game.currentFrame.Score()
		game.beforePreviousFrame = game.previousFrame
		game.previousFrame = game.currentFrame
		game.currentFrame = *new(Frame)
	}
}
