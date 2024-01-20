package locationiq

import (
	"context"
	"fmt"
	"net/http"
	"os"

	liq "github.com/location-iq/locationiq-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type AdminData struct {
	Token       string
	Config      *liq.Configuration
	Client      *liq.APIClient
	AuthContext context.Context
}

func NewAdminData(ctx context.Context, d *plugin.QueryData) (*AdminData, error) {
	t, err := getToken(ctx, d)
	if err != nil {
		return nil, err
	}
	liqconfig := liq.NewConfiguration()
	client := liq.NewAPIClient(liqconfig)
	ac := getAuthContext(ctx, t)
	ad := &AdminData{
		Token:       t,
		Config:      liqconfig,
		Client:      client,
		AuthContext: ac,
	}
	return ad, nil
}

func getToken(ctx context.Context, d *plugin.QueryData) (string, error) {
	token := os.Getenv("LOCATIONIQ_TOKEN")
	if token == "" {
		locConfig := GetConfig(d.Connection)
		if locConfig.Token != nil {
			token = *locConfig.Token
		} else {
			return "", fmt.Errorf("could not retrieve API token from config")
		}
	}

	if token == "" {
		//plugin.Logger(ctx).Error("token must be set somewhere")
		return "", fmt.Errorf("LocationIQ API token is not set")
	}
	return token, nil
}

func getAuthContext(ctx context.Context, token string) context.Context {
	auth := context.WithValue(context.Background(), liq.ContextAPIKey, liq.APIKey{
		Key: token,
	})
	return auth
}

/* func GetClient(config *liq.Configuration) *liq.APIClient {
	return liq.NewAPIClient(config)
} */

func (ad AdminData) GetSearchService(ctx context.Context, token string) *liq.SearchApiService {
	//liqconfig := liq.NewConfiguration()
	plugin.Logger(ctx).Info("liq config", "default", fmt.Sprintf("%+v", ad.Config))
	//client := liq.NewAPIClient(liqconfig)
	return ad.Client.SearchApi
}

func (ad AdminData) GetReverseSearchService(ctx context.Context, token string) *liq.ReverseApiService {
	return ad.Client.ReverseApi
}

func (ad AdminData) GetBalance(ctx context.Context, token string) (liq.Balance, error) {
	bs := ad.Client.BalanceApi
	b, r, err := bs.Balance(ctx)
	if err != nil {
		return liq.Balance{}, err
	}
	if r.StatusCode != http.StatusOK {
		return liq.Balance{}, fmt.Errorf("balance request returned status code %v", r.StatusCode)
	}
	return b, nil
}
