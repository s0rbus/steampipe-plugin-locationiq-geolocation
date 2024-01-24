---
organization: s0rbus
category: ["software development"]
icon_url: "/images/plugins/s0rbus/locationiq.svg"
brand_color: "#FF5C5C"
display_name: "LocationIQ"
short_name: "locationiq"
description: "Steampipe plugin for querying geolocation data from LocationIQ."
og_description: "Query LocationIQ geolocation service with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/s0rbus/locationiq-social-graphic.png"
---

# LocationIQ + Steampipe

[LocationIQ](https://locationiq.com/ offering a range of location based services to developers, focussed primarily on providing affordable and scalable APIs for Geocoding, Maps and Routing.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql

```

```

```

## Documentation

- **[Table definitions & examples →](https://hub.steampipe.io/plugins/s0rbus/locationiq/tables)**

## Get started

### Install

Download and install the latest LocationIQ plugin:

```bash
steampipe plugin install s0rbus/locationiq
```

### Credentials

| Item        | Description                                                                                                            |
| :---------- | :--------------------------------------------------------------------------------------------------------------------- |
| Credentials | LocationIQ requires registration and an [API token](https://docs.locationiq.com/docs/authentication) for all requests. |

### Configuration

Installing the latest LocationIQ plugin will create a config file (`~/.steampipe/config/locationiq.spc`) with a single connection named `locationiq`:

```hcl
connection "location iq" {
    plugin    = "s0rbus/locationiq"
    # token   =
}
```

- `token` - [API token](https://docs.locationiq.com/docs/authentication)

## Get involved

- Open source: https://github.com/s0rbus/steampipe-plugin-locationiq
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
