package sensors

import "strconv"

type VoltageMeasurement struct {
	Mac     string `json:"m"`
	Voltage string `json:"v"`
}

func (s VoltageMeasurement) GetVoltageFloat() (float64, error) {
	voltageFloat, err := strconv.ParseFloat(s.Voltage, 64)
	if err != nil {
		return 0, err
	}
	return voltageFloat, nil
}
