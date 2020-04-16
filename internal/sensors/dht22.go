package sensors

type DHT22Sensor struct {
	Mac         string  `json:"mac"`
	Humidity    float32 `json:"humidity"`
	Temperature float32 `json:"temp"`
}
