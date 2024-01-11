package main

import (
	//"github.com/location-iq/locationiq-go-client"
	"github.com/s0rbus/steampipe-plugin-locationiq/locationiq"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: locationiq.Plugin})
}
