package shodan

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type shodanConfig struct {
	APIKey *string `hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &shodanConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) shodanConfig {
	if connection == nil || connection.Config == nil {
		return shodanConfig{}
	}
	config, _ := connection.Config.(shodanConfig)
	return config
}
