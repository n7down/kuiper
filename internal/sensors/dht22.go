package sensors

type DHT22Sensor struct {
	ID          string `json:"id"`
	Humidity    string `json:"humidity"`
	Temperature string `json:"temp"`
}
