package locationiq

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-locationiq",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"locationiq_place2latlong": tableLocationIQPlace2Latlong(),
			"locationiq_latlong2place": tableLocationIQLatlong2place(),
			"locationiq_balance":       tableLocationIQBalance(),
		},
	}
	return p
}
