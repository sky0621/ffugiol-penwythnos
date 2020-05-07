package auth

import (
	"context"
	"time"

	"golang.org/x/xerrors"

	"gopkg.in/auth0.v2/management"
)

type NewManagementFunc func(ctx context.Context) (*management.Management, error)

type Auth0Client interface {
	NewManagement(ctx context.Context) (*management.Management, error)
}

type auth0Client struct {
	fn NewManagementFunc
}

func NewAuth0Client(domain, clientID, clientSecret string, timeout time.Duration, debug bool) Auth0Client {
	return &auth0Client{
		fn: func(ctx context.Context) (*management.Management, error) {
			m, err := management.New(domain, clientID, clientSecret, management.WithTimeout(timeout), management.WithDebug(debug))
			if err != nil {
				return nil, xerrors.Errorf("failed to management.New: %w", err)
			}
			return m, nil
		},
	}
}

func (c *auth0Client) NewManagement(ctx context.Context) (*management.Management, error) {
	return c.fn(ctx)
}
