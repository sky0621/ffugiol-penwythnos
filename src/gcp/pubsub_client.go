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
	SendTopic(ctx context.Context, topic Topic, metadata map[string]string, bodyJSON string) error
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

func (c *pubSubClient) SendTopic(ctx context.Context, topic Topic, metadata map[string]string, bodyJSON string) (e error) {
	t, err := pubsub.OpenTopic(ctx, fmt.Sprintf("gcppubsub://%s", c.topicList[topic]))
	if err != nil {
		return xerrors.Errorf("failed to Open Topic: %w", err)
	}
	defer func() {
		if t != nil {
			if err := t.Shutdown(ctx); err != nil {
				e = xerrors.Errorf("failed to Shutdown Topic: %w", err)
			}
		}
	}()

	err = t.Send(ctx, &pubsub.Message{
		Metadata: metadata,
		Body:     []byte(bodyJSON),
	})
	if err != nil {
		return xerrors.Errorf("failed to Send Topic: %w", err)
	}
	return nil
}
