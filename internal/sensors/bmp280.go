package sensors

type BMP280Sensor struct {
	ID          string `json:"id"`
	Pressure    string `json:"pres"`
	Temperature string `json:"temp"`
}
