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

func TestFrameIsNotAlteredByARollWhenItIsComplete(t *testing.T) {
	frame := new(Frame)
	expectedRolls := [2]int{1, 2}

	frame.Roll(1)
	frame.Roll(2)
	frame.Roll(5)

	if !frame.IsComplete() {
		t.Error("Frame must be complete after 2 rolls, but it isn't.")
	}

	actualRolls := frame.rolls
	if actualRolls != expectedRolls {
		t.Errorf("Expected rolls to be %q. Got %q instead.", expectedRolls, actualRolls)
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

func TestReturnsFrameScore(t *testing.T) {
	expectedScore := 9
	frame := new(Frame)

	frame.Roll(5)
	frame.Roll(4)

	score := frame.Score()
	if score != expectedScore {
		t.Errorf("Expected frame score to be %d. Got %d instead.", expectedScore, score)
	}
}

func TestIsSpareWhenFirstRollIsLessThanTenButRollsSumTen(t *testing.T) {
	frame := new(Frame)

	frame.Roll(0)
	frame.Roll(10)

	if !frame.IsSpare() {
		t.Error("Expected frame to be a spare, but it isn't.")
	}
}

func TestIsNotSpare(t *testing.T) {
	cases := [][2]int{
		{0, 0},
		{10, 0},
		{2, 3},
		{7, 2},
	}

	for _, rolls := range cases {
		frame := new(Frame)

		for _, pinsDown := range rolls {
			frame.Roll(pinsDown)
		}

		if frame.IsSpare() {
			t.Error("Expected frame not to be a spare, but it is.")
		}
	}
}

func TestCalculatesBonusOfASpareWhenGivenNextFrame(t *testing.T) {
	nextFrame := Frame{[2]int{1, 2}, 2}
	frame := &Frame{[2]int{2, 8}, 2}
	expectedBonus := 1

	bonus := frame.Bonus(nextFrame)

	if bonus != expectedBonus {
		t.Errorf("Expected bonus to be %d. Got %d instead.", expectedBonus, bonus)
	}
}

func TestBonusIsZeroWhenFrameIsNotASpare(t *testing.T) {
	nextFrame := Frame{[2]int{1, 2}, 2}
	frame := &Frame{[2]int{2, 6}, 2}

	if frame.IsSpare() {
		t.Error("Expected frame not to be a spare, but it is.")
	}

	bonus := frame.Bonus(nextFrame)

	if bonus != 0 {
		t.Errorf("Expected bonus to be zero. Got %d instead.", bonus)
	}
}

func TestAStrikeIsCompleteAfterFirstRoll(t *testing.T) {
	frame := new(Frame)

	frame.Roll(10)

	if !frame.IsComplete() {
		t.Error("Expected frame to be complete, but it isn't.")
	}
}

func TestFrameIsStrikeWhenFirstRollKnocksDownAllPins(t *testing.T) {
	frame := new(Frame)

	frame.Roll(10)

	if !frame.IsStrike() {
		t.Error("Expected frame to be a strike, but it isn't.")
	}
}

func TestIsNotStrike(t *testing.T) {
	cases := [][2]int{
		{0, 0},
		{0, 10},
		{2, 3},
		{7, 3},
	}

	for _, rolls := range cases {
		frame := new(Frame)

		for _, pinsDown := range rolls {
			frame.Roll(pinsDown)
		}

		if frame.IsStrike() {
			t.Error("Expected frame not to be a strike, but it is.")
		}
	}
}

func TestBonusIsTheScoreOfTheNextFrameWhenFrameIsStrike(t *testing.T) {
	nextFrame := new(Frame)
	nextFrame.Roll(5)
	nextFrame.Roll(4)
	frame := new(Frame)
	frame.Roll(10)
	expectedBonus := 9

	actualBonus := frame.Bonus(*nextFrame, Frame{})

	if actualBonus != expectedBonus {
		t.Errorf("Expected bonus to be %d. Got %d instead.", expectedBonus, actualBonus)
	}
}

func TestBonusIsCalculatedWhenThereAreTwoStrikesInARow(t *testing.T) {
	nextFrame := new(Frame)
	nextFrame.Roll(10)
	secondNextFrame := new(Frame)
	secondNextFrame.Roll(2)
	secondNextFrame.Roll(6)
	frame := new(Frame)
	frame.Roll(10)
	expectedBonus := 12

	actualBonus := frame.Bonus(*nextFrame, *secondNextFrame)

	if actualBonus != expectedBonus {
		t.Errorf("Expected bonus to be %d. Got %d instead.", expectedBonus, actualBonus)
	}

}
