# planeta

Package ```planeta``` provides geospatial utilities for Go.

It is a lightly modified, stand-alone fork of CockroachDB's [cockroach/pkg/geo](https://github.com/cockroachdb/cockroach/tree/master/pkg/geo) subpackage which provides PostgreSQL/PostGIS spatial data compatability. 

## Key Features

* **Geometry types** - OGC-style geometries based on [twpayne/go-geom](https://github.com/twpayne/go-geom/) 
* **Spatial types** - supports both geometric and geographic spatial types
* **Extended spatial functionality** - Includes bindings to GEOS (binary predicates, buffers, etc) and PROJ (projection transformations)  
* **Encoding and decoding support** - EWKB, EWKB Hex, EWKT, GeoJSON, Geohash, encoded polylines and KML (encoding only).

## Geometry types

Geometry types in ```planeta``` are based on the types provided in the [```twpayne/go-geom``` package](https://github.com/twpayne/go-geom/), which support "OpenGeo Consortium-style geometries, 2D and 3D geometries, measures (time and/or distance), and unlimited extra dimensions".

* Point, MultiPoint
* LineString, MultiLineString
* Polygon, MultiPolygon
* Geometry, GeometryCollection
* 2D, Z, M and ZM

Geometries and geographies in ```planeta``` can be easily converted into ```go-geom```'s generic spatial interface objects (```geom.T```) to provide compatibility with all features in ```twpyane/go-geom```.

## Install

To install the library:

```go get github.com/engelsjk/planeta```

### GEOS

Some functionality in ```planeta``` is dependent on CockroachDB-built custom versions of the GEOS library.

To install this library, follow the instructions ([mac](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-mac.html) / [linux](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-linux.html)) in CockroachDB's *Download the Binary*, steps 1-3.

### PROJ

Projection transformations require installation of the PROJ library. Instructions can be found [here](https://proj.org/install.html).

## Usage Examples

For usage examples, see [```engelsjk/planeta-examples```](https://github.com/engelsjk/planeta-examples/).

## License

Usage of this library must follow the requirements of the Cockroach BSL license which is copied here as ```licenses/BSL.txt```.
