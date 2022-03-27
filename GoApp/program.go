package GoApp

import (
	"github.com/spindlygo/spindly-demo/spindlyapp"
	"github.com/spindlygo/spindly/Spindly"
)

func Main() {

	// OnInitialize is called before anything is initialized
	onInitialize()

	// Configure Hubs and Server Routes
	spindlyapp.Configure()

	// OnConfigure is called after host is configured
	onConfigure()

	// Start serving the UI. This will block until the server is stopped.
	spindlyapp.Serve(onServerStart)

	// OnExit is called after the server is stopped
	onExit()
}

// MainMobile is called from the mobile app (iOS or Android).
// This is the entry point of your mobile app.
// This function could be called multiple times, depending on the lifecycle of mobile app.
func MainMobile(currentWorkingDir string) {

	if Spindly.InitializeForMobile(defaultServerPort, currentWorkingDir) {
		// If there's another instance of the server running, we will exit this one
		return
	}

	Main()
}
