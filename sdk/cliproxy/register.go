package cliproxy

import (
	configaccess "github.com/router-for-me/CLIProxyAPI/v6/internal/access/config_access"
)

func init() {
	// Ensure default request-access providers are registered even when embedding.
	configaccess.Register()
}
