package locationiq

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/antihax/optional"
	liq "github.com/location-iq/locationiq-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLocationIQLocation() *plugin.Table {
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
				Name:        "importance",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Importance"),
				Description: "match importance",
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
			{
				Name:        "match",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("MatchCode"),
				Description: "match level",
			},
		},
	}

}

func getLocation(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	token, err := GetToken(ctx, d)
	if err != nil {
		return nil, err
	}
	authContext := GetAuthContext(ctx, token)
	service := GetSearchService(ctx, token)
	pquery := d.EqualsQuals["placequery"].GetStringValue()
	pcquery := d.EqualsQuals["postcodequery"].GetStringValue()
	var loc []liq.Location
	var resp *http.Response
	opts := &liq.SearchOpts{
		Matchquality: optional.NewInt32(1),
		Limit:        optional.NewInt32(1),
	}
	if pquery != "" {
		loc, resp, err = service.Search(authContext, pquery, "JSON", 1, opts)
	} else if pcquery != "" {
		loc, resp, err = service.Search(authContext, pcquery, "JSON", 1, opts)
	} else {
		return nil, fmt.Errorf("no query params")
	}
	if err != nil {
		plugin.Logger(ctx).Error("getting location", "err", err)
		return nil, err
	}
	plugin.Logger(ctx).Debug("getting location", "resp status", resp.Status)
	plugin.Logger(ctx).Debug("getting location", "locs len", len(loc))

	type Row struct {
		Address    string
		Postcode   string
		Lat        float64
		Long       float64
		MatchCode  string
		Importance float64
	}
	for i, l := range loc {

		lat, laterr := strconv.ParseFloat(l.Lat, 64)
		lng, lngerr := strconv.ParseFloat(l.Lon, 64)
		if laterr != nil || lngerr != nil {
			plugin.Logger(ctx).Error("getting loc", "lat parse error", laterr, "lng parse error", lngerr)
			continue
		}
		plugin.Logger(ctx).Debug("getting location", "idx", i, "display name", l.DisplayName, "pc", l.Address.Postcode, "lat", lat, "long", lng)
		d.StreamListItem(ctx, Row{
			Address:    l.DisplayName,
			Postcode:   l.Address.Postcode,
			Lat:        lat,
			Long:       lng,
			MatchCode:  l.Matchquality.Matchcode,
			Importance: float64(l.Importance),
		})
	}

	return nil, nil
}
