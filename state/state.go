package state

import (
	mCard "github.com/hmuar/dominion-replay/card"
)

type Player struct {
	Action int
	Buy    int
	Coin   int
	Point  int
	Deck   Deck
}

type Deck struct {
	Hand     []mCard.Card
	Draw     []mCard.Card
	Discard  []mCard.Card
	Duration []mCard.Card
	Play     []mCard.Card
}

type Board struct {
	Trash  []mCard.Card
	Supply []mCard.CardSet
}

type State struct {
	Players    map[string]*Player
	Board      Board
	TurnPlayer string
	TurnNum    int
}

func (s *State) SetPlayers(players []string) {
	s.Players = make(map[string]*Player)
	for _, player := range players {
		s.Players[player] = &Player{}
	}
}

func (s *State) getPlayer(player string) *Player {
	return s.Players[player]
}

func (s *State) SetSupply(cardSets []mCard.CardSet) {
	s.Board.Supply = cardSets
}

func (s *State) SetHand(player string, cards []mCard.Card) {
	p := s.getPlayer(player)
	p.Deck.Hand = cards
}

func (s *State) GetHand(player string) []mCard.Card {
	return s.getPlayer(player).Deck.Hand
}

func (s *State) AddPlay(player string, card mCard.Card) {
	p := s.getPlayer(player)
	p.Deck.Play = append(p.Deck.Play, card)
}
