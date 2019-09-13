package influx

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

type HumidityData struct {
	ID            string `json:"id"`
	DHT22Humidity string `json:"dht22hum"`
}

func (i Influx) LogHumidity(measurement string, humidityData HumidityData) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// indexed
	tags := map[string]string{
		"id": humidityData.ID,
	}

	// not indexed
	fields := map[string]interface{}{
		"dht22_humidity": humidityData.DHT22Humidity,
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
