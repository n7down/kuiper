package listeners

import (
	"github.com/n7down/kuiper/internal/sensors/persistence/influxdb"
)

type SensorsListenersEnv struct {
	influxDB *influxdb.InfluxDB
}

func NewSensorsListenersEnv(influxDB *influxdb.InfluxDB) *SensorsListenersEnv {
	return &SensorsListenersEnv{
		influxDB: influxDB,
	}
}
