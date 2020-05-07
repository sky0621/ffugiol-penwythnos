package graph

import (
	"database/sql"

	"github.com/sky0621/fs-mng-backend/src/auth"

	"github.com/sky0621/fs-mng-backend/src/gcp"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB          *sql.DB
	GCSClient   gcp.CloudStorageClient
	Auth0Client auth.Auth0Client
}
