package bowlingkata

// Frame represents every frame in a bowling game.
type Frame struct {
	rolls     [2]int
	rollCount int
}

// Roll annotates the pins taken down in a frame roll.
func (frame *Frame) Roll(pinsDown int) {
	frame.rolls[frame.rollCount] = pinsDown
	frame.rollCount++
}

// IsComplete determines if a frame has finished
func (frame *Frame) IsComplete() bool {
	return frame.rollCount == 2
}

// Score calculates total score of a frame.
func (frame *Frame) Score() int {
	score := 0

	for _, pinsDown := range frame.rolls {
		score += pinsDown
	}

	return score
}
