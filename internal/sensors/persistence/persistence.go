package persistence

import sensors "github.com/n7down/kuiper/internal/sensors/persistence/devicesensors"

type Persistence interface {
	LogBMP280(sensor *sensors.BMP280Sensor) error
	LogDHT22(sensor *sensors.DHT22Sensor) error
	LogHDC1080(sensor *sensors.HDC1080Sensor) error
	LogStats(sensor *sensors.StatsSensor) error
	LogVoltage(sensor *sensors.VoltageSensor) error
}
