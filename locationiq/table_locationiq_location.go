package locationiq

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableLocationIQLocation() *plugin.Table {
	/* auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
		Key: "APIKEY",
		Prefix: "Bearer", // Omit if not necessary.
	})
	r, err := client.Service.Operation(auth, args) */

	//apiKey :=

	return &plugin.Table{
		Name:        "locationiq_location",
		Description: "Get lat/long from place name",
		List: &plugin.ListConfig{
			Hydrate: getLocation,
		},
		Columns: []*plugin.Column{
			{
				Name:        "lat",
				Type:        proto.ColumnType_DOUBLE,
				Description: "latitude",
			},
			{
				Name:        "long",
				Type:        proto.ColumnType_DOUBLE,
				Description: "longitude",
			},
		},
	}

}

func getLocation(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	token := GetToken(ctx, d)
	plugin.Logger(ctx).Info("getting api token", "value", token)
	type Row struct {
		Lat  float64
		Long float64
	}
	d.StreamListItem(ctx, Row{
		Lat:  1.2,
		Long: 3.4,
	})
	return nil, nil
}
