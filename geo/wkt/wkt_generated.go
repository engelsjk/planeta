// Code generated by goyacc -o wkt_generated.go -p wkt wkt.y. DO NOT EDIT.

//line wkt.y:12

package wkt

import __yyfmt__ "fmt"

//line wkt.y:13

import "github.com/twpayne/go-geom"

// TODO(ayang): move these into lex.go
func isValidLineString(wktlex wktLexer, flatCoords []float64, stride int) bool {
	if len(flatCoords) < 2*stride {
		wktlex.(*wktLex).setParseError("non-empty linestring with only one point", "minimum number of points is 2")
		return false
	}
	return true
}

func isValidPolygonRing(wktlex wktLexer, flatCoords []float64, stride int) bool {
	if len(flatCoords) < 4*stride {
		wktlex.(*wktLex).setParseError("polygon ring doesn't have enough points", "minimum number of points is 4")
		return false
	}
	for i := 0; i < stride; i++ {
		if flatCoords[i] != flatCoords[len(flatCoords)-stride+i] {
			wktlex.(*wktLex).setParseError("polygon ring not closed", "ensure first and last point are the same")
			return false
		}
	}
	return true
}

type geomFlatCoordsRepr struct {
	flatCoords []float64
	ends       []int
}

func makeGeomFlatCoordsRepr(flatCoords []float64) geomFlatCoordsRepr {
	return geomFlatCoordsRepr{flatCoords: flatCoords, ends: []int{len(flatCoords)}}
}

func appendGeomFlatCoordsReprs(p1 geomFlatCoordsRepr, p2 geomFlatCoordsRepr) geomFlatCoordsRepr {
	if len(p1.ends) > 0 {
		p1LastEnd := p1.ends[len(p1.ends)-1]
		for i, _ := range p2.ends {
			p2.ends[i] += p1LastEnd
		}
	}
	return geomFlatCoordsRepr{flatCoords: append(p1.flatCoords, p2.flatCoords...), ends: append(p1.ends, p2.ends...)}
}

type multiPolygonFlatCoordsRepr struct {
	flatCoords []float64
	endss      [][]int
}

func makeMultiPolygonFlatCoordsRepr(p geomFlatCoordsRepr) multiPolygonFlatCoordsRepr {
	if p.flatCoords == nil {
		return multiPolygonFlatCoordsRepr{flatCoords: nil, endss: [][]int{nil}}
	}
	return multiPolygonFlatCoordsRepr{flatCoords: p.flatCoords, endss: [][]int{p.ends}}
}

func appendMultiPolygonFlatCoordsRepr(
	p1 multiPolygonFlatCoordsRepr, p2 multiPolygonFlatCoordsRepr,
) multiPolygonFlatCoordsRepr {
	p1LastEndsLastEnd := 0
	for i := len(p1.endss) - 1; i >= 0; i-- {
		if len(p1.endss[i]) > 0 {
			p1LastEndsLastEnd = p1.endss[i][len(p1.endss[i])-1]
			break
		}
	}
	if p1LastEndsLastEnd > 0 {
		for i, _ := range p2.endss {
			for j, _ := range p2.endss[i] {
				p2.endss[i][j] += p1LastEndsLastEnd
			}
		}
	}
	return multiPolygonFlatCoordsRepr{
		flatCoords: append(p1.flatCoords, p2.flatCoords...), endss: append(p1.endss, p2.endss...),
	}
}

//line wkt.y:95
type wktSymType struct {
	yys               int
	str               string
	geom              geom.T
	coord             float64
	coordList         []float64
	flatRepr          geomFlatCoordsRepr
	multiPolyFlatRepr multiPolygonFlatCoordsRepr
	geomList          []geom.T
	geomCollect       *geom.GeometryCollection
}

