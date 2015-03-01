package main

import (
	"fmt"
	mLogParse "github.com/hmuar/dominion-replay/logparse"
	// "github.com/gorilla/mux"
	"encoding/json"
	mGame "github.com/hmuar/dominion-replay/game"
	mMsg "github.com/hmuar/dominion-replay/message"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"net/http"
)

var game mGame.Game

func turnHandler(msg mMsg.Msg) {
	// fmt.Printf("%T", data)
	// turnNum := data[mMsg.KEY_DATA_TURN_NUM]
	// playerNum := data[mMsg.KEY_DATA_PLAYER_NUM]
	turnNum := msg.MData["num"].(float64)
	playerNum := msg.MData["pnum"].(float64)
	fmt.Printf("turn %v\n", turnNum)
	fmt.Printf("playerTurn %v", playerNum)

}

func socketMsgHandler(session sockjs.Session) {
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
				turnHandler(*msg)
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

	logFile := "test/testlogs/testlog_turns.txt"
	history := mLogParse.ParseLog(logFile)
	gameBuilder := mGame.NewGameBuilder()
	gameBuilder.FeedHistory(history)

}
