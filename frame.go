package bowlingkata

// Frame represents every frame in a bowling game.
type Frame struct {
	rolls    [2]int
	nextRoll int
}

// Roll annotates the pins taken down in a frame roll.
func (frame *Frame) Roll(pinsDown int) {
	frame.rolls[frame.nextRoll] = pinsDown
	frame.nextRoll++
}
