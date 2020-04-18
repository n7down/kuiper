package sensors

import "strconv"

type StatsSensor struct {
	Mac     string `json:"mac"`
	Voltage string `json:"voltage"`
	Connect string `json:"connect"`
}

func (s StatsSensor) GetVoltageFloat() (float64, error) {
	voltageFloat, err := strconv.ParseFloat(s.Voltage, 64)
	if err != nil {
		return 0, err
	}
	return voltageFloat, nil
}

func (s StatsSensor) GetConnectFloat() (float64, error) {
	connectFloat, err := strconv.ParseFloat(s.Connect, 64)
	if err != nil {
		return 0, err
	}
	return connectFloat, nil
}
