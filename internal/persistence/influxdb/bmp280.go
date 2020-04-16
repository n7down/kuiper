package influxdb

import (
	"strconv"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/n7down/iota/internal/sensors"
)

func (i InfluxDB) LogBMP280(measurement string, sensor *sensors.BMP280Sensor) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	pressureFloat, err := strconv.ParseFloat(sensor.Pressure, 64)
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
		"bmp280_pressure": pressureFloat,
		"bmp280_temp":     temperatureFloat,
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
