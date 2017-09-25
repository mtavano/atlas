package service

import (
	"log"
	"time"

	"github.com/paypal/gatt"
	rpio "github.com/stianeikeland/go-rpio"
)

// NewAuthService ...
func NewAuthService() *gatt.Service {
	s := gatt.NewService(gatt.MustParseUUID("2000000F-3ED1-488D-B6E8-6A27D888E256"))
	s.AddCharacteristic(gatt.MustParseUUID("2000001F-3ED1-488D-B6E8-6A27D888E256")).HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			key := string(data)

			if key == "development-key" {
				log.Println("Opening door ...")

				_ = rpio.Open()
				defer rpio.Close()

				relay := rpio.Pin(16)
				relay.Output()

				relay.High()
				time.Sleep(3 * time.Second)
				relay.Low()

				return gatt.StatusSuccess
			}

			log.Println("Incorrect key: ", key)
			return gatt.StatusUnexpectedError
		})
	return s
}
