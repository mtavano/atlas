package main

import (
	"log"

	"github.com/mtavano/atlas/service"
	"github.com/paypal/gatt"
)

// TODO: review why it compiles with gatt/examples/option and noth with atlas/option

func main() {
	defaultServerOptions := []gatt.Option{}
	dvc, err := gatt.NewDevice(defaultServerOptions...)
	check(err)

	// registed optional handlers, as a kind of middlewares
	dvc.Handle(
		gatt.CentralConnected(onConnection),
		gatt.CentralDisconnected(onDisconnection),
	)

	// A mandatory handler for monitoring device state. as a server listen
	onStateChanged := func(d gatt.Device, s gatt.State) {
		log.Printf("State: %s\n", s)

		switch s {
		case gatt.StatePoweredOn:
			d.AddService(service.NewGapService("Atlas"))

			auth := atlas.service.NewAuthService()
			d.AddService(auth)

			d.AdvertiseNameAndServices("Atlas", []gatt.UUID{auth.UUID()})

		default:
		}
	}

	dvc.Init(onStateChanged)
	select {}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func onConnection(c gatt.Central) { log.Printf("Device %s connected\n", c.ID()) }

func onDisconnection(c gatt.Central) { log.Printf("Device %s disconnected\n", c.ID()) }
