package locationiq

import (
	"context"
	"fmt"
	"net/http"

	//"strconv"

	"github.com/antihax/optional"
	liq "github.com/location-iq/locationiq-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLocationIQLatlong2place() *plugin.Table {
	return &plugin.Table{
		Name:        "locationiq_latlong2place",
		Description: "Get place name from lat/long",
		List: &plugin.ListConfig{
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "latitude", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "longitude", Require: plugin.Optional, Operators: []string{"="}},
			},
			Hydrate: getPlace,
		},
		Columns: []*plugin.Column{
			{
				Name:        "latitude",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromQual("latitude"),
				Description: "latitude",
			},
			{
				Name:        "longitude",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromQual("longitude"),
				Description: "longitude",
			},
			{
				Name:        "address",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Address"),
				Description: "the address",
			},
			{
				Name:        "postcode",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Postcode"),
				Description: "the postcode",
			},
			{
				Name:        "distance",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Distance"),
				Description: "distance from requested point",
			},
		},
	}

}

func getPlace(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	liqAdminData, err := NewAdminData(ctx, d)
	if err != nil {
		return nil, err
	}
	service := liqAdminData.GetReverseSearchService(ctx, liqAdminData.Token)
	lat := d.EqualsQuals["latitude"].GetDoubleValue()
	lng := d.EqualsQuals["longitude"].GetDoubleValue()
	var loc liq.Location
	var resp *http.Response
	opts := &liq.ReverseOpts{
		Addressdetails: optional.NewInt32(1),
		Showdistance:   optional.NewInt32(1),
	}

	loc, resp, err = service.Reverse(liqAdminData.AuthContext, float32(lat), float32(lng), "JSON", 1, opts)

	if err != nil {
		plugin.Logger(ctx).Error("getting location", "err", err)
		return nil, err
	}
	plugin.Logger(ctx).Debug("getting location", "resp status", resp.Status)
	plugin.Logger(ctx).Debug("getting location", "loc all", fmt.Sprintf("%+v", loc))

	type Row struct {
		Address  string
		Postcode string
		Distance float64
	}

	plugin.Logger(ctx).Debug("getting location", "display name", loc.DisplayName, "pc", loc.Address.Postcode, "lat", lat, "long", lng)
	d.StreamListItem(ctx, Row{
		Address:  loc.DisplayName,
		Postcode: loc.Address.Postcode,
		Distance: float64(loc.Distance),
	})

	return nil, nil
}
