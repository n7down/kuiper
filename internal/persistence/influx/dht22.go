package influx

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/n7down/iota/internal/sensors"
)

func (i Influx) LogDHT22(measurement string, sensors *sensors.DHT22Sensors) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// indexed
	tags := map[string]string{
		"id": sensors.ID,
	}

	// not indexed
	fields := map[string]interface{}{
		"dht22_humidity": sensors.DHT22Humidity,
		"dht22_temp":     sensors.DHT22Temperature,
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
