package bowlingkata

import "testing"

func TestGameStartsWithZeroScore(t *testing.T) {
	game := New()
	if game.score != 0 {
		t.Errorf("The initial score of a game is zero. Got %d instead.", game.score)
	}
}

func TestCalculatesGameScoreWhenNoSparesNorStrikes(t *testing.T) {
	rolls := []int{5, 3, 4, 4, 6, 2, 7, 1, 1, 7, 2, 6, 3, 5, 4, 4, 8, 0, 0, 8}
	expectedScore := 80
	game := New()

	for _, pinsDown := range rolls {
		game.Roll(pinsDown)
	}

	score := game.Score()
	if score != expectedScore {
		t.Errorf("Expected score to be %d. Actual is %d.", expectedScore, score)
	}
}
