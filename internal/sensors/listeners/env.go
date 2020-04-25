package listeners

import (
	"github.com/n7down/kuiper/internal/sensors/persistence/influxdb"
)

type Env struct {
	influxDB *influxdb.InfluxDB
}

func NewEnv(influxDB *influxdb.InfluxDB) *Env {
	return &Env{
		influxDB: influxDB,
	}
}
