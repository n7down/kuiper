package sensors

type DHT22Sensors struct {
	ID               string `json:"id"`
	DHT22Humidity    string `json:"dht22hum"`
	DHT22Temperature string `json:"dht22temp"`
}
