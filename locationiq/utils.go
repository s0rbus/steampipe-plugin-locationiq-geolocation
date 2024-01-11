package locationiq

import (
	"context"
	"os"

	liq "github.com/location-iq/locationiq-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func GetToken(ctx context.Context, d *plugin.QueryData) string {
	token := os.Getenv("LOCATIONIQ_TOKEN")
	if token == "" {
		locConfig := GetConfig(d.Connection)
		if locConfig.Token != nil {
			token = *locConfig.Token
		}
	}

	if token == "" {
		plugin.Logger(ctx).Error("token must be set somewhere")
	}
	liqconfig := liq.NewConfiguration()
	plugin.Logger(ctx).Info("liq config", "value", liqconfig)

	return token
}
