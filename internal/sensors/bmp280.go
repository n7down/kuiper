package sensors

type BMP280Sensor struct {
	Mac         string `json:"mac"`
	Pressure    string `json:"pres"`
	Temperature string `json:"temp"`
}
