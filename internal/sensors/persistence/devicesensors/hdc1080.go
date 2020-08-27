package sensors

import "time"

type HDC1080Measurement struct {
	Mac         string  `json:"m"`
	Humidity    float64 `json:"h"`
	Temperature float64 `json:"t"`
}

type HDC1080TemperatureMeasurement struct {
	timestamp   time.Time
	Temperature float64
}

type HDC1080TemperatureMeasurements struct {
	Mac                     string
	TemperatureMeasurements []HDC1080TemperatureMeasurement
}

type HDC1080HumidityMeasurement struct {
	timestamp time.Time
	Humidity  float64
}

type HDC1080HumidityMeasurements struct {
	Mac                  string
	HumidityMeasurements []HDC1080HumidityMeasurement
}
