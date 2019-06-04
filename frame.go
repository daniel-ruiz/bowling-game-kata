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
