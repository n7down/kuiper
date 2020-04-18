package sensors

import "strconv"

type BMP280Sensor struct {
	Mac         string `json:"m"`
	Pressure    string `json:"p"`
	Temperature string `json:"t"`
}

func (s BMP280Sensor) GetPressureFloat() (float64, error) {
	pressureFloat, err := strconv.ParseFloat(s.Pressure, 64)
	if err != nil {
		return 0, err
	}
	return pressureFloat, nil
}

func (s BMP280Sensor) GetTemperatureFloat() (float64, error) {
	temperatureFloat, err := strconv.ParseFloat(s.Temperature, 64)
	if err != nil {
		return 0, err
	}
	return temperatureFloat, nil
}
