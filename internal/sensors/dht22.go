package sensors

type DHT22Sensor struct {
	Mac         string `json:"mac"`
	Humidity    string `json:"humidity"`
	Temperature string `json:"temp"`
}
