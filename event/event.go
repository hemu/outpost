package event

import (
	mCard "github.com/hmuar/dominion-replay/card"
)

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
	ACTION_PLACE_ON_DECK string = "place"
	ACTION_LOOK_AT       string = "look"
	ACTION_TRASH         string = "trash"
)

type Event struct {
	Player string
	Action string
	Cards  []mCard.CardSet
}
