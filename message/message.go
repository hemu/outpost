package message

import (
	"encoding/json"
	// "github.com/hmuar/dominion-replay/card"
)

const (
	KEY_PLAYER string = "player"
	KEY_TURN   string = "turn"
	KEY_EVENT  string = "event"
)

const (
	KEY_EVENT_ACTION string = "action"
)

const (
	KEY_ACTION_DRAW    string = "draw"
	KEY_ACTION_PLAY    string = "play"
	KEY_ACTION_BUY     string = "buy"
	KEY_ACTION_GAIN    string = "gain"
	KEY_ACTION_DISCARD string = "discard"
)

//, cards []card.Card
func GetMsgActionDraw(turn string, player string) []byte {
	msg := map[string]string{KEY_TURN: turn, KEY_PLAYER: player, KEY_EVENT: KEY_EVENT_ACTION}
	jsonMsg, _ := json.Marshal(msg)
	return jsonMsg
}

// func GetMsgActionPlay() {
// }
