# planeta

Package ```planeta``` provides geospatial utilities for Go. It is intended to be an unaffiliated, stand-alone copy of [cockroach/pkg/geo](https://github.com/cockroachdb/cockroach/tree/master/pkg/geo), a library that powers the spatial data support in CockroachDB.

A work in progress.

## Install

To install the library:

```go get github.com/engelsjk/planeta```

### GEOS

Package ```planeta``` is dependent on the GEOS libraries for some spatial functions. Modified versions of these GEOS libraries have been custom-built by Cockroach and need to be installed. 

To install these libraries, follow the instructions ([mac](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-mac.html) / [linux](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-linux.html)) in *Download the Binary*, steps 1-3.

## Usage & Examples

In development.

## Not Implemented Yet

The following sub-packages from [cockroach/pkg/geo](https://github.com/cockroachdb/cockroach/tree/master/pkg/geo) have not been implemented in ```planeta``` yet.

* geoindex
* geoproj
* geotransform

So far, one new sub-package has been added (```geojsonext```) to extend handling of GeoJSON properties.

## License

Usage of this library must follow the requirements of the Cockroach BSL license which is copied here as ```licenses/BSL.txt```.
