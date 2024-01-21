# Steampipe plugin for LocatioIQ's gelocation API

https://locationiq.com/

Use SQL to get location (lat/long) from town/city names worldwide.

The plugin utilises the Go client provided by LocationIQ - https://github.com/location-iq/locationiq-go-client

LocationIQ provides a free tier, but you have to register and apply for an API token. Free tier use is limited to 5,000 requests per day. Rate limiting to address this has not yet been implemented/configured in the plugin.

## Quick start

Install the plugin with [Steampipe](https://steampipe.io/downloads):

```shell
steampipe plugin install locationiq
```

[Configure the plugin](https://hub.steampipe.io/plugins/path/to/locationiq#configuration) using the configuration file:

```shell
vi ~/.steampipe/locationiq.spc
```

Or environment variables:

```shell
export LOCATIONIQ_TOKEN=liq_YOURTOKENHERE
```

Start Steampipe:

```shell
steampipe query
```

Run a query:

```sql
select
    latitude,
    longitude
from
    locationiq_place2latlong
where
    placequery = 'Trafalgar Square, London';
```

```
+-----------+----------------------+
| latitude  | longitude            |
+-----------+----------------------+
| 51.508037 | -0.12804941070724718 |
+-----------+----------------------+
```

Or a reverse search:

```sql
select
    address,
    postcode,
    distance
from
    locationiq_latlong2place
where
    latitude = 51.5053
and
    longitude = -0.0755;
```

```
(some of the address has been snipped for clarity)
+-------------------------------------------------------------------------------------------------------+----------+----------+
| address                                                                                               | postcode | distance |
+-------------------------------------------------------------------------------------------------------+----------+----------+
| Tower Bridge, Horsleydown Old Stairs, Butler's Wharf, Bermondsey, London Borough of Southwark, London | SE1 2LY  | 25       |
+-------------------------------------------------------------------------------------------------------+----------+----------+

```

You can also retrieve your request balance:

```sql
select
    balance,
    bonus
from
    locationiq_balance;
```

