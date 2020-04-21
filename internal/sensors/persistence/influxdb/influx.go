package influxdb

import (
	"fmt"
	"net/url"

	client "github.com/influxdata/influxdb1-client/v2"
)

type InfluxDB struct {
	Client   client.Client
	Database string
}

func NewInfluxDB(url *url.URL) (*InfluxDB, error) {
	i := &InfluxDB{}
	username := url.User.Username()
	password, _ := url.User.Password()

	influxClient, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     fmt.Sprintf("%s://%s", url.Scheme, url.Host),
		Username: username,
		Password: password,
	})

	if err != nil {
		return i, err
	}

	database := url.Path[1:len(url.Path)]
	if database == "" {
		database = "test"
	}

	i = &InfluxDB{
		Client:   influxClient,
		Database: database,
	}

	return i, nil
}
