package deck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallDeckToString(t *testing.T) {
	deck := NewEmptyDeck()
	deck.Cards = append(deck.Cards, Card{ACE, HEART}, Card{KING, HEART})
	result := fmt.Sprintf("%s", deck)
	assert.Equal(t, "A♥\nK♥\n", result, "These should be equal")
}

func TestDeckToString(t *testing.T) {
	deck := NewDeck(false)
	result := fmt.Sprintf("%s", deck)
	assert.Equal(t, "A♣\n2♣\n3♣\n4♣\n5♣\n6♣\n7♣\n8♣\n9♣\nT♣\nJ♣\nQ♣\nK♣\n"+
		"A♦\n2♦\n3♦\n4♦\n5♦\n6♦\n7♦\n8♦\n9♦\nT♦\nJ♦\nQ♦\nK♦\n"+
		"A♥\n2♥\n3♥\n4♥\n5♥\n6♥\n7♥\n8♥\n9♥\nT♥\nJ♥\nQ♥\nK♥\n"+
		"A♠\n2♠\n3♠\n4♠\n5♠\n6♠\n7♠\n8♠\n9♠\nT♠\nJ♠\nQ♠\nK♠\n", result, "These should be equal")
}

func TestDeckSignature(t *testing.T) {
	deck := NewEmptyDeck()
	deck.Cards = append(deck.Cards, Card{ACE, HEART}, Card{KING, HEART}, Card{TEN, CLUB})
	result := deck.GetSignature()
	assert.Equal(t, "02c290", result, "These should be equal")
}

func TestCompleteDeckSignature(t *testing.T) {
	deck := NewDeck(false)
	result := deck.GetSignature()
	assert.Equal(t, "00102030405060708090a0b0c001112131415161718191a1b1c102122232425262728292a2b2c203132333435363738393a3b3c3", result, "These should be equal")
}

func TestNewTinyDeck(t *testing.T) {
	deck := NewSpecificDeck(true, FACES, []Suit{SPADE})
	assert.Equal(t, 13, deck.NumberOfCards())
}

func TestNewSmallDeck(t *testing.T) {
	deck := NewSpecificDeck(false, FACES, []Suit{SPADE, HEART})
	assert.Equal(t, 26, deck.NumberOfCards())
}

func TestGetCardFromDeck(t *testing.T) {
	deck := NewDeck(false)
	card := deck.GetCard(1)
	assert.Equal(t, 52, deck.NumberOfCards())
	assert.NotNil(t, card)
}

func TestDrawFromDeck(t *testing.T) {
	deck := NewDeck(false)
	deck.Draw()
	assert.Equal(t, 51, deck.NumberOfCards())
	deck.Draw()
	assert.Equal(t, 50, deck.NumberOfCards())
}
func TestDrawFromDeckGetCard(t *testing.T) {
	deck := NewDeck(false)
	card := deck.Draw()
	assert.NotNil(t, card, "Draw should return a Card")
}

func TestDrawHandFromDeck(t *testing.T) {
	deck := NewDeck(false)

	hand := deck.DrawHand(4)
	assert.Equal(t, 48, deck.NumberOfCards())
	assert.Equal(t, 4, hand.NumberOfCards())

	otherHand := deck.DrawHand(6)
	assert.Equal(t, 42, deck.NumberOfCards())
	assert.Equal(t, 6, otherHand.NumberOfCards())
}

func BenchmarkTinyDeckShuffle(b *testing.B) {
	deck := NewSpecificDeck(false, FACES, []Suit{SPADE})
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}

func BenchmarkSmallDeckShuffle(b *testing.B) {
	deck := NewSpecificDeck(false, FACES, []Suit{SPADE, HEART})
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}

func BenchmarkMediumDeckShuffle(b *testing.B) {
	deck := NewSpecificDeck(false, FACES, []Suit{SPADE, HEART, DIAMOND})
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}

func BenchmarkDeckShuffle(b *testing.B) {
	deck := NewSpecificDeck(false, FACES, SUITS)
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}
