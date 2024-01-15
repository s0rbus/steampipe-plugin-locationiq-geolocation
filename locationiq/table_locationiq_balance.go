package locationiq

import (
	"context"
	"fmt"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLocationIQBalance() *plugin.Table {
	return &plugin.Table{
		Name:        "locationiq_balance",
		Description: "Get LocationIQ account request balance",
		List: &plugin.ListConfig{
			Hydrate: getBalance,
		},
		Columns: []*plugin.Column{
			{
				Name:        "balance",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Balance"),
				Description: "day balance",
			},
			{
				Name:        "bonus",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Bonus"),
				Description: "bonus",
			},
		},
	}

}

func getBalance(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	liqAdminData, err := NewAdminData(ctx, d)
	if err != nil {
		return nil, err
	}
	if err != nil {
		plugin.Logger(ctx).Error("getting lociq admindata", "err", err)
		return nil, err
	}
	bal, err := liqAdminData.GetBalance(liqAdminData.AuthContext, liqAdminData.Token)
	if err != nil {
		plugin.Logger(ctx).Error("getting lociq balance", "err", err)
		return nil, err
	}
	if strings.ToLower(bal.Status) != "ok" {
		return nil, fmt.Errorf("getting lociq balance returned status %v", bal.Status)
	}

	type Row struct {
		Balance int
		Bonus   int
	}
	d.StreamListItem(ctx, Row{
		Balance: int(bal.Balance.Day),
		Bonus:   int(bal.Balance.Bonus),
	})

	return nil, nil
}
