package noGracias

import "strconv"

type Player struct {
	Cards  []Card
	Name   string
	tokens int
}

func NewPlayer(name string) Player {
	return Player{
		Name:   name,
		tokens: 11,
	}
}

func (player *Player) TakeTable(table *Table) {
	player.tokens += table.Tokens
	player.Cards = append(player.Cards, table.PlayCard)
	table.PlayCard = Card{}
	table.Tokens = 0
}

func (player *Player) Bet(table *Table) {
	player.tokens--
	table.Tokens++
}

func (player Player) GetState() []string {
	cards := ""
	points := int64(0)
	for i := 0; i < len(player.Cards); i++ {
		cards = cards + " - " + strconv.FormatInt(player.Cards[i].Number, 10)
		points += player.Cards[i].Number
	}
	state := []string{cards, "Total: " + strconv.FormatInt(points, 10)}
	return state
}
