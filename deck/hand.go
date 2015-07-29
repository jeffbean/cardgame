package deck

import "strings"

//Hand is an array of cards
type Hand struct {
	Cards []Card
}

//DealCard deals a card from a hand and removes it from the hand
func (h *Hand) DealCard(i int) Card {
	card := h.Cards[len(h.Cards)-1:]
	h.Cards = append(h.Cards[:i], h.Cards[i+1:]...)
	return card[0]
}

// NumberOfCards is a utility function that tells you how many cards are left in the deck
func (h *Hand) NumberOfCards() int {
	return len(h.Cards)
}

// String is the hand representation of all the cards
func (h *Hand) String() string {
	var cardsSymbol []string
	for _, card := range h.Cards {
		cardsSymbol = append(cardsSymbol, card.String())
	}
	return strings.Join(cardsSymbol, " ")
}
