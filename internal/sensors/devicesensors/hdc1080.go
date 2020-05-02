package sensors

import "strconv"

type HDC1080Sensor struct {
	Mac         string `json:"m"`
	Humidity    string `json:"h"`
	Temperature string `json:"t"`
}

func (s HDC1080Sensor) GetHumidityFloat() (float64, error) {
	humidityFloat, err := strconv.ParseFloat(s.Humidity, 64)
	if err != nil {
		return 0, err
	}
	return humidityFloat, nil
}

func (s HDC1080Sensor) GetTemperatureFloat() (float64, error) {
	temperatureFloat, err := strconv.ParseFloat(s.Temperature, 64)
	if err != nil {
		return 0, err
	}
	return temperatureFloat, nil
}
