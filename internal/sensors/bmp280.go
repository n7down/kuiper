package sensors

type BMP280Sensor struct {
	Mac         string  `json:"mac"`
	Pressure    float32 `json:"pres"`
	Temperature float32 `json:"temp"`
}
