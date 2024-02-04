package noGracias

import "math/rand"

type Table struct {
	deck     []Card
	Tokens   int
	PlayCard Card
}

func NewTable() Table {
	deck := []Card{}
	for i := 3; i <= 35; i++ {
		deck = append(deck, NewCard(int64(i)))
	}
	return Table{
		deck:   deck,
		Tokens: 0,
	}
}

func (table Table) DeckShuffle(seed int64) {
	r := rand.New(rand.NewSource(seed))
	for i := 0; i < 250; i++ {
		rand1 := r.Int63n(33)
		rand2 := r.Int63n(33)
		support := table.deck[rand1]
		table.deck[rand1] = table.deck[rand2]
		table.deck[rand2] = support
	}
}

func (table *Table) RemoveCards(number int) {
	supp := table.deck
	table.deck = supp[number:]
}
