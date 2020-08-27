package influxpersistence

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	sensors "github.com/n7down/kuiper/internal/sensors/persistence/devicesensors"
)

func (i InfluxPersistence) CreateDHT22Measurement(sensor *sensors.DHT22Measurement) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// indexed
	tags := map[string]string{
		"mac": sensor.Mac,
	}

	// not indexed
	fields := map[string]interface{}{
		"humidity": sensor.Humidity,
		"temp":     sensor.Temperature,
	}

	point, err := client.NewPoint(
		"dht22_listener",
		tags,
		fields,
		time.Now().UTC(),
	)

	bp.AddPoint(point)

	err = i.client.Write(bp)
	if err != nil {
		return err
	}

	return nil
}
