package spindlyapp

import (
	"github.com/gorilla/mux"
	"github.com/spindlygo/spindly/Spindly"
)

var router *mux.Router
var DefaultPort string = "0"
var HostURL string

func Configure() {
	InitializeHubs()
	router = Spindly.NewRouter()
	Spindly.HandleHub(router, HubManager)
	Spindly.HandleStatic(router, "public", "index.html")
}

func Serve(onServerStart func()) bool {
	HostURL, DefaultPort = Spindly.Serve(router, DefaultPort, onServerStart)

	if HostURL == "" {
		return false
	}

	Spindly.BlockWhileHostRunning()

	return true
}
