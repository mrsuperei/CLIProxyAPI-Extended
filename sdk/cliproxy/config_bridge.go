package cliproxy

import (
	"fmt"

	internalconfig "github.com/router-for-me/CLIProxyAPI/v6/internal/config"
)

// Config exposes the full server configuration structure used by the embedded service.
// It aliases the internal Config type so embedding applications can work with the same
// fields without importing an internal package.
type Config = internalconfig.Config

// LoadConfig loads the CLI Proxy configuration from the supplied file path.
// It mirrors the CLI binary's loader so embedded hosts can share the same YAML format.
func LoadConfig(path string) (*Config, error) {
	if path == "" {
		return nil, fmt.Errorf("cliproxy: configuration path is required")
	}
	return internalconfig.LoadConfig(path)
}
