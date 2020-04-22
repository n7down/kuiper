package sensors

import "strconv"

type StatsSensor struct {
	Mac            string `json:"m"`
	Voltage        string `json:"v"`
	ConnectionTime string `json:"c"`
}

func (s StatsSensor) GetVoltageFloat() (float64, error) {
	voltageFloat, err := strconv.ParseFloat(s.Voltage, 64)
	if err != nil {
		return 0, err
	}
	return voltageFloat, nil
}

func (s StatsSensor) GetConnectionTimeFloat() (float64, error) {
	connectFloat, err := strconv.ParseFloat(s.ConnectionTime, 64)
	if err != nil {
		return 0, err
	}
	return connectFloat, nil
}

func (s StatsSensor) GetConnectionTimeInt() (int, error) {
	connectInt, err := strconv.ParseInt(s.ConnectionTime, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(connectInt), nil
}
