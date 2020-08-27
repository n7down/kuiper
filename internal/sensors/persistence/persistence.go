package persistence

import (
	"time"

	sensors "github.com/n7down/kuiper/internal/sensors/persistence/devicesensors"
)

type BMP280Sensor interface {
	CreateBMP280Measurement(sensor *sensors.BMP280Measurement) error
}

type DHT22Sensor interface {
	CreateDHT22Measurement(sensor *sensors.DHT22Measurement) error
}

type HDC1080Sensor interface {
	CreateHDC1080Measurement(sensor *sensors.HDC1080Measurement) error
	GetHDC1080TemperatureMeasurements(mac string, startTime, endTime time.Time) (sensors.HDC1080TemperatureMeasurements, error)
	GetHDC1080HumidityMeasurements(mac string, startTime, endTime time.Time) (sensors.HDC1080HumidityMeasurements, error)
}

type Stats interface {
	CreateStatsMeasurement(sensor *sensors.StatsMeasurement) error
}

type Voltage interface {
	CreateVoltageMeasurement(sensor *sensors.VoltageMeasurement) error
}

type Persistence interface {
	BMP280Sensor
	DHT22Sensor
	HDC1080Sensor
	Stats
	Voltage
}
