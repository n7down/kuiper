package sensors

type VoltageSensor struct {
	Mac     string  `json:"mac"`
	Voltage float64 `json:"voltage"`
}
