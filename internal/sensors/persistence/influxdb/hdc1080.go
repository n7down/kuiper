package influxdb

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	sensors "github.com/n7down/kuiper/internal/sensors/devicesensors"
)

func (i InfluxDB) LogHDC1080(sensor *sensors.HDC1080Sensor) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// humidityFloat, err := sensor.GetHumidityFloat()
	// if err != nil {
	// 	return err
	// }

	// temperatureFloat, err := sensor.GetTemperatureFloat()
	// if err != nil {
	// 	return err
	// }

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
		"hdc1080_listener",
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
