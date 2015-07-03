package event

import (
	"fmt"
	mCard "github.com/hmuar/dominion-replay/card"
	mEvent "github.com/hmuar/dominion-replay/event"
)

const (
	BUILD_STATE_INIT = iota
	BUILD_STATE_SETUP
	BUILD_STATE_TURN
	BUILD_STATE_END
)

type playerTurn struct {
	num    int
	player string
	events []mEvent.Event
}

func (pt *playerTurn) GetEvents() []mEvent.Event {
	return pt.events
}

func (pt *playerTurn) addEvent(player string,
	action string,
	cards []mCard.Card) {
	newEvent := mEvent.Event{Player: player, Action: action, Cards: cards}
	pt.events = append(pt.events, newEvent)
}

type turn struct {
	num         int
	playerTurns []playerTurn
}

func (t *turn) GetTurnNum() int {
	return t.num
}

func (t *turn) addPlayerTurn(pt playerTurn) {
	t.playerTurns = append(t.playerTurns, pt)
}

func (t *turn) getCurPlayerTurn() *playerTurn {
	numPlayerTurns := t.GetNumPlayerTurns()
	if numPlayerTurns > 0 {
		return &t.playerTurns[numPlayerTurns-1]
	} else {
		return &playerTurn{}
	}
}

func (t *turn) GetPlayerEvents(playerInd int) []mEvent.Event {
	if playerInd > len(t.playerTurns) {
		return []mEvent.Event{}
	}
	return t.playerTurns[playerInd].events
}

func (t *turn) GetNumPlayerTurns() int {
	return len(t.playerTurns)
}

type History struct {
	LogFile     string
	Players     []string
	Supply      []mCard.CardSet
	Turns       []turn
	PlayerOrder []string
	Rating      string
	Winner      string
}

func (h *History) Print() {
	fmt.Println("[Game Info]")
	fmt.Printf("  Logfile: %v\n", h.LogFile)
	fmt.Printf("  Players: %v\n", h.Players)
	fmt.Println("  Supply:")
	for _, cardSet := range h.Supply {
		fmt.Printf("    %v\n", cardSet)
	}
	fmt.Printf("  Rating: %v\n", h.Rating)
	fmt.Printf("  Winner: %v\n\n", h.Winner)
	fmt.Println("[Game Start]")
	for _, turn := range h.Turns {
		fmt.Printf("  [Turn %d]\n", turn.num)
		for _, pTurn := range turn.playerTurns {
			fmt.Printf("    [%v turn %d]\n", pTurn.player, pTurn.num)
			for _, ev := range pTurn.events {
				actString := ev.Action
				paddingSize := 8 - len(actString)
				if paddingSize > 0 {
					for i := 0; i < paddingSize; i++ {
						actString += " "
					}
				}
				fmt.Printf("      %v| %v\n", actString, ev.Cards)
			}
		}
		fmt.Println("")
	}
}

type HistoryBuilder struct {
	History History
	state   int
}

// factory function gameBuilder
func NewHistoryBuilder() HistoryBuilder {
	game := History{}
	hb := HistoryBuilder{History: game}
	hb.state = BUILD_STATE_INIT
	return hb
}

func (hb *HistoryBuilder) SetSupply(cards []mCard.CardSet) {
	hb.History.Supply = cards
}

func (hb *HistoryBuilder) getCurTurn() *turn {
	if len(hb.History.Turns) > 0 {
		return &hb.History.Turns[len(hb.History.Turns)-1]
	} else {
		return &turn{}
	}
}

func (hb *HistoryBuilder) getCurPlayerTurn() *playerTurn {
	return hb.getCurTurn().getCurPlayerTurn()
}

// this should never get called directly by client.
// should be controlled by startNewPlayerTurn.
func (hb *HistoryBuilder) startNewTurn(turnNum int) {
	hb.state = BUILD_STATE_TURN
	hb.History.Turns = append(hb.History.Turns, turn{num: turnNum})
}

func (hb *HistoryBuilder) startNewPlayerTurn(player string, turnNum int) {
	// if turn 1, use this to build up player order
	if turnNum == 1 {
		hb.History.PlayerOrder = append(hb.History.PlayerOrder, player)
	}
	// if turnNum is diff than current turn num,
	// then need to start a new turn
	if hb.getCurTurn().num != turnNum {
		hb.startNewTurn(turnNum)
	}
	curTurn := hb.getCurTurn()
	curTurn.addPlayerTurn(playerTurn{num: turnNum, player: player})
}

func (hb *HistoryBuilder) endPrevPlayerTurn(player string, turnNum int) {
	curPlayerTurn := hb.getCurPlayerTurn()
	if curPlayerTurn.player != "" && curPlayerTurn.num != 0 {
		curPlayerTurn.addEvent(player, mEvent.ACTION_END_TURN, []mCard.Card{})
	}
}

func (hb *HistoryBuilder) RegisterGameSetup() {
	hb.startNewTurn(0)
	hb.state = BUILD_STATE_SETUP
}

func (hb *HistoryBuilder) StartPlayerTurn(player string, turnNum int) {
	if turnNum == 1 && hb.getCurTurn().GetNumPlayerTurns() > 0 {
		if hb.getCurTurn().GetNumPlayerTurns() > 0 {
			curPlayerTurn := hb.getCurPlayerTurn()
			hb.endPrevPlayerTurn(curPlayerTurn.player, curPlayerTurn.num)
		}
	} else {
		curPlayerTurn := hb.getCurPlayerTurn()
		hb.endPrevPlayerTurn(curPlayerTurn.player, curPlayerTurn.num)
	}
	hb.startNewPlayerTurn(player, turnNum)
}

func (hb *HistoryBuilder) AddEvent(player string, action string, cards []mCard.Card) {
	if hb.state == BUILD_STATE_SETUP {
		// if it's shuffle during game setup, ignore
		if action == mEvent.ACTION_SHUFFLE {
			return
		} else {
			hb.startNewPlayerTurn(player, 0)
		}
	}
	curPlayerTurn := hb.getCurPlayerTurn()
	curPlayerTurn.addEvent(player, action, cards)
}
