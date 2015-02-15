package engine

// Feed engine an action, and a state, and
// engine knows how to interpret action and
// modify the state

import (
	mCard "github.com/hmuar/dominion-replay/card"
	mEvent "github.com/hmuar/dominion-replay/event"
	mState "github.com/hmuar/dominion-replay/state"
)

type Engine struct {
}

func (eng *Engine) SetSupply(cards []mCard.CardSet, s *mState.State) {
	s.SetSupply(cards)
}

func (eng *Engine) RegisterEvent(ev mEvent.Event, s *mState.State) {

	switch ev.Action {

	case mEvent.ACTION_DRAW:
		s.SetHand(ev.Player, ev.Cards)

	case mEvent.ACTION_PLAY:
		for _, cardSet := range ev.Cards {
			s.AddPlay(ev.Player, cardSet)
		}
	}
}
