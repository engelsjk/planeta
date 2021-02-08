# planeta

Planeta is a geospatial library for Go. A work in progress.

It's a stand-alone, lightly modified version of the Cockroach [geo package](https://github.com/cockroachdb/cockroach/tree/master/pkg/geo) that drives the spatial data functionality in CockroachDB.

## Install

To install the library:

```go get github.com/engelsjk/planeta```

Note that most of the functionality in Planeta is dependent on versions of the GEOS libraries that have been custom-built by Cockroach.

To install these libraries, follow the instructions provided by Cockroach ([mac](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-mac.html) / [linux](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-linux.html)), specifically the section *Download the Binary*, steps 1-3.

## Usage

In development.

## License

Usage of this library must follow the requirements of the Cockroach BSL license which is copied here as ```licenses/BSL.txt```.
