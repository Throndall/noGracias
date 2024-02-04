package noGracias

type Card struct {
	Number int64
}

func NewCard(number int64) Card {
	return Card{
		Number: number,
	}
}


type Context struct {
	Table   Table
	Players []Player
	Seed    int64
}

func NewContext(names []string, seed int64) Context {
	players := []Player{}
	for i := 0; i < len(names); i++ {
		players = append(players, NewPlayer(names[i]))
	}
	table := NewTable()

	return Context{
		Players: players,
		Table:   table,
		Seed:    seed,
	}
}
