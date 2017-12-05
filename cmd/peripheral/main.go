package main

import (
	"fmt"
	"log"

	"github.com/mtavano/atlas/postgres"
	"github.com/mtavano/atlas/service"
	"github.com/paypal/gatt"
)

// TODO: review why it compiles with gatt/examples/option and noth with atlas/option

func main() {
	host := "gate-dev.cfjgzmj532wd.us-east-1.rds.amazonaws.com"
	port := 5432
	user := "gate"
	password := "Gate12345"
	dbname := "gate_api"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := postgres.NewDatastore(psqlInfo)
	keyStore := &postgres.KeyStore{Datastore: db}

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

			auth := service.NewAuthService(keyStore)
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