const POINT = 57346
const POINTM = 57347
const POINTZ = 57348
const POINTZM = 57349
const LINESTRING = 57350
const LINESTRINGM = 57351
const LINESTRINGZ = 57352
const LINESTRINGZM = 57353
const POLYGON = 57354
const POLYGONM = 57355
const POLYGONZ = 57356
const POLYGONZM = 57357
const MULTIPOINT = 57358
const MULTIPOINTM = 57359
const MULTIPOINTZ = 57360
const MULTIPOINTZM = 57361
const MULTILINESTRING = 57362
const MULTILINESTRINGM = 57363
const MULTILINESTRINGZ = 57364
const MULTILINESTRINGZM = 57365
const MULTIPOLYGON = 57366
const MULTIPOLYGONM = 57367
const MULTIPOLYGONZ = 57368
const MULTIPOLYGONZM = 57369
const GEOMETRYCOLLECTION = 57370
const GEOMETRYCOLLECTIONM = 57371
const GEOMETRYCOLLECTIONZ = 57372
const GEOMETRYCOLLECTIONZM = 57373
const EMPTY = 57374
const NUM = 57375

var wktToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"POINT",
	"POINTM",
	"POINTZ",
	"POINTZM",
	"LINESTRING",
	"LINESTRINGM",
	"LINESTRINGZ",
	"LINESTRINGZM",
	"POLYGON",
	"POLYGONM",
	"POLYGONZ",
	"POLYGONZM",
	"MULTIPOINT",
	"MULTIPOINTM",
	"MULTIPOINTZ",
	"MULTIPOINTZM",
	"MULTILINESTRING",
	"MULTILINESTRINGM",
	"MULTILINESTRINGZ",
	"MULTILINESTRINGZM",
	"MULTIPOLYGON",
	"MULTIPOLYGONM",
	"MULTIPOLYGONZ",
	"MULTIPOLYGONZM",
	"GEOMETRYCOLLECTION",
	"GEOMETRYCOLLECTIONM",
	"GEOMETRYCOLLECTIONZ",
	"GEOMETRYCOLLECTIONZM",
	"EMPTY",
	"NUM",
	"'('",
	"')'",
	"','",
}

var wktStatenames = [...]string{}

const wktEofCode = 1
const wktErrCode = 2
const wktInitialStackSize = 16

//line yacctab:1
var wktExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const wktPrivate = 57344

const wktLast = 218

var wktAct = [...]int{
	66, 2, 116, 126, 131, 106, 65, 111, 128, 121,
	60, 57, 137, 118, 155, 156, 61, 70, 58, 101,
	75, 114, 81, 63, 87, 69, 63, 64, 104, 63,
	61, 63, 90, 63, 62, 63, 107, 68, 63, 61,
	72, 84, 77, 61, 83, 78, 89, 108, 59, 94,
	58, 67, 153, 154, 71, 92, 74, 97, 80, 61,
	86, 151, 152, 93, 149, 150, 147, 148, 145, 146,
	143, 144, 141, 142, 139, 140, 57, 102, 138, 57,
	61, 97, 58, 27, 113, 109, 26, 25, 70, 63,
	24, 70, 23, 22, 136, 63, 56, 21, 123, 133,
	20, 63, 19, 18, 17, 95, 124, 16, 15, 14,
	13, 12, 134, 11, 99, 10, 1, 91, 119, 135,
	88, 85, 130, 125, 129, 132, 127, 82, 79, 120,
	115, 122, 117, 76, 73, 110, 103, 112, 105, 100,
	98, 96, 9, 8, 7, 57, 6, 57, 5, 102,
	4, 161, 113, 70, 160, 70, 163, 63, 165, 164,
	162, 63, 158, 133, 123, 63, 3, 0, 0, 0,
	0, 0, 124, 159, 0, 0, 134, 0, 0, 0,
	0, 0, 0, 0, 0, 119, 0, 0, 157, 129,
	28, 29, 30, 31, 32, 33, 34, 35, 36, 37,
	38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
	48, 49, 50, 51, 52, 53, 54, 55,
}

