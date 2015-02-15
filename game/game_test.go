package game

import (
	mParse "github.com/hmuar/dominion-replay/logparse"
	"testing"
)

func TestHistoryInit(t *testing.T) {
	history := mParse.ParseLog("../logparse/test/testlogs/testlog_turns.txt")
	history.Print()
	// test supply
	// eng := NewEngine()
	// eng.FeedHistory(history)
	// supply0 = eng.state[0].board.supply
	// expectSup := []string{"Chapel", "Courtyard", "Haven", "Fishing Village",
	//  "Village", "Warehouse", "Moneylender", "Monument", "Navigator", "Bank",
	//  "Copper", "Silver", "Gold", "Estate", "Duchy", "Province", "Curse"}
	// if len(supply0) != 18 {
	//  t.Errorf("Expected %d supply cards but got %d", 17, len(supply0))
	// }
	// for i, val := range expectSup {
	//  if state[0].Supply[i].Card.Name != val {
	//    t.Errorf("Expected %v supply card but got %v", val, history.Supply[i].Card.Name)
	//  }
	// }
}
