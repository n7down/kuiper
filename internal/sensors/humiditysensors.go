package sensors

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/n7down/iota/internal/stores"
)

type HumiditySensors struct {
	ID string `json:"id"`
	//Voltage           string `json:"volt"`
	//DHT22Temperature  string `json:"dht22temp"`
	DHT22Humidity string `json:"dht22hum"`
	//BMP280Temperature string `json:"bmp280temp"`
}

func (i HumiditySensors) LogSensors(store *stores.InfluxStore, measurement string) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  store.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// indexed
	tags := map[string]string{
		"id": i.ID,
	}

	// not indexed
	fields := map[string]interface{}{
		//"voltage":        i.Voltage,
		//"dht22_temp":     i.DHT22Temperature,
		"dht22_humidity": i.DHT22Humidity,
		//"bmp280_temp":    i.BMP280Temperature,
	}

	point, err := client.NewPoint(
		measurement,
		tags,
		fields,
		time.Now(),
	)

	bp.AddPoint(point)

	err = store.Client.Write(bp)
	if err != nil {
		return err
	}

	return nil
}
