package kernel

import (
	"code.google.com/p/go.net/websocket"
	"github.com/jmorgan1321/golang-games/core/support"
	"net/http"
)

type NetworkManager struct {
}

func (n *NetworkManager) StartUp(config GameObject) {
	// get ip-addr, port from config
	ipaddr, port := "127.0.0.1", "8080"
	// open websocket
	http.Handle("/ws/", websocket.Handler(wsHandler))
	http.Handle("/www/",
		http.StripPrefix("/www/", http.FileServer(http.Dir("./www"))))

	// start listening for new connections
	go func() {
		support.Log("opening websocket on: %s:%s", ipaddr, port)
		http.ListenAndServe(ipaddr+":"+port, nil)
	}()
}

func (n *NetworkManager) ShutDown() {
	// close websocket
}

func (n *NetworkManager) BeginFrame() {
	// process last frames messages, sending them to rest of game
}

func (n *NetworkManager) EndFrame() {
	// send out world state
}

func wsHandler(ws *websocket.Conn) {
	support.Log("incoming connection")
	//handle connection
}
