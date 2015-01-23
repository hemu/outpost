package main

import (
	// "fmt"
	// "github.com/hmuar/dominion-replay/card"
	"github.com/hmuar/dominion-replay/logparse"
)

/*

- log parser
- write a tester

*/

func main() {
	logparse.ParseLog("logparse/test/testlogs/testlog_turns.txt")
	// newCard := card.Card{Name: "hey", Cost: 4, Ctype: card.TREASURE}
}