var wktPact = [...]int{
	186, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	16, 27, 27, 16, 27, 27, 16, 27, 27, -16,
	11, -16, 7, -16, -2, 21, 27, 27, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 24, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 24, -1000, -1000, -1000,
	16, -1000, -1000, -1000, -1000, 48, -1000, -1000, 48, -1000,
	-1000, -16, -1000, -1000, -16, -1000, -1000, -16, -1000, -1000,
	-16, -1000, 186, -1000, -1000, -23, 45, -1000, 39, -1000,
	37, -1000, -1000, 35, -1000, -1000, -1000, -1000, -1000, -1000,
	33, -1000, -1000, -1000, -1000, 31, -1000, -1000, -1000, -1000,
	29, -1000, -1000, -1000, -1000, 26, -1000, -1000, -1000, -1000,
	17, -1000, -1000, -1000, -1000, -21, -1000, -1000, -1000, -1000,
	24, -1000, 16, -1000, 48, -1000, 48, -1000, -16, -1000,
	-16, -1000, -16, -1000, -16, -1000, 186, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000,
}

var wktPgo = [...]int{
	0, 1, 166, 150, 148, 146, 144, 143, 142, 36,
	21, 10, 141, 47, 85, 140, 6, 13, 19, 139,
	8, 5, 138, 137, 28, 7, 136, 135, 134, 133,
	132, 131, 2, 9, 130, 129, 128, 127, 126, 125,
	3, 4, 123, 122, 121, 120, 119, 117, 116, 115,
	113, 111, 110, 109, 108, 107, 104, 103, 102, 100,
	97, 93, 92, 90, 87, 86, 83, 0,
}

var wktR1 = [...]int{
	0, 48, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 2, 49, 49, 50, 51, 51, 51, 3, 3,
	3, 52, 52, 53, 54, 54, 54, 4, 4, 4,
	55, 55, 56, 57, 57, 57, 5, 5, 5, 5,
	58, 59, 59, 59, 6, 6, 6, 6, 60, 61,
	61, 61, 7, 7, 7, 7, 62, 63, 63, 63,
	8, 8, 8, 47, 46, 46, 64, 64, 65, 66,
	66, 66, 67, 44, 45, 43, 43, 42, 42, 40,
	41, 38, 38, 39, 39, 36, 37, 34, 34, 35,
	35, 32, 33, 30, 30, 31, 31, 28, 29, 26,
	26, 27, 27, 24, 25, 22, 22, 23, 23, 21,
	21, 20, 19, 19, 18, 17, 16, 15, 15, 14,
	13, 12, 12, 9, 10, 11,
}

var wktR2 = [...]int{
	0, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 2, 1, 1, 1, 1, 1, 1, 2, 2,
	2, 1, 1, 1, 1, 1, 1, 2, 2, 2,
	1, 1, 1, 1, 1, 1, 2, 2, 2, 2,
	1, 1, 1, 1, 2, 2, 2, 2, 1, 1,
	1, 1, 2, 2, 2, 2, 1, 1, 1, 1,
	2, 2, 2, 3, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 3, 3, 1, 3, 1, 1,
	1, 1, 1, 1, 1, 3, 3, 3, 1, 3,
	1, 1, 1, 1, 1, 1, 1, 3, 3, 3,
	1, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 3, 1, 1, 1, 3, 3, 1, 3,
	1, 2, 1, 1, 1, 1,
}

var wktChk = [...]int{
	-1000, -48, -1, -2, -3, -4, -5, -6, -7, -8,
	-49, -50, -51, -52, -53, -54, -55, -56, -57, -58,
	-59, -60, -61, -62, -63, -64, -65, -66, 4, 5,
	6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, -14, -67, 34, -9,
	-11, 32, -10, -11, -17, -16, -67, -9, -10, -20,
	-67, -9, -10, -28, -9, -67, -29, -10, 34, -36,
	-9, -67, -37, -10, 34, -44, -9, -67, -45, -10,
	34, -47, 34, -9, -10, -13, -12, 33, -15, -13,
	-19, -18, -16, -26, -24, -22, -21, -9, -13, -14,
	-27, -25, -23, -21, -10, -34, -32, -30, -17, -9,
	-35, -33, -31, -17, -10, -42, -40, -38, -20, -9,
	-43, -41, -39, -20, -10, -46, -1, 35, 33, 35,
	36, 35, 36, 35, 36, 35, 36, 35, 36, 35,
	36, 35, 36, 35, 36, 35, 36, -13, -18, -24,
	-25, -32, -33, -40, -41, -1,
}

