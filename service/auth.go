package service

import (
	"log"
	"time"

	"github.com/mtavano/atlas"
	"github.com/paypal/gatt"
	rpio "github.com/stianeikeland/go-rpio"
)

// NewAuthService ...
func NewAuthService(kst atlas.KeyStore, s *gatt.Service) *gatt.Service {
	s.AddCharacteristic(gatt.MustParseUUID("2000001F-3ED1-488D-B6E8-6A27D888E256")).HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			token := string(data)

			k, err := kst.FindKeyByToken(token)
			if err != nil {
				log.Println("Key: ", token)
				log.Println("Incorrect key: ", data)
				return gatt.StatusUnexpectedError
			}

			if k.ExpiredAt.UnixNano() < time.Now().UnixNano() {
				return gatt.StatusUnexpectedError
			}

			log.Println("key struct: ", &k)

			_ = rpio.Open()
			defer rpio.Close()

			relay := rpio.Pin(16)
			relay.Output()
			relay.High()
			time.Sleep(500 * time.Millisecond)
			relay.Low()

			return gatt.StatusSuccess
		})
	return s
}

// NewStatusKeyService ...
func NewStatusKeyService(s *gatt.Service) *gatt.Service {

	s.AddCharacteristic(gatt.MustParseUUID("2000002F-3ED1-488D-B6E8-6A27D888E256")).HandleReadFunc(
		func(rsw gatt.ResponseWriter, rrq *gatt.ReadRequest) {
			rsw.SetStatus(2)
			rsw.Write([]byte("2"))
		})

	return s
}
