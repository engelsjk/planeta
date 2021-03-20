# planeta

Package ```planeta``` provides geospatial utilities for Go.

It is a lightly modified, stand-alone copy of [cockroach/pkg/geo](https://github.com/cockroachdb/cockroach/tree/master/pkg/geo), the library that powers spatial data support in CockroachDB.

A work in progress.

## Install

To install the library:

```go get github.com/engelsjk/planeta```

### GEOS

Some functionality in ```planeta``` is dependent on custom-built versions of the GEOS library.

To install this library, follow the instructions ([mac](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-mac.html) / [linux](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-linux.html)) in *Download the Binary*, steps 1-3.

### PROJ

Projection transformations require the PROJ library to be installed. Instructions can be found [here](https://proj.org/install.html).

## Usage & Examples

In development.

## License

Usage of this library must follow the requirements of the Cockroach BSL license which is copied here as ```licenses/BSL.txt```.
