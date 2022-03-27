package spindlyapp

import (
	"time"

	"github.com/spindlygo/SpindlyExports"
	"github.com/spindlygo/spindly/Spindly"
)

var HubManager *Spindly.HubManager = Spindly.NewHubManager()

type GlobalHub struct {
	Instance     *Spindly.HubInstance
	AppName      Spindly.SpindlyStore
	SaidHello    Spindly.SpindlyStore
	HelloMessage Spindly.SpindlyStore
	Events       Spindly.SpindlyStore
}

var GlobalHub_OnInstanciate func(*GlobalHub)

type GlobalHubExported struct {
	AppName      *SpindlyExports.ExportedStore
	SaidHello    *SpindlyExports.ExportedStore
	HelloMessage *SpindlyExports.ExportedStore
	Events       *SpindlyExports.ExportedStore
}

func (hub *GlobalHub) ToExported() *GlobalHubExported {
	return &GlobalHubExported{
		AppName:      hub.AppName.ToExported(),
		SaidHello:    hub.SaidHello.ToExported(),
		HelloMessage: hub.HelloMessage.ToExported(),
		Events:       hub.Events.ToExported(),
	}
}

var Global *GlobalHub

func (hub GlobalHub) New(InstanceID string) *GlobalHub {
	hub.Instanciate(InstanceID)
	return &hub
}

func (hub *GlobalHub) Instanciate(InstanceID string) *Spindly.HubInstance {
	hub.Instance = &Spindly.HubInstance{
		HubClass:   "GlobalHub",
		InstanceID: InstanceID,
		Stores:     make(map[string]*Spindly.SpindlyStore),
	}

	hub.AppName = Spindly.NewSpindlyStore(
		"AppName",
		func() interface{} {
			return ``
		},
		`Spindly`,
	)
	hub.Instance.Register(&hub.AppName)

	hub.SaidHello = Spindly.NewSpindlyStore(
		"SaidHello",
		func() interface{} {
			return 0.0
		},
		nil,
	)
	hub.Instance.Register(&hub.SaidHello)

	hub.HelloMessage = Spindly.NewSpindlyStore(
		"HelloMessage",
		func() interface{} {
			return ""
		},
		nil,
	)
	hub.Instance.Register(&hub.HelloMessage)

	hub.Events = Spindly.NewSpindlyStore(
		"Events",
		func() interface{} {
			return []interface{}{}
		},
		[]interface{}{},
	)
	hub.Instance.Register(&hub.Events)

	HubManager.Register(hub.Instance)
	if GlobalHub_OnInstanciate != nil {
		go GlobalHub_OnInstanciate(hub)
	}
	return hub.Instance
}

func (hub *GlobalHub) GetAppName() string {
	return hub.AppName.Get().(string)
}
func (hub *GlobalHub) GetSaidHello() float64 {
	return hub.SaidHello.Get().(float64)
}
func (hub *GlobalHub) GetHelloMessage() string {
	return hub.HelloMessage.Get().(string)
}
func (hub *GlobalHub) GetEvents() []interface{} {
	return hub.Events.Get().([]interface{})
}

type ClockHub struct {
	Instance  *Spindly.HubInstance
	ClockFace Spindly.SpindlyStore
	TimeZone  Spindly.SpindlyStore
}

var ClockHub_OnInstanciate func(*ClockHub)

type ClockHubExported struct {
	ClockFace *SpindlyExports.ExportedStore
	TimeZone  *SpindlyExports.ExportedStore
}

func (hub *ClockHub) ToExported() *ClockHubExported {
	return &ClockHubExported{
		ClockFace: hub.ClockFace.ToExported(),
		TimeZone:  hub.TimeZone.ToExported(),
	}
}

func (hub ClockHub) New(InstanceID string) *ClockHub {
	hub.Instanciate(InstanceID)
	return &hub
}

func (hub *ClockHub) Instanciate(InstanceID string) *Spindly.HubInstance {
	hub.Instance = &Spindly.HubInstance{
		HubClass:   "ClockHub",
		InstanceID: InstanceID,
		Stores:     make(map[string]*Spindly.SpindlyStore),
	}

	hub.ClockFace = Spindly.NewSpindlyStore(
		"ClockFace",
		func() interface{} {
			return ""
		},
		nil,
	)
	hub.Instance.Register(&hub.ClockFace)

	hub.TimeZone = Spindly.NewSpindlyStore(
		"TimeZone",
		func() interface{} {
			return ""
		},
		`Local`,
	)
	hub.Instance.Register(&hub.TimeZone)

	HubManager.Register(hub.Instance)
	if ClockHub_OnInstanciate != nil {
		go ClockHub_OnInstanciate(hub)
	}
	return hub.Instance
}

func (hub *ClockHub) GetClockFace() string {
	return hub.ClockFace.Get().(string)
}
func (hub *ClockHub) GetTimeZone() string {
	return hub.TimeZone.Get().(string)
}

type ExampleHub struct {
	Instance *Spindly.HubInstance
	Name     Spindly.SpindlyStore
	Greating Spindly.SpindlyStore
}

var ExampleHub_OnInstanciate func(*ExampleHub)

type ExampleHubExported struct {
	Name     *SpindlyExports.ExportedStore
	Greating *SpindlyExports.ExportedStore
}

func (hub *ExampleHub) ToExported() *ExampleHubExported {
	return &ExampleHubExported{
		Name:     hub.Name.ToExported(),
		Greating: hub.Greating.ToExported(),
	}
}

func (hub ExampleHub) New(InstanceID string) *ExampleHub {
	hub.Instanciate(InstanceID)
	return &hub
}

func (hub *ExampleHub) Instanciate(InstanceID string) *Spindly.HubInstance {
	hub.Instance = &Spindly.HubInstance{
		HubClass:   "ExampleHub",
		InstanceID: InstanceID,
		Stores:     make(map[string]*Spindly.SpindlyStore),
	}

	hub.Name = Spindly.NewSpindlyStore(
		"Name",
		func() interface{} {
			return ""
		},
		nil,
	)
	hub.Instance.Register(&hub.Name)

	hub.Greating = Spindly.NewSpindlyStore(
		"Greating",
		func() interface{} {
			return ""
		},
		nil,
	)
	hub.Instance.Register(&hub.Greating)

	HubManager.Register(hub.Instance)
	if ExampleHub_OnInstanciate != nil {
		go ExampleHub_OnInstanciate(hub)
	}
	return hub.Instance
}

func (hub *ExampleHub) GetName() string {
	return hub.Name.Get().(string)
}
func (hub *ExampleHub) GetGreating() string {
	return hub.Greating.Get().(string)
}

func InitializeHubs() {
	HubManager.RegisterClass("GlobalHub", func() Spindly.HubClass { return &GlobalHub{} })
	Global = GlobalHub{}.New("Global")
	HubManager.RegisterClass("ClockHub", func() Spindly.HubClass { return &ClockHub{} })
	HubManager.RegisterClass("ExampleHub", func() Spindly.HubClass { return &ExampleHub{} })
	namedExports = &NamedExportedHubs{
		Global: Global.ToExported(),
	}
}

type NamedExportedHubs struct {
	Global *GlobalHubExported
}

var namedExports *NamedExportedHubs = nil

func NamedExports() *NamedExportedHubs {
	for namedExports == nil {
		println("Waiting for NamedExports")
		time.Sleep(100 * time.Millisecond)
	}

	return namedExports
}
