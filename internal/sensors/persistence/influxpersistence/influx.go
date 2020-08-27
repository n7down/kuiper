package influxpersistence

import (
	"fmt"
	"net/url"

	"github.com/n7down/kuiper/internal/logger"

	client "github.com/influxdata/influxdb1-client/v2"
)

type InfluxPersistence struct {
	client   client.Client
	database string
	logger   logger.Logger
}

func NewInfluxPersistence(url *url.URL, logger logger.Logger) (*InfluxPersistence, error) {
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
		client:   influxClient,
		database: database,
		logger:   logger,
	}

	return i, nil
}
