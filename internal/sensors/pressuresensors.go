package sensors

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/n7down/iota/internal/stores"
)

type PressureSensors struct {
	ID             string `json:"id"`
	BMP280Pressure string `json:"bmp280pres"`
}

func (i PressureSensors) LogSensors(store *stores.InfluxStore, measurement string) error {
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
		"bmp280_pressure": i.BMP280Pressure,
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
