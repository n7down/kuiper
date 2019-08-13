package sensors

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/n7down/IOTWeatherStation/internal/stores"
)

type SensorDataPoint struct {
	Location                 string `json:"location"`
	Label                    string `json:"label"`   // label is a the label on the device
	Version                  string `json:"version"` // version is the firmware version on the device
	Device                   string `json:"device"`  // device is determined by the hardware that makes up the device
	Voltage                  string `json:"voltage"`
	DHT22Temperature         string `json:"dht22temp"`
	DHT22Humidity            string `json:"dht22humidity"`
	BMP280Temperature        string `json:"bmp280temperture"`
	BMP280BarometricPressure string `json:"bmp280pressure"`
}

func (s SensorDataPoint) LogSensor(store *stores.InfluxStore) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  store.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// indexed
	tags := map[string]string{
		"location": s.Location,
		"label":    s.Label,
		"version":  s.Version,
		"device":   s.Device,
	}

	// not indexed
	fields := map[string]interface{}{
		"voltage":              s.Voltage,
		"dht22_temp":           s.DHT22Temperature,
		"dht22_humidity":       s.DHT22Humidity,
		"bmp280_temp":          s.BMP280Temperature,
		"bmp280_bero_pressure": s.BMP280BarometricPressure,
	}

	point, err := client.NewPoint(
		"sensor",
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
