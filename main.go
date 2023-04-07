package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-shodan/shodan"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: shodan.Plugin})
}
