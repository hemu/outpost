package main

import (
	"fmt"
	mLogParse "github.com/hmuar/dominion-replay/logparse"
	// "github.com/gorilla/mux"
	"encoding/json"
	mGame "github.com/hmuar/dominion-replay/game"
	mMsg "github.com/hmuar/dominion-replay/message"
	mState "github.com/hmuar/dominion-replay/state"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"net/http"
)

// var game mGame.Game

func turnHandler(msg mMsg.Msg, game mGame.Game) mState.State {
	// fmt.Printf("%T", data)
	// turnNum := data[mMsg.KEY_DATA_TURN_NUM]
	// playerNum := data[mMsg.KEY_DATA_PLAYER_NUM]
	turnNumFloat := msg.MData["num"].(float64)
	playerNumFloat := msg.MData["pnum"].(float64)
	turnNumInt := int(turnNumFloat)
	playerNumInt := int(playerNumFloat)
	fmt.Printf("turn %v\n", turnNumInt)
	fmt.Printf("playerTurn %v\n", playerNumInt)
	// fmt.Printf("gameState length: %v\n", len(game.GetState(0, 0)))
	// return mState.State{}
	return game.GetState(turnNumInt, playerNumInt)
}

func socketMsgHandler(session sockjs.Session) {
	fmt.Println("socketMsgHandlerrrrrrrr")
	logFile := "logparse/test/testlogs/testlog_turns.txt"
	history := mLogParse.ParseLog(logFile)
	gameBuilder := mGame.NewGameBuilder()
	gameBuilder.FeedHistory(history)
	fmt.Println("[MSG] Received message")
	msg := &mMsg.Msg{}
	// var msg map[string]string
	for {
		if rawMsg, err := session.Recv(); err == nil {
			fmt.Println(rawMsg)
			// turnHandler(rawMsg)
			json.Unmarshal([]byte(rawMsg), &msg)
			fmt.Printf("%+v\n", msg)
			fmt.Println(msg.MData)
			// msgType := msg.MType
			// msgTypeKey := mMsg.KEY_MSG_TYPE
			// msgType := msg[msgTypeKey]
			switch msg.MType {
			case mMsg.KEY_MSG_TYPE_TURN:
				data := turnHandler(*msg, gameBuilder.GetGame())
				m, _ := json.Marshal(data)
				session.Send(string(m))
			}
		}
		break
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "app/index.html")
}

func main() {
	// logparse.ParseLog("logparse/test/testlogs/testlog_turns.txt")
	router := httprouter.New()
	// handle all static files
	router.ServeFiles("/public/*filepath", http.Dir("app"))
	// socket handler
	router.Handler(
		"GET",
		"/echo/*subpath",
		sockjs.NewHandler("/echo", sockjs.DefaultOptions, socketMsgHandler),
	)
	// root page request
	router.GET("/", indexHandler)
	log.Println("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