var wktDef = [...]int{
	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 12, 13, 0, 21, 22, 0, 30, 31, 0,
	0, 0, 0, 0, 0, 0, 66, 67, 14, 15,
	16, 17, 23, 24, 25, 26, 32, 33, 34, 35,
	40, 41, 42, 43, 48, 49, 50, 51, 56, 57,
	58, 59, 68, 69, 70, 71, 9, 0, 72, 10,
	123, 125, 11, 124, 18, 115, 0, 19, 20, 27,
	0, 28, 29, 36, 38, 0, 37, 39, 0, 44,
	46, 0, 45, 47, 0, 52, 54, 0, 53, 55,
	0, 60, 0, 61, 62, 0, 120, 122, 0, 118,
	0, 113, 114, 0, 100, 103, 105, 106, 109, 110,
	0, 102, 104, 107, 108, 0, 88, 91, 93, 94,
	0, 90, 92, 95, 96, 0, 78, 79, 81, 82,
	0, 76, 80, 83, 84, 0, 65, 119, 121, 116,
	0, 111, 0, 97, 0, 98, 0, 85, 0, 86,
	0, 73, 0, 74, 0, 63, 0, 117, 112, 99,
	101, 87, 89, 77, 75, 64,
}

var wktTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	34, 35, 3, 3, 36,
}

var wktTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33,
}

var wktTok3 = [...]int{
	0,
}

var wktErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	wktDebug        = 0
	wktErrorVerbose = true
)

type wktLexer interface {
	Lex(lval *wktSymType) int
	Error(s string)
}

type wktParser interface {
	Parse(wktLexer) int
	Lookahead() int
}

type wktParserImpl struct {
	lval  wktSymType
	stack [wktInitialStackSize]wktSymType
	char  int
}

func (p *wktParserImpl) Lookahead() int {
	return p.char
}

func wktNewParser() wktParser {
	return &wktParserImpl{}
}

const wktFlag = -1000

