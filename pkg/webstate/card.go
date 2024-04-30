package webstate

type Cards = []Card

type Card struct {
	Title  string
	Artist string
}

func NewCard(name, artist string) Card {
	return Card{
		Title:  name,
		Artist: artist,
	}
}

