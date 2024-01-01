# Steampipe plugin for LocatioIQ's gelocation API

https://locationiq.com/

Use SQL to get location (lat/long) from town/city names worldwide.

The plugin utilises the Go client provided by LocationIQ - https://github.com/location-iq/locationiq-go-client

LocationIQ provides a free tier, but you have to register and apply for an API token. Free tier use is limited to 5,000 requests per day. Rate limiting to address this has not yet been implemented/configured in the plugin.

Once installed, run a query (you must provide a placename, postcode etc):
```
```
