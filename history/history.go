package event

import (
	"fmt"
	mCard "github.com/hmuar/dominion-replay/card"
)

// type GameState struct {
//  logFile       string
//  playerDraw    []mCard.CardSet
//  playerDiscard []mCard.CardSet
//  playerHand    []mCard.CardSet
// }

// TODO:
// ACTION_REVEAL
// ACTION_RECEIVE
// ACTION_DURATION

const (
	ACTION_DRAW          string = "draw"
	ACTION_PLAY          string = "play"
	ACTION_BUY           string = "buy"
	ACTION_GAIN          string = "gain"
	ACTION_DISCARD       string = "discard"
	ACTION_SHUFFLE       string = "shuffle"
	ACTION_PLACE_ON_DECK string = "place-on-deck"
	ACTION_LOOK_AT       string = "look-at"
	ACTION_TRASH         string = "trash"
)

const (
	BUILD_STATE_INIT = iota
	BUILD_STATE_SETUP
	BUILD_STATE_TURN
	BUILD_STATE_END
)

type event struct {
	Player string
	Action string
	Cards  []mCard.CardSet
}

type playerTurn struct {
	num    int
	player string
	events []event
}

func (pt *playerTurn) GetEvents() []event {
	return pt.events
}

func (pt *playerTurn) addEvent(player string,
	action string,
	cards []mCard.CardSet) {
	newEvent := event{Player: player, Action: action, Cards: cards}
	pt.events = append(pt.events, newEvent)
}

type turn struct {
	num         int
	playerTurns []playerTurn
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

func (t *turn) GetPlayerEvents(playerInd int) []event {
	if playerInd > len(t.playerTurns) {
		return []event{}
	}
	return t.playerTurns[playerInd].events
}

func (t *turn) GetNumPlayerTurns() int {
	return len(t.playerTurns)
}

type History struct {
	LogFile string
	Players []string
	Supply  []mCard.CardSet
	Turns   []turn
	Rating  string
	Winner  string
}

func (g *History) PrintGame() {
	for _, turn := range g.Turns {
		fmt.Printf("******** Turn %d *******\n", turn.num)
		for _, pTurn := range turn.playerTurns {
			fmt.Printf("--- %v turn %d ---\n", pTurn.player, turn.num)
			for _, ev := range pTurn.events {
				fmt.Printf("%v ", ev.Action)
				fmt.Println(ev.Cards)
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
	// if turnNum is diff than current turn num,
	// then need to start a new turn
	if hb.getCurTurn().num != turnNum {
		hb.startNewTurn(turnNum)
	}
	curTurn := hb.getCurTurn()
	curTurn.addPlayerTurn(playerTurn{num: turnNum, player: player})
}

func (hb *HistoryBuilder) RegisterGameSetup() {
	hb.startNewTurn(0)
	hb.state = BUILD_STATE_SETUP
}

func (hb *HistoryBuilder) StartPlayerTurn(player string, turnNum int) {
	hb.startNewPlayerTurn(player, turnNum)
}

func (hb *HistoryBuilder) AddEvent(player string, action string, cards []mCard.CardSet) {
	if hb.state == BUILD_STATE_SETUP {
		// if it's shuffle during game setup, ignore
		if action == ACTION_SHUFFLE {
			return
		} else {
			hb.startNewPlayerTurn(player, 0)
		}
	}
	curPlayerTurn := hb.getCurPlayerTurn()
	curPlayerTurn.addEvent(player, action, cards)
}
