package sensors

type BMP280Sensor struct {
	Mac         string  `json:"mac"`
	Pressure    float64 `json:"pres"`
	Temperature float64 `json:"temp"`
}
