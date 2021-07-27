package main

import (
	"github.com/turbot/steampipe-plugin-azuread/azuread"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: azuread.Plugin})
}
