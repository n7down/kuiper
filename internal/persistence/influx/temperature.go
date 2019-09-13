package influx

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

type TemperatureData struct {
	ID                string `json:"id"`
	DHT22Temperature  string `json:"dht22temp"`
	BMP280Temperature string `json:"bmp280temp"`
}

func (i Influx) LogTemperature(measurement string, temperatureData TemperatureData) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// indexed
	tags := map[string]string{
		"id": temperatureData.ID,
	}

	// not indexed
	fields := map[string]interface{}{
		"dht22_temp":  temperatureData.DHT22Temperature,
		"bmp280_temp": temperatureData.BMP280Temperature,
	}

	point, err := client.NewPoint(
		measurement,
		tags,
		fields,
		time.Now(),
	)

	bp.AddPoint(point)

	err = i.Client.Write(bp)
	if err != nil {
		return err
	}

	return nil
}
