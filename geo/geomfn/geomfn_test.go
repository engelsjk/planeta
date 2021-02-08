// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package geomfn

import (
	"testing"

	"github.com/engelsjk/planeta/geo"
	"github.com/stretchr/testify/require"
)

var mismatchingSRIDGeometryA = geo.MustParseGeometry("SRID=4004;POINT(1.0 1.0)")
var mismatchingSRIDGeometryB = geo.MustParseGeometry("SRID=4326;LINESTRING(1.0 1.0, 2.0 2.0)")

// requireMismatchingSRIDError checks errors fall as expected for mismatching SRIDs.
func requireMismatchingSRIDError(t *testing.T, err error) {
	require.Error(t, err)
	require.EqualError(t, err, `operation on mixed SRIDs forbidden: (Point, 4004) != (LineString, 4326)`)
}

func TestRemoveConsecutivePointsFromGeomT(t *testing.T) {
	testCases := []struct {
		desc         string
		inEWKT       string
		expectedEWKT string
	}{
		{
			desc:         "POINT",
			inEWKT:       "POINT(1 2)",
			expectedEWKT: "POINT(1 2)",
		},
		{
			desc:         "LINESTRING",
			inEWKT:       "LINESTRING(1 1, 2 2, 2 2, 2 2, 3 3, 3 3)",
			expectedEWKT: "LINESTRING(1 1, 2 2, 3 3)",
		},
		{
			desc:         "LINESTRING collapsed into empty string",
			inEWKT:       "LINESTRING(1 1, 1 1, 1 1)",
			expectedEWKT: "LINESTRING EMPTY",
		},
		{
			desc:         "POLYGON",
			inEWKT:       "POLYGON((0 0, 1 0, 1 1, 1 1, 0 0), (0.1 0.1, 0.2 0.2, 0.2 0.3, 0.2 0.3, 0.1 0.1))",
			expectedEWKT: "POLYGON((0 0, 1 0, 1 1, 0 0), (0.1 0.1, 0.2 0.2, 0.2 0.3, 0.1 0.1))",
		},
		{
			desc:         "POLYGON, collapsed hole",
			inEWKT:       "POLYGON((0 0, 1 0, 1 1, 1 1, 0 0), (0.1 0.1, 0.1 0.1, 0.1 0.1, 0.1 0.1), (0.1 0.1, 0.2 0.2, 0.2 0.3, 0.2 0.3, 0.1 0.1))",
			expectedEWKT: "POLYGON((0 0, 1 0, 1 1, 0 0), (0.1 0.1, 0.2 0.2, 0.2 0.3, 0.1 0.1))",
		},
		{
			desc:         "POLYGON, collapsed base",
			inEWKT:       "POLYGON((0 0, 1 1, 1 1, 1 1, 0 0), (0.1 0.1, 0.1 0.1, 0.1 0.1, 0.1 0.1), (0.1 0.1, 0.2 0.2, 0.2 0.3, 0.2 0.3, 0.1 0.1))",
			expectedEWKT: "POLYGON EMPTY",
		},
		{
			desc:         "MULTIPOINT",
			inEWKT:       "MULTIPOINT(0 0, 1 1)",
			expectedEWKT: "MULTIPOINT(0 0, 1 1)",
		},
		{
			desc:         "MULTILINESTRING, some collapses",
			inEWKT:       "MULTILINESTRING((1 1, 2 2, 2 2, 3 3), (2 2, 2 2))",
			expectedEWKT: "MULTILINESTRING((1 1, 2 2, 3 3))",
		},
		{
			desc:         "MULTILINESTRING, all collapses",
			inEWKT:       "MULTILINESTRING((2 2, 2 2, 2 2, 2 2), (2 2, 2 2))",
			expectedEWKT: "MULTILINESTRING EMPTY",
		},
		{
			desc: "MULTIPOLYGON, some collapses",
			inEWKT: `MULTIPOLYGON(
				((0 0, 1 0, 1 1, 1 1, 0 0), (0.1 0.1, 0.2 0.2, 0.2 0.3, 0.2 0.3, 0.1 0.1)),
				((0 0, 1 1, 1 1, 1 1, 0 0), (0.1 0.1, 0.1 0.1, 0.1 0.1, 0.1 0.1), (0.1 0.1, 0.2 0.2, 0.2 0.3, 0.2 0.3, 0.1 0.1))
			)`,
			expectedEWKT: `MULTIPOLYGON(
				((0 0, 1 0, 1 1, 0 0), (0.1 0.1, 0.2 0.2, 0.2 0.3, 0.1 0.1))
			)`,
		},
		{
			desc: "MULTIPOLYGON, all collapses",
			inEWKT: `MULTIPOLYGON(
				((0 0, 0 0, 1 1, 1 1, 0 0), (0.1 0.1, 0.2 0.2, 0.2 0.3, 0.2 0.3, 0.1 0.1)),
				((0 0, 1 1, 1 1, 1 1, 0 0), (0.1 0.1, 0.1 0.1, 0.1 0.1, 0.1 0.1), (0.1 0.1, 0.2 0.2, 0.2 0.3, 0.2 0.3, 0.1 0.1))
			)`,
			expectedEWKT: `MULTIPOLYGON EMPTY`,
		},
		{
			desc: "GEOMETRYCOLLECTION",
			inEWKT: `GEOMETRYCOLLECTION(
	POLYGON((0 0, 1 0, 1 1, 1 1, 0 0), (0.1 0.1, 0.2 0.2, 0.2 0.3, 0.2 0.3, 0.1 0.1)),
	LINESTRING(1 1, 1 1, 1 1),
	MULTIPOINT(0 0, 1 1)
)`,
			expectedEWKT: `GEOMETRYCOLLECTION(
	POLYGON((0 0, 1 0, 1 1, 0 0), (0.1 0.1, 0.2 0.2, 0.2 0.3, 0.1 0.1)),
	MULTIPOINT(0 0, 1 1)
)`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			in, err := geo.ParseGeometry(tc.inEWKT)
			require.NoError(t, err)
			inT, err := in.AsGeomT()
			require.NoError(t, err)

			expected, err := geo.ParseGeometry(tc.expectedEWKT)
			require.NoError(t, err)
			expectedT, err := expected.AsGeomT()
			require.NoError(t, err)

			actual, err := removeConsecutivePointsFromGeomT(inT)
			require.NoError(t, err)
			require.Equal(t, expectedT, actual)
		})
	}
}
