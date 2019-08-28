package sensors

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/n7down/iota/internal/stores"
)

type AmbientEnvironmentalSensor struct {
	Label             string `json:"label"`
	Voltage           string `json:"volt"`
	DHT22Temperature  string `json:"dht22temp"`
	DHT22Humidity     string `json:"dht22hum"`
	BMP280Temperature string `json:"bmp280temp"`
}

func (a AmbientEnvironmentalSensor) LogSensor(store *stores.InfluxStore) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  store.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// indexed
	tags := map[string]string{
		"label": a.Label,
	}

	// not indexed
	fields := map[string]interface{}{
		"voltage":        a.Voltage,
		"dht22_temp":     a.DHT22Temperature,
		"dht22_humidity": a.DHT22Humidity,
		"bmp280_temp":    a.BMP280Temperature,
	}

	point, err := client.NewPoint(
		"indoor_humidity",
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
