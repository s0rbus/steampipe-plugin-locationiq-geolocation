package locationiq

import (
	"context"
	"fmt"
	"os"

	liq "github.com/location-iq/locationiq-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func GetToken(ctx context.Context, d *plugin.QueryData) (string, error) {
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

func GetAuthContext(ctx context.Context, token string) context.Context {
	auth := context.WithValue(context.Background(), liq.ContextAPIKey, liq.APIKey{
		Key: token,
	})
	return auth
}

func GetSearchService(ctx context.Context, token string) *liq.SearchApiService {
	liqconfig := liq.NewConfiguration()
	plugin.Logger(ctx).Info("liq config", "default", fmt.Sprintf("%+v", liqconfig))
	client := liq.NewAPIClient(liqconfig)
	return client.SearchApi
}
