package bowlingkata

const maxFrameScore int = 10

// Frame represents every frame in a bowling game.
type Frame struct {
	rolls     [2]int
	rollCount int
}

// Roll annotates the pins taken down in a frame roll.
func (frame *Frame) Roll(pinsDown int) {
	if frame.IsComplete() {
		return
	}

	frame.rolls[frame.rollCount] = pinsDown
	frame.rollCount++
}

// IsComplete determines if a frame has finished
func (frame *Frame) IsComplete() bool {
	return frame.IsStrike() || (frame.rollCount == 2)
}

// Score calculates total score of a frame.
func (frame *Frame) Score() int {
	score := 0

	for _, pinsDown := range frame.rolls {
		score += pinsDown
	}

	return score
}

// IsSpare determines if a frame is a spare.
func (frame *Frame) IsSpare() bool {
	return frame.rolls[0] < maxFrameScore && frame.Score() == maxFrameScore
}

// IsStrike determines if a frame is a strike.
func (frame *Frame) IsStrike() bool {
	return frame.rolls[0] == maxFrameScore
}

// Bonus calculates bonus of a frame given its next frame in the game.
func (frame *Frame) Bonus(nextFrames ...Frame) int {
	if frame.IsStrike() {
		return frame.calculateBonusForStrike(nextFrames...)
	}
	if frame.IsSpare() {
		return frame.calculateBonusForSpare(nextFrames[0])
	}
	return 0
}

func (frame *Frame) calculateBonusForSpare(nextFrame Frame) int {
	return nextFrame.rolls[0]
}

func (frame *Frame) calculateBonusForStrike(nextFrames ...Frame) int {
	return nextFrames[0].Score()
}
