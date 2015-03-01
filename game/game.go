package game

import (
	// "fmt"
	mCard "github.com/hmuar/dominion-replay/card"
	mEngine "github.com/hmuar/dominion-replay/engine"
	mEvent "github.com/hmuar/dominion-replay/event"
	mHistory "github.com/hmuar/dominion-replay/history"
	mState "github.com/hmuar/dominion-replay/state"
)

// turn-indexed array of game states
// stores gamestate at beginning of every turn
// turn 0 is start of game
type Game struct {
	states []mState.State
}

func NewGameBuilder() gameBuilder {
	g := Game{}
	gb := gameBuilder{game: g}
	state := mState.State{}
	gb.state = &state
	return gb
}

// gameBuilder uses an engine and feeds it actions
// and a current state and allows engine to modify
// that state
type gameBuilder struct {
	game  Game
	eng   mEngine.Engine
	state *mState.State
}

func (gb *gameBuilder) FeedHistory(h mHistory.History) {
	h.Print()
	gb.registerInitHistory(h)
	for _, turn := range h.Turns {
		gb.registerTurnStart(turn.GetTurnNum())
		for j := 0; j < turn.GetNumPlayerTurns(); j++ {
			player := h.PlayerOrder[j]
			gb.registerPlayerTurnStart(player)
			pturn := turn.GetPlayerEvents(j)
			for _, ev := range pturn {
				// fmt.Println(ev)
				gb.eng.RegisterEvent(ev, gb.state)
			}
			gb.registerPlayerTurnEnd(player)
		}
		gb.state.Print()
		return
	}
}

// extract init info about game such as initial supply,
// start draw pile (7 copper, 3 estate), players
func (gb *gameBuilder) registerInitHistory(h mHistory.History) {
	gb.registerSupply(h.Supply)
	gb.registerPlayers(h.PlayerOrder)
	//XXX: for now init draw deck is harcoded to
	//     7 coppers and 3 estates, but eventually
	//     read log to get this information
	initDraw := []mCard.Card{}
	initDraw = append(initDraw, mCard.NewCards("Copper", 7)...)
	initDraw = append(initDraw, mCard.NewCards("Estate", 3)...)
	for _, player := range h.PlayerOrder {
		gb.registerInitDraw(player, initDraw)
	}
}

func (gb *gameBuilder) registerSupply(cards []mCard.CardSet) {
	gb.eng.SetSupply(cards, gb.state)
}

func (gb *gameBuilder) registerPlayers(players []string) {
	gb.eng.SetPlayers(players, gb.state)
}

func (gb *gameBuilder) registerInitDraw(player string,
	cards []mCard.Card) {
	gb.eng.SetInitDraw(player, cards, gb.state)
}

func (gb *gameBuilder) registerPlayerTurnStart(player string) {
	gb.state.TurnPlayer = player
}

func (gb *gameBuilder) registerPlayerTurnEnd(player string) {
	endEvent := mEvent.Event{Player: player,
		Action: mEvent.ACTION_END_TURN,
		Cards:  []mCard.Card{},
	}
	gb.eng.RegisterEvent(endEvent, gb.state)
}

func (gb *gameBuilder) registerTurnStart(num int) {
	gb.state.TurnNum = num
}
