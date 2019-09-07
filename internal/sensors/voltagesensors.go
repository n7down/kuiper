package sensors

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/n7down/iota/internal/stores"
)

type VoltageSensors struct {
	ID      string `json:"id"`
	Voltage string `json:"voltage"`
}

func (i VoltageSensors) LogSensors(store *stores.InfluxStore, measurement string) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  store.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// indexed
	tags := map[string]string{
		"id": i.ID,
	}

	// not indexed
	fields := map[string]interface{}{
		"voltage": i.Voltage,
	}

	point, err := client.NewPoint(
		measurement,
		tags,
		fields,
		time.Now(),
	)

	bp.AddPoint(point)

	err = store.Client.Write(bp)
	if err != nil {
		return err
	}

	return nil
}
