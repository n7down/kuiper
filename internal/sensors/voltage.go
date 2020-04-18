package sensors

import "strconv"

type VoltageSensor struct {
	Mac     string `json:"mac"`
	Voltage string `json:"voltage"`
}

func (s VoltageSensor) GetVoltageFloat() (float64, error) {
	voltageFloat, err := strconv.ParseFloat(s.Voltage, 64)
	if err != nil {
		return 0, err
	}
	return voltageFloat, nil
}
