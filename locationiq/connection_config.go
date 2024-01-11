package locationiq

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type locationiqConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &locationiqConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) locationiqConfig {
	if connection == nil || connection.Config == nil {
		return locationiqConfig{}
	}
	config, ok := connection.Config.(locationiqConfig)
	if !ok {
		plugin.Logger(context.Background()).Error("casting config is not ok")
	}
	return config
}
