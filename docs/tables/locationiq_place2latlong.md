# Table: locationiq_place2latlong

Latitude and longitude for a given location

## Examples

### Get lat/long for Cambridge, England

```sql
select
   latitude,
   longitude
from
   locationiq_place2latlong
where
   placequery = 'Cambridge, England';
```

```
+------------+-----------+
| latitude   | longitude |
+------------+-----------+
| 52.2055314 | 0.1186637 |
+------------+-----------+
```

### Or you can be more specific
```sql
select
   latitude,
   longitude
from
   locationiq_place2latlong
where
   placequery = '20 Trinity St, Cambridge, England';
```
