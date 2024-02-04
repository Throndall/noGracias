package noGracias

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewContext(t *testing.T) {
	r := require.New(t)
	names := []string{"P1", "P2", "P3"}
	seed := int64(1)
	context := NewContext(names, seed)

	r.Len(context.Players, len(names))
	r.EqualValues(context.Table.Tokens, 0)
	context.Players[0].Bet(&context.Table)	
}

func TestNewPlayer(t *testing.T) {
	r := require.New(t)
	player := NewPlayer("P1")
	r.Len(player.Cards, 0)
	r.EqualValues(player.tokens, 11)
}

func TestNewTable(t *testing.T) {
	r := require.New(t)

	table := NewTable()
	r.EqualValues(table.Tokens, 0)
	table.DeckShuffle(int64(96))
	r.Len(table.deck, 33)
	r.NotEqualValues(table.deck[0].Number, 3)
	r.NotEqualValues(table.deck[32].Number, 35)
	table.RemoveCards(9)
	r.Len(table.deck, 24)
}

