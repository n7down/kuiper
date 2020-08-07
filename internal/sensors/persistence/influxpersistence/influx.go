package influxpersistence

import (
	"fmt"
	"net/url"

	client "github.com/influxdata/influxdb1-client/v2"
)

type InfluxPersistence struct {
	Client   client.Client
	Database string
}

func NewInfluxPersistence(url *url.URL) (*InfluxPersistence, error) {
	i := &InfluxPersistence{}
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

	i = &InfluxPersistence{
		Client:   influxClient,
		Database: database,
	}

	return i, nil
}
