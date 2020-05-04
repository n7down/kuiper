package sensors

type HDC1080Sensor struct {
	Mac         string  `json:"m"`
	Humidity    float64 `json:"h"`
	Temperature float64 `json:"t"`
}

// func (s HDC1080Sensor) GetHumidityFloat() (float64, error) {
// 	humidityFloat, err := strconv.ParseFloat(s.Humidity, 64)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return humidityFloat, nil
// }

// func (s HDC1080Sensor) GetTemperatureFloat() (float64, error) {
// 	temperatureFloat, err := strconv.ParseFloat(s.Temperature, 64)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return temperatureFloat, nil
// }
