package sensors

type BMP280Sensors struct {
	ID                string `json:"id"`
	BMP280Pressure    string `json:"bmp280pres"`
	BMP280Temperature string `json:"bmp280temp"`
}
