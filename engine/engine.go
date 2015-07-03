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

func (eng *Engine) SetPlayers(players []string, s *mState.State) {
	s.SetPlayers(players)
}

func (eng *Engine) SetInitDraw(player string, cards []mCard.Card, s *mState.State) {
	s.SetDraw(player, cards)
}

func (eng *Engine) RegisterPlayerTurnStart(player string, s *mState.State) {
	s.TurnPlayer = player
	// initialize player stats
	s.ResetPlayerStats(player)

}

func (eng *Engine) RegisterEvent(ev mEvent.Event, s *mState.State) error {

	var err error

	switch ev.Action {

	case mEvent.ACTION_DRAW:
		s.SetHand(ev.Player, ev.Cards)
		s.RemoveFromDraw(ev.Player, ev.Cards)

	case mEvent.ACTION_PLAY:
		for _, card := range ev.Cards {
			if card.Ctype == mCard.ACTION {
				s.AddPlayerActionCount(ev.Player, -1)
			}
			s.AddPlay(card)
			err = s.RemoveFromHand(card)
			stats, err := mCard.GetCardStats(card.Name)
			if err == nil {
				s.AddPlayerStats(ev.Player, stats)
			}
		}

	case mEvent.ACTION_BUY:
		for _, card := range ev.Cards {
			s.RemoveSupplyCard(card.Name)
		}

	case mEvent.ACTION_GAIN:
		for _, card := range ev.Cards {
			s.AddDiscardCard(ev.Player, card)
		}

	case mEvent.ACTION_DISCARD:
		for _, card := range ev.Cards {
			s.AddDiscardCard(ev.Player, card)
		}

		// Turn ended, so:
		// 1. move hand and play cards to discard
		// 2. set actions to 0
		// 3. set buys to 0
		// 4. set coints to 0
	case mEvent.ACTION_END_TURN:
		handCards := s.GetHand(ev.Player)
		playCards := s.GetPlay()
		for _, card := range handCards {
			s.AddDiscardCard(ev.Player, card)
		}
		for _, card := range playCards {
			s.AddDiscardCard(ev.Player, card)
		}
		s.SetHand(ev.Player, []mCard.Card{})
		s.SetPlay([]mCard.Card{})
		s.ClearPlayerStats(ev.Player)

	case mEvent.ACTION_SHUFFLE:

	}

	return err

}
