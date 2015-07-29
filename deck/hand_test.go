package deck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHand(t *testing.T) {
	cards := make([]Card, 3)
	hand := Hand{cards}
	assert.Equal(t, hand.NumberOfCards(), 3, "These should be equal")
}
func TestDealFromHand(t *testing.T) {
	cards := make([]Card, 3)
	hand := Hand{cards}
	card := hand.DealCard(0)
	assert.NotNil(t, card, "Card should be returned by dealing from hand.")
	assert.Equal(t, hand.NumberOfCards(), 2, "These should be equal")
}
func TestHandString(t *testing.T) {
	cards := make([]Card, 3)
	hand := Hand{cards}
	assert.Equal(t, hand.String(), "A♣ A♣ A♣", "These should be equal")
}
