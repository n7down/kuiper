package sensors

type DHT22Sensor struct {
	Mac         string  `json:"mac"`
	Humidity    float64 `json:"humidity"`
	Temperature float64 `json:"temp"`
}
