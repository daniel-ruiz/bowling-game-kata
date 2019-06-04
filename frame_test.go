package bowlingkata

import "testing"

func TestInitializesFrame(t *testing.T) {
	frame := new(Frame)

	if frame.rollCount != 0 {
		t.Errorf("Next roll in a new frame must be zero. Got %d instead.", frame.rollCount)
	}

	if rollsCount := len(frame.rolls); rollsCount != 2 {
		t.Errorf("A frame only has 2 rolls. Found %d instead.", rollsCount)
	}
}

func TestUpdatesFrameWhenRolling(t *testing.T) {
	pinsDown := 5
	frame := new(Frame)

	frame.Roll(pinsDown)

	if frame.rollCount != 1 {
		t.Errorf("Roll count in a frame after first roll must be 1. Got %d instead.", frame.rollCount)
	}

	firstRoll := frame.rolls[0]
	if firstRoll != pinsDown {
		t.Errorf("Number of pins down in first roll must be %d. Got %d instead.", pinsDown, firstRoll)
	}
}

func TestFrameIsCompleteAfterTwoRolls(t *testing.T) {
	frame := new(Frame)

	frame.Roll(1)
	frame.Roll(2)

	if !frame.IsComplete() {
		t.Error("Frame must be complete after 2 rolls, but it isn't.")
	}
}

func TestFrameIsNotComplete(t *testing.T) {
	frame := new(Frame)

	if frame.IsComplete() {
		t.Error("Frame must not be complete when no rolls are made, but it is.")
	}

	frame.Roll(7)
	if frame.IsComplete() {
		t.Error("Frame must not be complete after the first roll, but it is.")
	}
}
