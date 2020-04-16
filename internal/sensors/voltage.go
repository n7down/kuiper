package sensors

type VoltageSensor struct {
	Mac     string  `json:"mac"`
	Voltage float32 `json:"voltage"`
}
