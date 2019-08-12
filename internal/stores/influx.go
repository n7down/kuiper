package stores

import (
	"net/url"

	client "github.com/influxdata/influxdb1-client/v2"
)

type InfluxStore struct {
	Client   client.Client
	Database string
}

func NewInfluxStore(url *url.URL) (*InfluxStore, error) {
	i := &InfluxStore{}
	username := url.User.Username()
	password, _ := url.User.Password()

	influxClient, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: url.Host,
		//Username: "dbuser",
		//Password: "password",
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

	i = &InfluxStore{
		Client:   influxClient,
		Database: database,
	}

	return i, nil
}
