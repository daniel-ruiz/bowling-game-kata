package bowlingkata

import "testing"

func TestGameStartsWithZeroScore(t *testing.T) {
	game := new(Game)
	if game.score != 0 {
		t.Errorf("The initial score of a game is zero. Got %d instead.", game.score)
	}
}

func TestCalculatesGameScoreWhenNoSparesNorStrikes(t *testing.T) {
	rolls := []int{5, 3, 4, 4, 6, 2, 7, 1, 1, 7, 2, 6, 3, 5, 4, 4, 8, 0, 0, 8}
	expectedScore := 80
	game := new(Game)

	for _, pinsDown := range rolls {
		game.Roll(pinsDown)
	}

	score := game.Score()
	if score != expectedScore {
		t.Errorf("Expected score to be %d. Actual is %d.", expectedScore, score)
	}
}

func TestCalculatesGameScoreWhenThereIsASpare(t *testing.T) {
	rolls := []int{5, 5, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	expectedScore := 16
	game := new(Game)

	for _, pinsDown := range rolls {
		game.Roll(pinsDown)
	}

	score := game.Score()
	if score != expectedScore {
		t.Errorf("Expected score to be %d. Actual is %d.", expectedScore, score)
	}
}

func TestCalculatesGameScoreWhenThereAreManySparesInARow(t *testing.T) {
	rolls := []int{5, 5, 3, 7, 6, 4, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	expectedScore := 43
	game := new(Game)

	for _, pinsDown := range rolls {
		game.Roll(pinsDown)
	}

	score := game.Score()
	if score != expectedScore {
		t.Errorf("Expected score to be %d. Actual is %d.", expectedScore, score)
	}
}

// func TestCalculatesGameScoreWhenThereIsAStrike(t *testing.T) {
// 	rolls := []int{10, 3, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
// 	expectedScore := 26
// 	game := new(Game)

// 	for _, pinsDown := range rolls {
// 		game.Roll(pinsDown)
// 	}

// 	score := game.Score()

// 	if score != expectedScore {
// 		t.Errorf("Expected score to be %d. Actual is %d.", expectedScore, score)
// 	}
// }
