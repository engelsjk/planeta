package main

import (
	"encoding/json"
	"fmt"

	"github.com/engelsjk/planeta/geo"
	"github.com/engelsjk/planeta/geo/geojsonext"
	"github.com/engelsjk/planeta/geo/geos"
	"github.com/twpayne/go-geom/encoding/geojson"
)

func main() {

	s := `{
		"type": "Feature",
		"properties": {
			"name": "example"
		},
		"geometry": {
		  "type": "Polygon",
		  "coordinates": [
			[
			  [
				-83.023681640625,
				39.104488809440475
			  ],
			  [
				-81.968994140625,
				39.104488809440475
			  ],
			  [
				-81.968994140625,
				39.85072092501597
			  ],
			  [
				-83.023681640625,
				39.85072092501597
			  ],
			  [
				-83.023681640625,
				39.104488809440475
			  ]
			]
		  ]
		}
	  }`

	f := &geojson.Feature{}
	err := json.Unmarshal([]byte(s), f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", f.Properties)

	if name, err := geojsonext.PropertyString(f.Properties, "name"); err == nil {
		fmt.Printf("name: %s\n", name)
	} else {
		panic(err)
	}

	g, err := geo.MakeGeometryFromGeomT(f.Geometry)
	if err != nil {
		panic(err)
	}

	if isValid, err := geos.IsValid(g.EWKB()); err == nil {
		fmt.Printf("isvalid: %t\n", isValid)
	} else {
		panic(err)
	}
}