package persistence

import sensors "github.com/n7down/kuiper/internal/sensors/persistence/devicesensors"

type Persistence interface {
	CreateBMP280(sensor *sensors.BMP280Sensor) error
	CreateDHT22(sensor *sensors.DHT22Sensor) error
	CreateHDC1080(sensor *sensors.HDC1080Sensor) error
	CreateStats(sensor *sensors.StatsSensor) error
	CreateVoltage(sensor *sensors.VoltageSensor) error
}
