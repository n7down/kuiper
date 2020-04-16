package influxdb

import (
	"strconv"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/n7down/iota/internal/sensors"
)

func (i InfluxDB) LogDHT22(measurement string, sensor *sensors.DHT22Sensor) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	humidityFloat, err := strconv.ParseFloat(sensor.Humidity, 64)
	if err != nil {
		return err
	}

	temperatureFloat, err := strconv.ParseFloat(sensor.Temperature, 64)
	if err != nil {
		return err
	}

	// indexed
	tags := map[string]string{
		"mac": sensor.Mac,
	}

	// not indexed
	fields := map[string]interface{}{
		"dht22_humidity": humidityFloat,
		"dht22_temp":     temperatureFloat,
	}

	point, err := client.NewPoint(
		measurement,
		tags,
		fields,
		time.Now().UTC(),
	)

	bp.AddPoint(point)

	err = i.Client.Write(bp)
	if err != nil {
		return err
	}

	return nil
}
