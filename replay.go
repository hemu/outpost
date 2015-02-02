package main

import (
	"fmt"
	// "github.com/hmuar/dominion-replay/logparse"
	// "github.com/gorilla/mux"
	// "github.com/hmuar/dominion-replay/message"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"net/http"
)

func turnHandler() {
	turnNum := r.URL.Path[len("/replay/"):]
	fmt.Println("turn requested")
	fmt.Println(turnNum)
}

func socketMsgHandler(session sockjs.Session) {
	fmt.Println("[MSG] Received message")
	for {
		if msg, err := session.Recv(); err == nil {
			fmt.Println(msg)
			turnHandler(msg)
			// session.Send("for once in my liffffe")
			continue
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
