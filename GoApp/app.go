package GoApp

import (
	"strconv"
	"strings"
	"time"

	"github.com/spindlygo/spindly-demo/spindlyapp"
)

func Main() {

	println(" --- Spindly Server --- ")

	spindlyapp.DefaultPort = "43216"

	// Set OnInstanciate methods before anything is instantiated
	spindlyapp.ClockHub_OnInstanciate = StartClock
	spindlyapp.ExampleHub_OnInstanciate = ExampleHub_OnInstanciate

	// Configure Hubs and Server Routes
	spindlyapp.Configure()

	// You can access Hubs now
	println(spindlyapp.Global.GetAppName())
	spindlyapp.Global.HelloMessage.Set("Hello there, this messege is set from Go code!")

	// Handle an event. Here we are responding to the "SaidHello" event
	spindlyapp.Global.SaidHello.OnChange(SaidHello)

	// Start serving the UI. This will block until the server is stopped.
	spindlyapp.Serve()

}

func SaidHello(interface{}) {
	// This function is called everytime when the value of the SaidHello event occurs
	spindlyapp.Global.HelloMessage.Set("Hello there, it's " + time.Now().Format("15:04:05") + " now. You are right on time!")

	events := spindlyapp.Global.GetEvents()
	events = append(events, "Said hello at "+time.Now().Format("15:04:05"))
	if len(events) > 5 {
		events = events[1:]
	}

	spindlyapp.Global.Events.Set(events)
}

func ExampleHub_OnInstanciate(hub *spindlyapp.ExampleHub) {

	// Set a default value
	hub.Greating.Set("Hello from Go!")

	hub.Name.OnChange(
		// This function is called everytime when the value of the Name property changes
		func(newValue interface{}) {

			var name = newValue.(string)
			println("Message changed to: " + name)

			// Just for fun, we are going to change the greeting message
			if len(name) > 0 {
				num := (len(name)*11)%20 + 20
				hub.Greating.Set("Nice name '" + strings.Title(name) + "', your magic number is " + strconv.Itoa(num))
			} else {
				hub.Greating.Set("Can you tell me your name?")
			}

		},
	)
}

func StartClock(clock *spindlyapp.ClockHub) {
	// Create a timer to run every second
	timer := time.NewTicker(time.Second)

	// This loop executes every second
	for range timer.C {
		// Get the current time in clocks time zone
		displayTime := TimeIn(time.Now().UTC(), clock.GetTimeZone())

		// Set the spindly store value, this will be immediately synched with the correspoinding svelte store
		clock.ClockFace.Set(displayTime.Format("15:04:05"))
	}
}

// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) time.Time {
	loc, err := time.LoadLocation(name)
	if err == nil {
		return t.In(loc)
	} else {
		return t
	}
}
