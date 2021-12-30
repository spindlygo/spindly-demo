
package spindlyapp

import (
    "github.com/spindlygo/spindly/Spindly"
	"github.com/gorilla/mux"
)

var router *mux.Router

func Configure() {
	InitializeHubs()
	router = Spindly.NewRouter()
	Spindly.HandleHub(router, HubManager)
	Spindly.HandleStatic(router, "public", "index.html")
}

func Serve() {
	Spindly.Serve(router, "32510")
}

