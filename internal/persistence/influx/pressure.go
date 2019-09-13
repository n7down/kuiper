package influx

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

type PressureData struct {
	ID             string `json:"id"`
	BMP280Pressure string `json:"bmp280pres"`
}

func (i Influx) LogPressure(measurement string, pressureData PressureData) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// indexed
	tags := map[string]string{
		"id": pressureData.ID,
	}

	// not indexed
	fields := map[string]interface{}{
		"bmp280_pressure": pressureData.BMP280Pressure,
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
