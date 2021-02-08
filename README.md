# planeta

Package ```planeta``` provides geospatial utilities for Go. It is intended to be a stand-alone, lightly modified version of [cockroach/pkg/geo](https://github.com/cockroachdb/cockroach/tree/master/pkg/geo). A work in progress.

## Install

To install the library:

```go get github.com/engelsjk/planeta```

### GEOS

Package ```planeta``` is dependent on modified versions of the GEOS libraries that have been custom-built by Cockroach.

To install these libraries, follow the instructions provided by Cockroach ([mac](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-mac.html) / [linux](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-linux.html)), specifically section *Download the Binary*, steps 1-3.

## Usage

In development.

## License

Usage of this library must follow the requirements of the Cockroach BSL license which is copied here as ```licenses/BSL.txt```.
