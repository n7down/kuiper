package googlepubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/n7down/kuiper/internal/logger"
	"github.com/n7down/kuiper/internal/sensors/persistence"
)

type GooglePubSub struct {
	client      *pubsub.Client
	persistence persistence.Persistence
	logger      logger.Logger
}

func NewGooglePubSub(ctx context.Context, projectID string, persistence persistence.Persistence, logger logger.Logger) (*GooglePubSub, error) {

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return &GooglePubSub{}, err
	}

	return &GooglePubSub{
		client:      client,
		persistence: persistence,
		logger:      logger,
	}, nil
}