func wktTokname(c int) string {
	if c >= 1 && c-1 < len(wktToknames) {
		if wktToknames[c-1] != "" {
			return wktToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func wktStatname(s int) string {
	if s >= 0 && s < len(wktStatenames) {
		if wktStatenames[s] != "" {
			return wktStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func wktErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !wktErrorVerbose {
		return "syntax error"
	}

	for _, e := range wktErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + wktTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := wktPact[state]
	for tok := TOKSTART; tok-1 < len(wktToknames); tok++ {
		if n := base + tok; n >= 0 && n < wktLast && wktChk[wktAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if wktDef[state] == -2 {
		i := 0
		for wktExca[i] != -1 || wktExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; wktExca[i] >= 0; i += 2 {
			tok := wktExca[i]
			if tok < TOKSTART || wktExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if wktExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += wktTokname(tok)
	}
	return res
}

func wktlex1(lex wktLexer, lval *wktSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = wktTok1[0]
		goto out
	}
	if char < len(wktTok1) {
		token = wktTok1[char]
		goto out
	}
	if char >= wktPrivate {
		if char < wktPrivate+len(wktTok2) {
			token = wktTok2[char-wktPrivate]
			goto out
		}
	}
	for i := 0; i < len(wktTok3); i += 2 {
		token = wktTok3[i+0]
		if token == char {
			token = wktTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = wktTok2[1] /* unknown char */
	}
	if wktDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", wktTokname(token), uint(char))
	}
	return char, token
}

func wktParse(wktlex wktLexer) int {
	return wktNewParser().Parse(wktlex)
}

func (wktrcvr *wktParserImpl) Parse(wktlex wktLexer) int {
	var wktn int
	var wktVAL wktSymType
	var wktDollar []wktSymType
	_ = wktDollar // silence set and not used
	wktS := wktrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	wktstate := 0
	wktrcvr.char = -1
	wkttoken := -1 // wktrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		wktstate = -1
		wktrcvr.char = -1
		wkttoken = -1
	}()
	wktp := -1
	goto wktstack

ret0:
	return 0

ret1:
	return 1

wktstack:
	/* put a state and value onto the stack */
	if wktDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", wktTokname(wkttoken), wktStatname(wktstate))
	}

	wktp++
	if wktp >= len(wktS) {
		nyys := make([]wktSymType, len(wktS)*2)
		copy(nyys, wktS)
		wktS = nyys
	}
	wktS[wktp] = wktVAL
	wktS[wktp].yys = wktstate

wktnewstate:
	wktn = wktPact[wktstate]
	if wktn <= wktFlag {
		goto wktdefault /* simple state */
	}
	if wktrcvr.char < 0 {
		wktrcvr.char, wkttoken = wktlex1(wktlex, &wktrcvr.lval)
	}
	wktn += wkttoken
	if wktn < 0 || wktn >= wktLast {
		goto wktdefault
	}
	wktn = wktAct[wktn]
	if wktChk[wktn] == wkttoken { /* valid shift */
		wktrcvr.char = -1
		wkttoken = -1
		wktVAL = wktrcvr.lval
		wktstate = wktn
		if Errflag > 0 {
			Errflag--
		}
		goto wktstack
	}

wktdefault:
	/* default state action */
	wktn = wktDef[wktstate]
	if wktn == -2 {
		if wktrcvr.char < 0 {
			wktrcvr.char, wkttoken = wktlex1(wktlex, &wktrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if wktExca[xi+0] == -1 && wktExca[xi+1] == wktstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			wktn = wktExca[xi+0]
			if wktn < 0 || wktn == wkttoken {
				break
			}
		}
		wktn = wktExca[xi+1]
		if wktn < 0 {
			goto ret0
		}
	}
	if wktn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			wktlex.Error(wktErrorMessage(wktstate, wkttoken))
			Nerrs++
			if wktDebug >= 1 {
				__yyfmt__.Printf("%s", wktStatname(wktstate))
				__yyfmt__.Printf(" saw %s\n", wktTokname(wkttoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for wktp >= 0 {
				wktn = wktPact[wktS[wktp].yys] + wktErrCode
				if wktn >= 0 && wktn < wktLast {
					wktstate = wktAct[wktn] /* simulate a shift of "error" */
					if wktChk[wktstate] == wktErrCode {
						goto wktstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if wktDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", wktS[wktp].yys)
				}
				wktp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if wktDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", wktTokname(wkttoken))
			}
			if wkttoken == wktEofCode {
				goto ret1
			}
			wktrcvr.char = -1
			wkttoken = -1
			goto wktnewstate /* try again in the same state */
		}
	}

	/* reduction by production wktn */
	if wktDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", wktn, wktStatname(wktstate))
	}

	wktnt := wktn
	wktpt := wktp
	_ = wktpt // guard against "declared and not used"

	wktp -= wktR2[wktn]
	// wktp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if wktp+1 >= len(wktS) {
		nyys := make([]wktSymType, len(wktS)*2)
		copy(nyys, wktS)
		wktS = nyys
	}
	wktVAL = wktS[wktp+1]

	/* consult goto table to find next state */
	wktn = wktR1[wktn]
	wktg := wktPgo[wktn]
	wktj := wktg + wktS[wktp].yys + 1

	if wktj >= wktLast {
		wktstate = wktAct[wktg]
	} else {
		wktstate = wktAct[wktj]
		if wktChk[wktstate] != -wktn {
			wktstate = wktAct[wktg]
		}
	}
	// dummy call; replaced with literal code
	switch wktnt {

	case 1:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:184
		{
			ok := wktlex.(*wktLex).validateLayoutStackAtEnd()
			if !ok {
				return 1
			}
			wktlex.(*wktLex).ret = wktDollar[1].geom
		}
	case 8:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:200
		{
			ok := wktlex.(*wktLex).validateAndPopLayoutStackFrame()
			if !ok {
				return 1
			}
			err := wktDollar[1].geomCollect.SetLayout(wktlex.(*wktLex).curLayout())
			if err != nil {
				wktlex.(*wktLex).setError(err)
				return 1
			}
			wktVAL.geom = wktDollar[1].geomCollect
		}
	case 9:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:215
		{
			wktVAL.geom = geom.NewPointFlat(wktlex.(*wktLex).curLayout(), wktDollar[2].coordList)
		}
	case 10:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:219
		{
			wktVAL.geom = geom.NewPointEmpty(wktlex.(*wktLex).curLayout())
		}
	case 11:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:223
		{
			wktVAL.geom = geom.NewPointEmpty(wktlex.(*wktLex).curLayout())
		}
	case 14:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:233
		{
			ok := wktlex.(*wktLex).validateBaseGeometryTypeAllowed()
			if !ok {
				return 1
			}
		}
	case 15:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:242
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYM)
			if !ok {
				return 1
			}
		}
	case 16:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:249
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYZ)
			if !ok {
				return 1
			}
		}
	case 17:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:256
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYZM)
			if !ok {
				return 1
			}
		}
	case 18:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:265
		{
			wktVAL.geom = geom.NewLineStringFlat(wktlex.(*wktLex).curLayout(), wktDollar[2].coordList)
		}
	case 19:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:269
		{
			wktVAL.geom = geom.NewLineString(wktlex.(*wktLex).curLayout())
		}
	case 20:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:273
		{
			wktVAL.geom = geom.NewLineString(wktlex.(*wktLex).curLayout())
		}
	case 23:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:283
		{
			ok := wktlex.(*wktLex).validateBaseGeometryTypeAllowed()
			if !ok {
				return 1
			}
		}
	case 24:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:292
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYM)
			if !ok {
				return 1
			}
		}
	case 25:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:299
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYZ)
			if !ok {
				return 1
			}
		}
	case 26:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:306
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYZM)
			if !ok {
				return 1
			}
		}
	case 27:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:315
		{
			wktVAL.geom = geom.NewPolygonFlat(wktlex.(*wktLex).curLayout(), wktDollar[2].flatRepr.flatCoords, wktDollar[2].flatRepr.ends)
		}
	case 28:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:319
		{
			wktVAL.geom = geom.NewPolygon(wktlex.(*wktLex).curLayout())
		}
	case 29:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:323
		{
			wktVAL.geom = geom.NewPolygon(wktlex.(*wktLex).curLayout())
		}
	case 32:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:333
		{
			ok := wktlex.(*wktLex).validateBaseGeometryTypeAllowed()
			if !ok {
				return 1
			}
		}
	case 33:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:342
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYM)
			if !ok {
				return 1
			}
		}
	case 34:

		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:349
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYZ)
			if !ok {
				return 1
			}
		}
	case 35:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:356
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYZM)
			if !ok {
				return 1
			}
		}
	case 36:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:365
		{
			wktVAL.geom = geom.NewMultiPointFlat(
				wktlex.(*wktLex).curLayout(), wktDollar[2].flatRepr.flatCoords, geom.NewMultiPointFlatOptionWithEnds(wktDollar[2].flatRepr.ends),
			)
		}
	case 37:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:371
		{
			wktVAL.geom = geom.NewMultiPointFlat(
				wktlex.(*wktLex).curLayout(), wktDollar[2].flatRepr.flatCoords, geom.NewMultiPointFlatOptionWithEnds(wktDollar[2].flatRepr.ends),
			)
		}
	case 38:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:377
		{
			wktVAL.geom = geom.NewMultiPoint(wktlex.(*wktLex).curLayout())
		}
	case 39:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:381
		{
			wktVAL.geom = geom.NewMultiPoint(wktlex.(*wktLex).curLayout())
		}
	case 40:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:387
		{
			ok := wktlex.(*wktLex).validateBaseGeometryTypeAllowed()
			if !ok {
				return 1
			}
		}
	case 41:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:396
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYM)
			if !ok {
				return 1
			}
		}
	case 42:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:403
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYZ)
			if !ok {
				return 1
			}
		}
	case 43:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:410
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYZM)
			if !ok {
				return 1
			}
		}
	case 44:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:419
		{
			wktVAL.geom = geom.NewMultiLineStringFlat(wktlex.(*wktLex).curLayout(), wktDollar[2].flatRepr.flatCoords, wktDollar[2].flatRepr.ends)
		}
	case 45:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:423
		{
			wktVAL.geom = geom.NewMultiLineStringFlat(wktlex.(*wktLex).curLayout(), wktDollar[2].flatRepr.flatCoords, wktDollar[2].flatRepr.ends)
		}
	case 46:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:427
		{
			wktVAL.geom = geom.NewMultiLineString(wktlex.(*wktLex).curLayout())
		}
	case 47:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:431
		{
			wktVAL.geom = geom.NewMultiLineString(wktlex.(*wktLex).curLayout())
		}
	case 48:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:437
		{
			ok := wktlex.(*wktLex).validateBaseGeometryTypeAllowed()
			if !ok {
				return 1
			}
		}
	case 49:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:446
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYM)
			if !ok {
				return 1
			}
		}
	case 50:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:453
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYZ)
			if !ok {
				return 1
			}
		}
	case 51:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:460
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYZM)
			if !ok {
				return 1
			}
		}
	case 52:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:469
		{
			wktVAL.geom = geom.NewMultiPolygonFlat(wktlex.(*wktLex).curLayout(), wktDollar[2].multiPolyFlatRepr.flatCoords, wktDollar[2].multiPolyFlatRepr.endss)
		}
	case 53:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:473
		{
			wktVAL.geom = geom.NewMultiPolygonFlat(wktlex.(*wktLex).curLayout(), wktDollar[2].multiPolyFlatRepr.flatCoords, wktDollar[2].multiPolyFlatRepr.endss)
		}
	case 54:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:477
		{
			wktVAL.geom = geom.NewMultiPolygon(wktlex.(*wktLex).curLayout())
		}
	case 55:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:481
		{
			wktVAL.geom = geom.NewMultiPolygon(wktlex.(*wktLex).curLayout())
		}
	case 56:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:487
		{
			ok := wktlex.(*wktLex).validateBaseGeometryTypeAllowed()
			if !ok {
				return 1
			}
		}
	case 57:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:496
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYM)
			if !ok {
				return 1
			}
		}
	case 58:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:503
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYZ)
			if !ok {
				return 1
			}
		}
	case 59:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:510
		{
			ok := wktlex.(*wktLex).validateAndSetLayoutIfNoLayout(geom.XYZM)
			if !ok {
				return 1
			}
		}
	case 60:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:519
		{
			newCollection := geom.NewGeometryCollection()
			err := newCollection.Push(wktDollar[2].geomList...)
			if err != nil {
				wktlex.(*wktLex).setError(err)
				return 1
			}
			wktVAL.geomCollect = newCollection
		}
	case 61:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:529
		{
			wktVAL.geomCollect = geom.NewGeometryCollection()
		}
	case 62:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:533
		{
			wktVAL.geomCollect = geom.NewGeometryCollection()
		}
	case 63:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:539
		{
			wktVAL.geomList = wktDollar[2].geomList
		}
	case 64:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:545
		{
			wktVAL.geomList = append(wktDollar[1].geomList, wktDollar[3].geom)
		}
	case 65:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:549
		{
			wktVAL.geomList = []geom.T{wktDollar[1].geom}
		}
	case 68:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:559
		{
			ok := wktlex.(*wktLex).validateAndPushLayoutStackFrame(geom.NoLayout)
			if !ok {
				return 1
			}
		}
	case 69:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:568
		{
			ok := wktlex.(*wktLex).validateAndPushLayoutStackFrame(geom.XYM)
			if !ok {
				return 1
			}
		}
	case 70:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:575
		{
			ok := wktlex.(*wktLex).validateAndPushLayoutStackFrame(geom.XYZ)
			if !ok {
				return 1
			}
		}
	case 71:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:582
		{
			ok := wktlex.(*wktLex).validateAndPushLayoutStackFrame(geom.XYZM)
			if !ok {
				return 1
			}
		}
	case 72:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:591
		{
			ok := wktlex.(*wktLex).validateNonEmptyGeometryAllowed()
			if !ok {
				return 1
			}
		}
	case 73:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:600
		{
			wktVAL.multiPolyFlatRepr = wktDollar[2].multiPolyFlatRepr
		}
	case 74:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:606
		{
			wktVAL.multiPolyFlatRepr = wktDollar[2].multiPolyFlatRepr
		}
	case 75:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:612
		{
			wktVAL.multiPolyFlatRepr = appendMultiPolygonFlatCoordsRepr(wktDollar[1].multiPolyFlatRepr, wktDollar[3].multiPolyFlatRepr)
		}
	case 77:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:619
		{
			wktVAL.multiPolyFlatRepr = appendMultiPolygonFlatCoordsRepr(wktDollar[1].multiPolyFlatRepr, wktDollar[3].multiPolyFlatRepr)
		}
	case 79:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:626
		{
			wktVAL.multiPolyFlatRepr = makeMultiPolygonFlatCoordsRepr(wktDollar[1].flatRepr)
		}
	case 80:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:632
		{
			wktVAL.multiPolyFlatRepr = makeMultiPolygonFlatCoordsRepr(wktDollar[1].flatRepr)
		}
	case 82:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:639
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 84:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:646
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 85:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:652
		{
			wktVAL.flatRepr = wktDollar[2].flatRepr
		}
	case 86:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:658
		{
			wktVAL.flatRepr = wktDollar[2].flatRepr
		}
	case 87:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:664
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 89:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:671
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 91:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:678
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 92:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:684
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 97:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:698
		{
			wktVAL.flatRepr = wktDollar[2].flatRepr
		}
	case 98:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:704
		{
			wktVAL.flatRepr = wktDollar[2].flatRepr
		}
	case 99:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:710
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 101:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:717
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 103:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:724
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 104:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:730
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 111:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:747
		{
			wktVAL.flatRepr = wktDollar[2].flatRepr
		}
	case 112:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:753
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 114:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:760
		{
			if !isValidPolygonRing(wktlex, wktDollar[1].coordList, wktlex.(*wktLex).curLayout().Stride()) {
				return 1
			}
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 115:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:769
		{
			if !isValidLineString(wktlex, wktDollar[1].coordList, wktlex.(*wktLex).curLayout().Stride()) {
				return 1
			}
		}
	case 116:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:777
		{
			wktVAL.coordList = wktDollar[2].coordList
		}
	case 117:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:783
		{
			wktVAL.coordList = append(wktDollar[1].coordList, wktDollar[3].coordList...)
		}
	case 119:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:790
		{
			wktVAL.coordList = wktDollar[2].coordList
		}
	case 120:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:796
		{
			switch len(wktDollar[1].coordList) {
			case 1:
				wktlex.(*wktLex).setParseError("not enough coordinates", "each point needs at least 2 coords")
				return 1
			case 2, 3, 4:
				ok := wktlex.(*wktLex).validateStrideAndSetDefaultLayoutIfNoLayout(len(wktDollar[1].coordList))
				if !ok {
					return 1
				}
			default:
				wktlex.(*wktLex).setParseError("too many coordinates", "each point can have at most 4 coords")
				return 1
			}
		}
	case 121:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:814
		{
			wktVAL.coordList = append(wktDollar[1].coordList, wktDollar[2].coord)
		}
	case 122:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:818
		{
			wktVAL.coordList = []float64{wktDollar[1].coord}
		}
	case 123:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:824
		{
			ok := wktlex.(*wktLex).validateBaseTypeEmptyAllowed()
			if !ok {
				return 1
			}
		}
	case 125:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:836
		{
			wktVAL.coordList = []float64(nil)
		}
	}
	goto wktstack /* stack new state and value */
}
