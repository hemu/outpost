package card

import (
// "fmt"
// "log"
)

type CardType int

const (
	ACTION CardType = iota
	TREASURE
	VICTORY
	CURSE
)

var CardFactory = map[string]Card{
	"Chapel":          {Name: "Chapel", Cost: 2, Ctype: ACTION},
	"Courtyard":       {Name: "Courtyard", Cost: 2, Ctype: ACTION},
	"Haven":           {Name: "Haven", Cost: 2, Ctype: ACTION},
	"Fishing Village": {Name: "Fishing Village", Cost: 3, Ctype: ACTION},
	"Village":         {Name: "Village", Cost: 3, Ctype: ACTION},
	"Warehouse":       {Name: "Warehouse", Cost: 3, Ctype: ACTION},
	"Moneylender":     {Name: "Moneylender", Cost: 4, Ctype: ACTION},
	"Monument":        {Name: "Monument", Cost: 4, Ctype: ACTION},
	"Navigator":       {Name: "Navigator", Cost: 4, Ctype: ACTION},
	"Bank":            {Name: "Bank", Cost: 7, Ctype: ACTION},
	"Copper":          {Name: "Copper", Cost: 0, Ctype: ACTION},
	"Silver":          {Name: "Silver", Cost: 3, Ctype: ACTION},
	"Gold":            {Name: "Gold", Cost: 6, Ctype: ACTION},
	"Estate":          {Name: "Estate", Cost: 2, Ctype: ACTION},
	"Duchy":           {Name: "Duchy", Cost: 5, Ctype: ACTION},
	"Province":        {Name: "Province", Cost: 8, Ctype: ACTION},
	"Curse":           {Name: "Curse", Cost: 0, Ctype: ACTION},
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
