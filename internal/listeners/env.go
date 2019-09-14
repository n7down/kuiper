package listeners

import (
	"github.com/n7down/iota/internal/persistence/influx"
)

type Env struct {
	influxDB *influx.Influx
}

func NewEnv(influxDB *influx.Influx) *Env {
	return &Env{
		influxDB: influxDB,
	}
}
