package shodan

import (
	"context"
	"net/http"
	"os"
	"strings"

	shodan "github.com/shadowscatcher/shodan"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*shodan.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "shodan"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*shodan.Client), nil
	}

	var apiKey string

	// First, use the api_key config
	shodanConfig := GetConfig(d.Connection)
	if &shodanConfig != nil {
		if shodanConfig.APIKey != nil {
			apiKey = *shodanConfig.APIKey
		}
	}

	// Otherwise, default to using SHODAN_KEY env var
	if apiKey == "" {
		apiKey = os.Getenv("SHODAN_API_KEY")
	}

	// Configure to automatically wait 1 sec between requests, per Shodan API requirements
	conn, err := shodan.GetClient(apiKey, http.DefaultClient, true)
	if err != nil {
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func isErrorWithMessage(err error, messages []string) bool {
	errMsg := err.Error()
	for _, msg := range messages {
		if strings.Contains(errMsg, msg) {
			return true
		}
	}
	return false
}

func queryString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	q := quals["query"].GetStringValue()
	return q, nil
}
