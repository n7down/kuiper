package googlepubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/n7down/kuiper/internal/logger"
	"github.com/n7down/kuiper/internal/settings/persistence"
)

type Client interface {
	Subscription(id string) *pubsub.Subscription
	SubscriptionInProject(id, projectID string) *pubsub.Subscription
	Subscriptions(ctx context.Context) *pubsub.SubscriptionIterator
	CreateSubscription(ctx context.Context, id string, cfg pubsub.SubscriptionConfig) (*pubsub.Subscription, error)
	CreateTopic(ctx context.Context, topicID string) (*pubsub.Topic, error)
	CreateTopicWithConfig(ctx context.Context, topicID string, tc *pubsub.TopicConfig) (*pubsub.Topic, error)
	Topic(id string) *pubsub.Topic
	TopicInProject(id, projectID string) *pubsub.Topic
	DetachSubscription(ctx context.Context, sub string) (*pubsub.DetachSubscriptionResult, error)
	Topics(ctx context.Context) *pubsub.TopicIterator
}

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

func NewGooglePubSubWithClient(ctx context.Context, projectID string, persistence persistence.Persistence, logger logger.Logger, client *pubsub.Client) (*GooglePubSub, error) {
	return &GooglePubSub{
		client:      client,
		persistence: persistence,
		logger:      logger,
	}, nil
}
