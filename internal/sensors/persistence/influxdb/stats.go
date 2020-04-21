package influxdb

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	sensors "github.com/n7down/iota/internal/sensors/devicesensors"
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

	connectionTimeFloat, err := sensor.GetConnectionTimeFloat()
	if err != nil {
		return err
	}

	deepSleepDelayInt, err := sensor.GetDeepSleepDelayInt()
	if err != nil {
		return err
	}

	// not indexed
	fields := map[string]interface{}{
		"voltage": voltageFloat,
		"connect": connectionTimeFloat,
		"delay":   deepSleepDelayInt,
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
