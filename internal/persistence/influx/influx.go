package influx

import (
	"fmt"
	"net/url"

	client "github.com/influxdata/influxdb1-client/v2"
)

type Influx struct {
	Client   client.Client
	Database string
}

func NewInflux(url *url.URL) (*Influx, error) {
	i := &Influx{}
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

	i = &Influx{
		Client:   influxClient,
		Database: database,
	}

	return i, nil
}
