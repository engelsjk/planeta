# planeta

Package ```planeta``` provides geospatial utilities for Go.

It is a lightly modified, stand-alone copy of [cockroach/pkg/geo](https://github.com/cockroachdb/cockroach/tree/master/pkg/geo), the library that powers spatial data support in CockroachDB.

## Key Features

* **Geometry types** - OGC-style geometries based on [twpayne/go-geom](https://github.com/twpayne/go-geom/) 
* **Spatial types** - supports both geometric and geographic spatial types
* **Extended spatial functionality** - Includes bindings to GEOS (binary predicates, buffers, etc) and PROJ (projection transformations)  
* **Encoding and decoding support** - EWKB, EWKB Hex, EWKT, GeoJSON, Geohash, encoded polylines and KML (encoding only).

## Geometry types

Geometry types in ```planeta``` are based on the types provided in the [```twpayne/go-geom``` package](https://github.com/twpayne/go-geom/), which support "OpenGeo Consortium-style geometries, 2D and 3D geometries, measures (time and/or distance), and unlimited extra dimensions".

Geometries and geographies in ```planeta``` can be easily converted into ```geom.T``` objects, the generic interface implemented by all geometry types in ```go-geom```, which should provide compatibility with all features in ```twpyane/go-geom```.

## Install

To install the library:

```go get github.com/engelsjk/planeta```

### GEOS

Some functionality in ```planeta``` is dependent on custom-built versions of the GEOS library.

To install this library, follow the instructions ([mac](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-mac.html) / [linux](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-linux.html)) in *Download the Binary*, steps 1-3.

### PROJ

Projection transformations require the PROJ library to already be installed. Instructions can be found [here](https://proj.org/install.html).

## Usage & Examples

link to go.pkgdocs f or the following...

geo.go
geomfn
geogfn

### Parsing
### Binary Predicates
### Transform
### Segmentize
### Decoding

## License

Usage of this library must follow the requirements of the Cockroach BSL license which is copied here as ```licenses/BSL.txt```.
