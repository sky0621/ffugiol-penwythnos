package gcp

import (
	"context"
	"fmt"
	"time"

	"gocloud.dev/pubsub/mempubsub"

	"golang.org/x/xerrors"

	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/gcppubsub"
	_ "gocloud.dev/pubsub/mempubsub"
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
	env       string
	projectID string
	topicList map[Topic]string
}

func CreateMemPubSubSet() (*pubsub.Topic, *pubsub.Subscription) {
	tpc := mempubsub.NewTopic()
	sub := mempubsub.NewSubscription(tpc, 1*time.Minute)
	return tpc, sub
}

func NewPubSubClient(env, projectID, createMovieTopic string) PubSubClient {
	return &pubSubClient{
		env:       env,
		projectID: projectID,
		topicList: map[Topic]string{
			CreateMovieTopic: createMovieTopic,
		},
	}
}

func (c *pubSubClient) SendTopic(ctx context.Context, topic Topic, metadata map[string]string, bodyJSON string) (e error) {
	targetTopic := fmt.Sprintf("gcppubsub://projects/%s/topics/%s", c.projectID, c.topicList[topic])
	t, err := pubsub.OpenTopic(ctx, targetTopic)
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

	if err := t.Send(ctx, &pubsub.Message{
		Metadata: metadata,
		Body:     []byte(bodyJSON),
	}); err != nil {
		return xerrors.Errorf("failed to Send Topic: %w", err)
	}

	return nil
}

type pubSubLocalClient struct {
	env       string
	projectID string
	topicList map[Topic]string

	memTopic        *pubsub.Topic
	memSubscription *pubsub.Subscription
}

func NewPubSubLocalClient(env, projectID, createMovieTopic string, memTopic *pubsub.Topic, memSubscription *pubsub.Subscription) PubSubClient {
	return &pubSubLocalClient{
		env:       env,
		projectID: projectID,
		topicList: map[Topic]string{
			CreateMovieTopic: createMovieTopic,
		},
		memTopic:        memTopic,
		memSubscription: memSubscription,
	}
}

func (c *pubSubLocalClient) SendTopic(ctx context.Context, topic Topic, metadata map[string]string, bodyJSON string) (e error) {
	if err := c.memTopic.Send(ctx, &pubsub.Message{
		Metadata: metadata,
		Body:     []byte(bodyJSON),
	}); err != nil {
		return xerrors.Errorf("failed to Send Topic: %w", err)
	}

	msg, err := c.memSubscription.Receive(ctx)
	if err != nil {
		return xerrors.Errorf("failed to Receive: %w", err)
	}
	fmt.Printf("%#v", msg)

	return nil
}
