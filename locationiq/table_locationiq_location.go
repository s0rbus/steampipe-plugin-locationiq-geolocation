package locationiq

import (
	"context"
	"strconv"

	liq "github.com/location-iq/locationiq-go-client"
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
	authContext := GetAuth(ctx, token)
	liqconfig := liq.NewConfiguration()
	//liqconfig.
	client := liq.NewAPIClient(liqconfig)
	loc, resp, err := client.SearchApi.Search(authContext, "Martlesham", "JSON", 1, nil)
	if err != nil {
		plugin.Logger(ctx).Error("getting location", "err", err)
		return nil, err
	}
	plugin.Logger(ctx).Info("getting location", "resp status", resp.Status)
	plugin.Logger(ctx).Info("getting location", "locs len", len(loc))

	type Row struct {
		Lat  float64
		Long float64
	}
	for i, l := range loc {
		lat, laterr := strconv.ParseFloat(l.Lat, 64)
		lng, lngerr := strconv.ParseFloat(l.Lon, 64)
		if laterr != nil || lngerr != nil {
			plugin.Logger(ctx).Error("getting loc", "lat parse error", laterr, "lng parse error", lngerr)
			continue
		}
		plugin.Logger(ctx).Info("getting location", "idx", i, "display name", l.DisplayName)
		d.StreamListItem(ctx, Row{
			Lat:  lat,
			Long: lng,
		})
	}

	return nil, nil
}
