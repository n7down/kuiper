package listeners

import (
	"github.com/n7down/iota/internal/sensors/persistence/influxdb"
)

type Env struct {
	influxDB *influxdb.InfluxDB
}

func NewEnv(influxDB *influxdb.InfluxDB) *Env {
	return &Env{
		influxDB: influxDB,
	}
}
