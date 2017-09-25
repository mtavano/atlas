package config

import (
	"github.com/paypal/gatt"
)

// DefaultServerOptions ...
var DefaultServerOptions = []gatt.Option{
	gatt.MacDeviceRole(1),
}
