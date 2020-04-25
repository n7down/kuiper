package influxdb

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	sensors "github.com/n7down/kuiper/internal/sensors/devicesensors"
)

func (i InfluxDB) LogBMP280(measurement string, sensor *sensors.BMP280Sensor) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	pressureFloat, err := sensor.GetPressureFloat()
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
