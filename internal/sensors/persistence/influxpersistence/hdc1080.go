package influxpersistence

import (
	"fmt"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	sensors "github.com/n7down/kuiper/internal/sensors/persistence/devicesensors"
)

func (i InfluxPersistence) CreateHDC1080Measurement(sensor *sensors.HDC1080Measurement) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

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

	err = i.client.Write(bp)
	if err != nil {
		return err
	}

	return nil
}

func (i InfluxPersistence) GetHDC1080TemperatureMeasurements(mac string, startTime, endTime time.Time) (sensors.HDC1080TemperatureMeasurements, error) {
	query := fmt.Sprintf("select temp from hdc1080_listener where mac = '%s' and time => '%s' and time < '%s'", mac, startTime.String, endTime.String)
	q := client.NewQuery(query, i.database, "s")
	response, err := i.client.Query(q)
	if err != nil {
		return sensors.HDC1080TemperatureMeasurements{}, err
	}
	if response.Error() != nil {
		return sensors.HDC1080TemperatureMeasurements{}, response.Error()
	}

	// fmt.Println(response.Results)
	i.logger.Info(response.Results)

	return sensors.HDC1080TemperatureMeasurements{}, nil
}

func (i InfluxPersistence) GetHDC1080HumidityMeasurements(mac string, startTime, endTime time.Time) (sensors.HDC1080HumidityMeasurements, error) {
	query := fmt.Sprintf("select humidity from hdc1080_listener where mac = '%s' and time => '%s' and time < '%s'", mac, startTime.String, endTime.String)
	q := client.NewQuery(query, i.database, "s")
	response, err := i.client.Query(q)
	if err != nil {
		return sensors.HDC1080HumidityMeasurements{}, err
	}
	if response.Error() != nil {
		return sensors.HDC1080HumidityMeasurements{}, response.Error()
	}

	// fmt.Println(response.Results)
	i.logger.Info(response.Results)

	return sensors.HDC1080HumidityMeasurements{}, nil
}
