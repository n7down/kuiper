package influx

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

type VoltageData struct {
	ID      string `json:"id"`
	Voltage string `json:"voltage"`
}

func (i Influx) LogVoltage(measurement string, voltageData VoltageData) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// indexed
	tags := map[string]string{
		"id": voltageData.ID,
	}

	// not indexed
	fields := map[string]interface{}{
		"voltage": voltageData.Voltage,
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
