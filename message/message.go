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

const (
	KEY_MSG_TYPE string = "mtype"
	KEY_MSG_DATA string = "mdata"
)

const (
	KEY_MSG_TYPE_TURN string = "turn"
)

const (
	KEY_DATA_TURN_NUM   string = "num"
	KEY_DATA_PLAYER_NUM string = "pnum"
)

type Msg struct {
	MType string                 `json:"mtype",string`
	MData map[string]interface{} `json:"mdata",string`
}

//, cards []card.Card
func GetMsgActionDraw(turn string, player string) []byte {
	msg := map[string]string{KEY_TURN: turn, KEY_PLAYER: player, KEY_EVENT: KEY_EVENT_ACTION}
	jsonMsg, _ := json.Marshal(msg)
	return jsonMsg
}

// func GetMsgActionPlay() {
// }
