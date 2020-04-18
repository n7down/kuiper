package sensors

import "strconv"

type DHT22Sensor struct {
	Mac         string `json:"m"`
	Humidity    string `json:"h"`
	Temperature string `json:"t"`
}

func (s DHT22Sensor) GetHumidityFloat() (float64, error) {
	humidityFloat, err := strconv.ParseFloat(s.Humidity, 64)
	if err != nil {
		return 0, err
	}
	return humidityFloat, nil
}

func (s DHT22Sensor) GetTemperatureFloat() (float64, error) {
	temperatureFloat, err := strconv.ParseFloat(s.Temperature, 64)
	if err != nil {
		return 0, err
	}
	return temperatureFloat, nil
}
