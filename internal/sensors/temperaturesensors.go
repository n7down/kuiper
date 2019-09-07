package sensors

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/n7down/iota/internal/stores"
)

type TemperatureSensors struct {
	ID                string `json:"id"`
	DHT22Temperature  string `json:"dht22temp"`
	BMP280Temperature string `json:"bmp280temp"`
}

func (i TemperatureSensors) LogSensors(store *stores.InfluxStore, measurement string) error {
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
		"dht22_temp":  i.DHT22Temperature,
		"bmp280_temp": i.BMP280Temperature,
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
