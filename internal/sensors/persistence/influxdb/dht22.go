package influxdb

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	sensors "github.com/n7down/iota/internal/sensors/devicesensors"
)

func (i InfluxDB) LogDHT22(measurement string, sensor *sensors.DHT22Sensor) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	humidityFloat, err := sensor.GetHumidityFloat()
	if err != nil {
		return err
	}

	temperatureFloat, err := sensor.GetTemperatureFloat()
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
