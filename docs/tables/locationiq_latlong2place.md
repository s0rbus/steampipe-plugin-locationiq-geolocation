# Table: locationiq_latlong2place

Address for a given latitude and longitude. This will be the nearest address as determined by the API. A distance column is also returned which is the straight line distance in meters between the returned address and the given lat/long position.

## Examples

### Get an address for a location in London, England

```sql
select
   address,
   distance
from 
   locationiq_latlong2place
where
   latitude = 51.51317
and
   longitude = -0.14047;
```

```
+--------------------------------------------------------------------------------------------------------------------------------------+----------+
| address                                                                                                                              | distance |
+--------------------------------------------------------------------------------------------------------------------------------------+----------+
| 200-206, Regent Street, East Marylebone, Fitzrovia, Islington, City of Westminster, Greater London, England, W1B 5BD, United Kingdom | 17       |
+--------------------------------------------------------------------------------------------------------------------------------------+----------+
```
