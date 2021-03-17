package main

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-shodan/shodan"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: shodan.Plugin})
}
