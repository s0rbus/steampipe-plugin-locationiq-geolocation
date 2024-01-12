package locationiq

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	liq "github.com/location-iq/locationiq-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
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
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "placequery", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "postcodequery", Require: plugin.Optional, Operators: []string{"="}},
			},
			Hydrate: getLocation,
		},
		Columns: []*plugin.Column{
			{
				Name:        "placequery",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("placequery"),
				Description: "the query address",
			},
			{
				Name:        "postcodequery",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("postcodequery"),
				Description: "the query postcode",
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
				Name:        "lat",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Lat"),
				Description: "latitude",
			},
			{
				Name:        "long",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Long"),
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
	addr := d.EqualsQuals["placequery"].GetStringValue()
	pcode := d.EqualsQuals["postcodequery"].GetStringValue()
	plugin.Logger(ctx).Info("get params", "addr", addr, "pcode", pcode)
	var loc []liq.Location
	var resp *http.Response
	var err error
	if addr != "" {
		loc, resp, err = client.SearchApi.Search(authContext, addr, "JSON", 1, nil)
	} else if pcode != "" {
		loc, resp, err = client.SearchApi.Search(authContext, pcode, "JSON", 1, nil)
	} else {
		return nil, fmt.Errorf("no query params")
	}
	if err != nil {
		plugin.Logger(ctx).Error("getting location", "err", err)
		return nil, err
	}
	plugin.Logger(ctx).Info("getting location", "resp status", resp.Status)
	plugin.Logger(ctx).Info("getting location", "locs len", len(loc))

	type Row struct {
		Address  string
		Postcode string
		Lat      float64
		Long     float64
	}
	for i, l := range loc {
		lat, laterr := strconv.ParseFloat(l.Lat, 64)
		lng, lngerr := strconv.ParseFloat(l.Lon, 64)
		if laterr != nil || lngerr != nil {
			plugin.Logger(ctx).Error("getting loc", "lat parse error", laterr, "lng parse error", lngerr)
			continue
		}
		plugin.Logger(ctx).Info("getting location", "idx", i, "display name", l.DisplayName, "pc", l.Address.Postcode, "lat", lat, "long", lng)
		d.StreamListItem(ctx, Row{
			Address:  l.DisplayName,
			Postcode: l.Address.Postcode,
			Lat:      lat,
			Long:     lng,
		})
	}

	return nil, nil
}
