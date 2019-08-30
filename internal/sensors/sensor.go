package sensors

import (
	"github.com/n7down/iota/internal/stores"
)

type Sensors interface {
	LogSensors(store *stores.InfluxStore, measurement string) error
}
