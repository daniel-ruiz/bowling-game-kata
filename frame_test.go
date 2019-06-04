package bowlingkata

import "testing"

func TestInitializesFrame(t *testing.T) {
	frame := new(Frame)

	if frame.nextRoll != 0 {
		t.Errorf("Next roll in a new frame must be zero. Got %d instead.", frame.nextRoll)
	}

	if rollsCount := len(frame.rolls); rollsCount != 2 {
		t.Errorf("A frame only has 2 rolls. Found %d instead.", rollsCount)
	}
}

func TestUpdatesFrameWhenRolling(t *testing.T) {
	pinsDown := 5
	frame := new(Frame)

	frame.Roll(pinsDown)

	if frame.nextRoll != 1 {
		t.Errorf("Roll count in a frame after first roll must be 1. Got %d instead.", frame.nextRoll)
	}

	firstRoll := frame.rolls[0]
	if firstRoll != pinsDown {
		t.Errorf("Number of pins down in first roll must be %d. Got %d instead.", pinsDown, firstRoll)
	}
}
