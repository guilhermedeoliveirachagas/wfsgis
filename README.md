# WFS 3.0 FES Hackathon Repo
Working repo for the 2018 WFS Hackathon in Ft. Collins

# REST Examples

* Create a new collection

```
curl -X POST \
  http://localhost:8080/collections \
  -H 'Accept: */*' \
  -H 'Content-Type: application/json' \
  -d '{
	"name": "test6",
	"title": "test 6",
	"description": "test 6 test 6",
	"crs": ["EPSG4326"]
}
'
```

* Get collection list

```
curl -X GET \
  http://localhost:8080/collections \
  -H 'Accept: */*' \
```

* Create a new Feature (from a complex GeoJSON)

```
curl -X POST \
  http://localhost:8080/collections/test6/items \
  -H 'Accept: */*' \
  -H 'Content-Type: application/json' \
  -d '{
   "type":"FeatureCollection",
   "features":[
      {
         "type":"Feature",
         "geometry":{
            "type":"Point",
            "coordinates":[
               102.0,
               0.5
            ]
         },
         "properties":{
            "prop0":"value0"
         },
         "when": {
         	"@type": "Instant",
         	"datetime": "2019-02-01T15:04:02Z"
         }
      },
      {
         "type":"Feature",
         "geometry":{
            "type":"LineString",
            "coordinates":[
               [
                  102.0,
                  0.0
               ],
               [
                  103.0,
                  1.0
               ],
               [
                  104.0,
                  0.0
               ],
               [
                  105.0,
                  1.0
               ]
            ]
         },
         "properties":{
            "prop0":"value1",
            "prop1":0.0
         }
      },
      {
         "type":"Feature",
         "geometry":{
            "type":"Polygon",
            "coordinates":[
               [
                  [
                     100.0,
                     0.0
                  ],
                  [
                     101.0,
                     0.0
                  ],
                  [
                     101.0,
                     1.0
                  ],
                  [
                     100.0,
                     1.0
                  ],
                  [
                     100.0,
                     0.0
                  ]
               ]
            ]
         },
         "properties":{
            "prop0":"value2",
            "prop1":{
               "this":"that"
            }
         },
        "when": {
            "@type": "Instant",
            "datetime": "2019-12-11T10:02:44Z"
        }
      }
   ]
}'
```

* Query Features by bounding box, time range and property value

```
curl -X GET \
  'http://localhost:8080/collections/test6/items?bbox=101,0,103,1&limit=10&time=2019-01-01T15:04:02Z/2019-12-31T15:04:02Z&prop0=value2' \
  -H 'Accept: */*' \
  -H 'cache-control: no-cache'
```

* Get a item

```
curl -X GET \
  'http://localhost:8080/collections/test6/items/2' \
  -H 'Accept: */*'
```

* Delete item

```
curl -X DELETE \
  http://localhost:8080/collections/test6/items/2 \
  -H 'cache-control: no-cache'
```

