package stores

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

func (i InfluxStore) LogHumiditySensors(measurement string, id string, humidity string) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// indexed
	tags := map[string]string{
		"id": id,
	}

	// not indexed
	fields := map[string]interface{}{
		"dht22_humidity": humidity,
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
