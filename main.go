package main

import (
    "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
    "github.com/location-iq/locationiq-go-client"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{PluginFunc: locationiq.Plugin})
}
