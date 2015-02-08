package game

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

type event struct {
	Player string
	Action string
	Cards  []mCard.CardSet
}

type playerTurn struct {
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

type Game struct {
	LogFile string
	Players []string
	Supply  []mCard.CardSet
	Turns   []turn
	Rating  string
	Winner  string
}

func (g *Game) PrintGame() {
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

type GameBuilder struct {
	Game Game
}

// factory function gameBuilder
func NewGameBuilder() GameBuilder {
	game := Game{}
	gb := GameBuilder{Game: game}
	return gb
}

func (gb *GameBuilder) SetSupply(cards []mCard.CardSet) {
	gb.Game.Supply = cards
}

func (gb *GameBuilder) getCurTurn() *turn {
	if len(gb.Game.Turns) > 0 {
		return &gb.Game.Turns[len(gb.Game.Turns)-1]
	} else {
		return &turn{}
	}
}

func (gb *GameBuilder) getCurPlayerTurn() *playerTurn {
	return gb.getCurTurn().getCurPlayerTurn()
}

// this should never get called directly by client.
// should be controlled by startNewPlayerTurn.
func (gb *GameBuilder) startNewTurn(turnNum int) {
	gb.Game.Turns = append(gb.Game.Turns, turn{num: turnNum})
}

func (gb *GameBuilder) startNewPlayerTurn(player string, turnNum int) {
	// if turnNum is diff than current turn num,
	// then need to start a new turn
	if gb.getCurTurn().num != turnNum {
		gb.startNewTurn(turnNum)
	}
	curTurn := gb.getCurTurn()
	curTurn.addPlayerTurn(playerTurn{player: player})
}

func (gb *GameBuilder) StartPlayerTurn(player string, turnNum int) {
	gb.startNewPlayerTurn(player, turnNum)
}

func (gb *GameBuilder) AddEvent(player string, action string, cards []mCard.CardSet) {
	curPlayerTurn := gb.getCurPlayerTurn()
	curPlayerTurn.addEvent(player, action, cards)
}
