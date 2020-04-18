package influxdb

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/n7down/iota/internal/sensors"
)

func (i InfluxDB) LogStats(measurement string, sensor *sensors.StatsSensor) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// indexed
	tags := map[string]string{
		"mac": sensor.Mac,
	}

	voltageFloat, err := sensor.GetVoltageFloat()
	if err != nil {
		return err
	}

	connectFloat, err := sensor.GetConnectFloat()
	if err != nil {
		return err
	}

	// not indexed
	fields := map[string]interface{}{
		"voltage": voltageFloat,
		"connect": connectFloat,
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
