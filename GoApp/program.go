package GoApp

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/spindlygo/spindly/Spindly"
)

func Main() {

	configure()
	runServer()
}

func MainMobile(currentWorkingDir string) {

	doubleRunningCheckCount := 8
	for Spindly.CheckIfAnotherSpindlyAppIsRunning(serverPort) {
		time.Sleep(500 * time.Millisecond)

		doubleRunningCheckCount--
		if doubleRunningCheckCount < 1 {
			println("Another Spindly app is already running. Exiting...")
			return
		}
	}

	os.Chdir(currentWorkingDir)
	cwd = currentWorkingDir

	oswd, _ := os.Getwd()
	println("Spindly Working Directory: " + oswd)

	// List all files in the current working directory
	files, _ := ioutil.ReadDir(cwd)
	for _, f := range files {
		println(f.Name())
	}

	configure()

	runServer()
}
