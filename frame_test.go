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
