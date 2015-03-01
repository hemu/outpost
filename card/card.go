package card

import (
	"fmt"
	// "log"
	"errors"
)

type CardType int

const (
	ACTION CardType = iota
	TREASURE
	VICTORY
	CURSE
)

var cardFactory = map[string]Card{
	"Chapel":          {Name: "Chapel", Cost: 2, Ctype: ACTION},
	"Courtyard":       {Name: "Courtyard", Cost: 2, Ctype: ACTION},
	"Haven":           {Name: "Haven", Cost: 2, Ctype: ACTION},
	"Fishing Village": {Name: "Fishing Village", Cost: 3, Ctype: ACTION},
	"Village":         {Name: "Village", Cost: 3, Ctype: ACTION},
	"Warehouse":       {Name: "Warehouse", Cost: 3, Ctype: ACTION},
	"Moneylender":     {Name: "Moneylender", Cost: 4, Ctype: ACTION},
	"Monument":        {Name: "Monument", Cost: 4, Ctype: ACTION},
	"Navigator":       {Name: "Navigator", Cost: 4, Ctype: ACTION},
	"Bank":            {Name: "Bank", Cost: 7, Ctype: TREASURE},
	"Copper":          {Name: "Copper", Cost: 0, Ctype: TREASURE},
	"Silver":          {Name: "Silver", Cost: 3, Ctype: TREASURE},
	"Gold":            {Name: "Gold", Cost: 6, Ctype: TREASURE},
	"Estate":          {Name: "Estate", Cost: 2, Ctype: VICTORY},
	"Duchy":           {Name: "Duchy", Cost: 5, Ctype: VICTORY},
	"Province":        {Name: "Province", Cost: 8, Ctype: VICTORY},
	"Curse":           {Name: "Curse", Cost: 0, Ctype: CURSE},
}

// {actions, buys, coins, victory}
var cardStats = map[string][4]int{
	"Chapel":          [4]int{0, 0, 0, 0},
	"Courtyard":       [4]int{0, 0, 0, 0},
	"Haven":           [4]int{1, 0, 0, 0},
	"Fishing Village": [4]int{2, 0, 1, 0},
	"Village":         [4]int{2, 0, 0, 0},
	"Warehouse":       [4]int{1, 0, 0, 0},
	"Moneylender":     [4]int{0, 0, 0, 0},
	"Monument":        [4]int{0, 0, 2, 1},
	"Navigator":       [4]int{0, 0, 0, 0},
	"Bank":            [4]int{0, 0, 0, 0},
	"Copper":          [4]int{0, 0, 1, 0},
	"Silver":          [4]int{0, 0, 2, 0},
	"Gold":            [4]int{0, 0, 3, 0},
	"Estate":          [4]int{0, 0, 0, 1},
	"Duchy":           [4]int{0, 0, 0, 3},
	"Province":        [4]int{0, 0, 0, 6},
	"Curse":           [4]int{0, 0, 0, 0},
}

func NewCard(name string) Card {
	if c, ok := cardFactory[name]; ok {
		newCard := Card{Name: c.Name, Cost: c.Cost, Ctype: c.Ctype}
		return newCard
	} else {
		return Card{Name: "UNKNOWN"}
	}
}

func NewCards(name string, num int) []Card {
	cards := []Card{}
	for i := 0; i < num; i++ {
		if c, ok := cardFactory[name]; ok {
			newCard := Card{Name: c.Name, Cost: c.Cost, Ctype: c.Ctype}
			cards = append(cards, newCard)
		} else {
			cards = append(cards, Card{Name: "UNKNOWN"})
		}
	}
	return cards
}

func GetCardStats(name string) ([4]int, error) {
	if c, ok := cardStats[name]; ok {
		return c, nil
	} else {
		return c, errors.New(fmt.Sprintf("Could not find stats for card %v", name))
	}
}

type Card struct {
	Name  string
	Cost  int
	Ctype CardType
}

type CardSet struct {
	Num  int
	Card Card
}
