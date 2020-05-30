package gcp

import (
	"context"

	"gocloud.dev/gcp"
	"golang.org/x/xerrors"
)

func GetProjectID() (string, error) {
	credentials, err := gcp.DefaultCredentials(context.Background())
	if err != nil {
		return "", xerrors.Errorf("failed to gcp.DefaultCredentials: %w", err)
	}
	projectID, err := gcp.DefaultProjectID(credentials)
	if err != nil {
		return "", xerrors.Errorf("failed to gcp.DefaultProjectID: %w", err)
	}
	return string(projectID), nil
}
