package game

import "testing"

func TestGameStartsWithZeroScore(t *testing.T) {
	game := New()
	if game.score != 0 {
		t.Errorf("The initial score of a game is zero. Got %d instead.", game.score)
	}
}
