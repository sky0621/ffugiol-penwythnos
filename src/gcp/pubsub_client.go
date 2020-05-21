package gcp

import (
	"context"
	"fmt"

	"golang.org/x/xerrors"

	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/gcppubsub"
)

type Topic int

const (
	// 動画作成トピック
	CreateMovieTopic Topic = iota + 1
)

type PubSubClient interface {
	SendCreateMovieTopic(ctx context.Context, facilityID, bodyJSON string) error
}

type pubSubClient struct {
	topicList map[Topic]string
}

func NewPubSubClient(createMovieTopic string) PubSubClient {
	return &pubSubClient{
		topicList: map[Topic]string{
			CreateMovieTopic: createMovieTopic,
		},
	}
}

func (c *pubSubClient) SendCreateMovieTopic(ctx context.Context, facilityID, bodyJSON string) (e error) {
	topic, err := pubsub.OpenTopic(ctx, fmt.Sprintf("gcppubsub://%s", c.topicList[CreateMovieTopic]))
	if err != nil {
		return xerrors.Errorf("failed to Open Topic: %w", err)
	}
	defer func() {
		if topic != nil {
			if err := topic.Shutdown(ctx); err != nil {
				e = xerrors.Errorf("failed to Shutdown Topic: %w", err)
			}
		}
	}()

	err = topic.Send(ctx, &pubsub.Message{
		Metadata: map[string]string{
			"facility-id": facilityID,
		},
		Body: []byte(bodyJSON),
	})
	if err != nil {
		return xerrors.Errorf("failed to Send Topic: %w", err)
	}
	return nil
}
