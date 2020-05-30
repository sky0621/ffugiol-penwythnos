package gcp

import (
	"context"
	"fmt"
	"log"

	"github.com/sky0621/fs-mng-backend/src/util"

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
	var targetTopic string
	{
		if util.IsLocal(c.env) {
			targetTopic = fmt.Sprintf("mem://%s", c.topicList[topic])
		} else {
			targetTopic = fmt.Sprintf("gcppubsub://projects/%s/topics/%s", c.projectID, c.topicList[topic])
		}
	}
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

	err = t.Send(ctx, &pubsub.Message{
		Metadata: metadata,
		Body:     []byte(bodyJSON),
	})
	if err != nil {
		return xerrors.Errorf("failed to Send Topic: %w", err)
	}

	if util.IsLocal(c.env) {
		sub, err := pubsub.OpenSubscription(ctx, targetTopic)
		if err != nil {
			log.Printf("%+v\n", err)
			return err
		}
		defer func() {
			if sub != nil {
				if err := sub.Shutdown(ctx); err != nil {
					log.Printf("%+v\n", err)
					e = err
				}
			}
		}()
		msg, err := sub.Receive(ctx)
		if err != nil {
			log.Printf("%+v\n", err)
			return err
		}
		if msg != nil {
			log.Printf("metadata: %#v", msg.Metadata)
			log.Printf("body: %s", string(msg.Body))

			msg.Ack()
		}
	}

	return nil
}
