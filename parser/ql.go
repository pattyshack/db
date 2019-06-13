// Code generated by goyacc -o ql.go -v ql.output -p ql ql.y. DO NOT EDIT.

//line ql.y:2
package parser

import __yyfmt__ "fmt"

//line ql.y:2

//line ql.y:5
type qlSymType struct {
	yys int
	*Token

	ControlFlowExpr
	Expr

	Statements []ControlFlowExpr
}

const LEX_ERROR = 57346
const BLOCK_COMMENT_END = 57347
const COMMENT = 57348
const L_BRACE = 57349
const R_BRACE = 57350
const L_PAREN = 57351
const R_PAREN = 57352
const L_BRACKET = 57353
const R_BRACKET = 57354
const ASSIGN = 57355
const COLON = 57356
const COMMA = 57357
const DOT = 57358
const STAR_STAR = 57359
const SEMICOLON = 57360
const NEWLINE = 57361
const LET = 57362
const IF = 57363
const ELSE = 57364
const RETURN = 57365
const OR = 57366
const AND = 57367
const NOT = 57368
const LT = 57369
const GT = 57370
const EQ = 57371
const NE = 57372
const LE = 57373
const GE = 57374
const BITWISE_OR = 57375
const BITWISE_AND = 57376
const XOR = 57377
const L_SHIFT = 57378
const R_SHIFT = 57379
const ADD = 57380
const SUB = 57381
const MUL = 57382
const DIV = 57383
const MOD = 57384
const UNARY = 57385
const IDENT = 57386
const CHAR = 57387
const STRING = 57388
const INT = 57389
const FLOAT = 57390
const BOOL = 57391
const NOOP = 57392

var qlToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"BLOCK_COMMENT_END",
	"COMMENT",
	"L_BRACE",
	"R_BRACE",
	"L_PAREN",
	"R_PAREN",
	"L_BRACKET",
	"R_BRACKET",
	"ASSIGN",
	"COLON",
	"COMMA",
	"DOT",
	"STAR_STAR",
	"SEMICOLON",
	"NEWLINE",
	"LET",
	"IF",
	"ELSE",
	"RETURN",
	"OR",
	"AND",
	"NOT",
	"LT",
	"GT",
	"EQ",
	"NE",
	"LE",
	"GE",
	"BITWISE_OR",
	"BITWISE_AND",
	"XOR",
	"L_SHIFT",
	"R_SHIFT",
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"MOD",
	"UNARY",
	"IDENT",
	"CHAR",
	"STRING",
	"INT",
	"FLOAT",
	"BOOL",
	"NOOP",
}
var qlStatenames = [...]string{}

const qlEofCode = 1
const qlErrCode = 2
const qlInitialStackSize = 16

//line ql.y:607

//line yacctab:1
var qlExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const qlPrivate = 57344

const qlLast = 259

var qlAct = [...]int{

	35, 84, 34, 11, 48, 49, 50, 51, 52, 53,
	54, 55, 56, 57, 33, 92, 21, 55, 56, 57,
	64, 36, 7, 8, 60, 61, 51, 52, 53, 54,
	55, 56, 57, 62, 53, 54, 55, 56, 57, 94,
	63, 66, 67, 68, 69, 70, 71, 72, 73, 74,
	75, 76, 77, 78, 79, 80, 81, 82, 83, 14,
	87, 31, 10, 59, 2, 10, 89, 36, 90, 93,
	58, 37, 14, 15, 5, 86, 10, 14, 85, 20,
	4, 19, 14, 38, 18, 1, 13, 12, 7, 8,
	16, 15, 6, 17, 39, 97, 95, 3, 22, 40,
	41, 65, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 88, 0,
	9, 0, 0, 0, 0, 0, 91, 0, 0, 0,
	0, 0, 40, 41, 0, 42, 43, 44, 45, 46,
	47, 48, 49, 50, 51, 52, 53, 54, 55, 56,
	57, 0, 40, 41, 96, 42, 43, 44, 45, 46,
	47, 48, 49, 50, 51, 52, 53, 54, 55, 56,
	57, 14, 0, 32, 50, 51, 52, 53, 54, 55,
	56, 57, 14, 0, 32, 15, 0, 0, 0, 0,
	24, 49, 50, 51, 52, 53, 54, 55, 56, 57,
	0, 24, 0, 23, 0, 0, 0, 0, 25, 26,
	27, 28, 29, 30, 23, 0, 0, 0, 0, 25,
	26, 27, 28, 29, 30, 41, 0, 42, 43, 44,
	45, 46, 47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 51, 52, 53, 54, 55, 56, 57,
}
var qlPact = [...]int{

	70, -1000, -1000, 70, -1000, -1000, 4, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 70, 175, -30, 164, -1000, -1000,
	63, 75, 54, 175, 175, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 175, 27, -1000, 128, -1000, -1000, -2, 65,
	175, 175, 175, 175, 175, 175, 175, 175, 175, 175,
	175, 175, 175, 175, 175, 175, 175, 175, -43, 175,
	-1000, 216, 108, 164, 52, -7, 200, 216, -29, -29,
	-29, -29, -29, -29, 157, 139, -10, -4, -4, -23,
	-23, -1000, -1000, -1000, -1000, 59, 24, 128, -1000, -1000,
	-1000, -1000, 52, -1000, 175, -1000, -1000, 128,
}
var qlPgo = [...]int{

	0, 2, 98, 0, 64, 97, 80, 92, 61, 3,
	87, 86, 85, 74, 78, 75,
}
var qlR1 = [...]int{

	0, 12, 4, 4, 5, 5, 13, 13, 6, 6,
	7, 7, 7, 7, 7, 8, 9, 9, 9, 9,
	9, 9, 10, 11, 1, 1, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 14, 14, 15,
	15,
}
var qlR2 = [...]int{

	0, 1, 0, 1, 1, 2, 1, 1, 1, 2,
	1, 1, 1, 1, 1, 3, 3, 4, 5, 6,
	5, 6, 4, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 4, 3, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 0, 1, 1,
	3,
}
var qlChk = [...]int{

	-1000, -12, -4, -5, -6, -13, -7, 18, 19, 50,
	-8, -9, -10, -11, 7, 21, 20, 23, -6, -13,
	-4, -3, -2, 39, 26, 44, 45, 46, 47, 48,
	49, -8, 9, 44, -1, -3, -9, 8, -8, 19,
	24, 25, 27, 28, 29, 30, 31, 32, 33, 34,
	35, 36, 37, 38, 39, 40, 41, 42, 16, 9,
	-3, -3, -3, 13, 22, -8, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, -3, -3, 44, -14, -15, -3, 10, -1,
	-9, -8, 22, 10, 15, -9, -8, -3,
}
var qlDef = [...]int{

	2, -2, 1, 3, 4, 8, 0, 6, 7, 10,
	11, 12, 13, 14, 2, 0, 0, 0, 5, 9,
	0, 0, 36, 0, 0, 26, 27, 28, 29, 30,
	31, 32, 0, 0, 23, 24, 25, 15, 16, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 57,
	55, 56, 0, 0, 0, 17, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 53, 54, 33, 0, 58, 59, 35, 22,
	18, 20, 0, 34, 0, 19, 21, 60,
}
var qlTok1 = [...]int{

	1,
}
var qlTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50,
}
var qlTok3 = [...]int{
	0,
}

var qlErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	qlDebug        = 0
	qlErrorVerbose = false
)

type qlLexer interface {
	Lex(lval *qlSymType) int
	Error(s string)
}

type qlParser interface {
	Parse(qlLexer) int
	Lookahead() int
}

type qlParserImpl struct {
	lval  qlSymType
	stack [qlInitialStackSize]qlSymType
	char  int
}

func (p *qlParserImpl) Lookahead() int {
	return p.char
}

func qlNewParser() qlParser {
	return &qlParserImpl{}
}

const qlFlag = -1000

func qlTokname(c int) string {
	if c >= 1 && c-1 < len(qlToknames) {
		if qlToknames[c-1] != "" {
			return qlToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func qlStatname(s int) string {
	if s >= 0 && s < len(qlStatenames) {
		if qlStatenames[s] != "" {
			return qlStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func qlErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !qlErrorVerbose {
		return "syntax error"
	}

	for _, e := range qlErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + qlTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := qlPact[state]
	for tok := TOKSTART; tok-1 < len(qlToknames); tok++ {
		if n := base + tok; n >= 0 && n < qlLast && qlChk[qlAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if qlDef[state] == -2 {
		i := 0
		for qlExca[i] != -1 || qlExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; qlExca[i] >= 0; i += 2 {
			tok := qlExca[i]
			if tok < TOKSTART || qlExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if qlExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += qlTokname(tok)
	}
	return res
}

func qllex1(lex qlLexer, lval *qlSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = qlTok1[0]
		goto out
	}
	if char < len(qlTok1) {
		token = qlTok1[char]
		goto out
	}
	if char >= qlPrivate {
		if char < qlPrivate+len(qlTok2) {
			token = qlTok2[char-qlPrivate]
			goto out
		}
	}
	for i := 0; i < len(qlTok3); i += 2 {
		token = qlTok3[i+0]
		if token == char {
			token = qlTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = qlTok2[1] /* unknown char */
	}
	if qlDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", qlTokname(token), uint(char))
	}
	return char, token
}

func qlParse(qllex qlLexer) int {
	return qlNewParser().Parse(qllex)
}

func (qlrcvr *qlParserImpl) Parse(qllex qlLexer) int {
	var qln int
	var qlVAL qlSymType
	var qlDollar []qlSymType
	_ = qlDollar // silence set and not used
	qlS := qlrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	qlstate := 0
	qlrcvr.char = -1
	qltoken := -1 // qlrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		qlstate = -1
		qlrcvr.char = -1
		qltoken = -1
	}()
	qlp := -1
	goto qlstack

ret0:
	return 0

ret1:
	return 1

qlstack:
	/* put a state and value onto the stack */
	if qlDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", qlTokname(qltoken), qlStatname(qlstate))
	}

	qlp++
	if qlp >= len(qlS) {
		nyys := make([]qlSymType, len(qlS)*2)
		copy(nyys, qlS)
		qlS = nyys
	}
	qlS[qlp] = qlVAL
	qlS[qlp].yys = qlstate

qlnewstate:
	qln = qlPact[qlstate]
	if qln <= qlFlag {
		goto qldefault /* simple state */
	}
	if qlrcvr.char < 0 {
		qlrcvr.char, qltoken = qllex1(qllex, &qlrcvr.lval)
	}
	qln += qltoken
	if qln < 0 || qln >= qlLast {
		goto qldefault
	}
	qln = qlAct[qln]
	if qlChk[qln] == qltoken { /* valid shift */
		qlrcvr.char = -1
		qltoken = -1
		qlVAL = qlrcvr.lval
		qlstate = qln
		if Errflag > 0 {
			Errflag--
		}
		goto qlstack
	}

qldefault:
	/* default state action */
	qln = qlDef[qlstate]
	if qln == -2 {
		if qlrcvr.char < 0 {
			qlrcvr.char, qltoken = qllex1(qllex, &qlrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if qlExca[xi+0] == -1 && qlExca[xi+1] == qlstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			qln = qlExca[xi+0]
			if qln < 0 || qln == qltoken {
				break
			}
		}
		qln = qlExca[xi+1]
		if qln < 0 {
			goto ret0
		}
	}
	if qln == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			qllex.Error(qlErrorMessage(qlstate, qltoken))
			Nerrs++
			if qlDebug >= 1 {
				__yyfmt__.Printf("%s", qlStatname(qlstate))
				__yyfmt__.Printf(" saw %s\n", qlTokname(qltoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for qlp >= 0 {
				qln = qlPact[qlS[qlp].yys] + qlErrCode
				if qln >= 0 && qln < qlLast {
					qlstate = qlAct[qln] /* simulate a shift of "error" */
					if qlChk[qlstate] == qlErrCode {
						goto qlstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if qlDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", qlS[qlp].yys)
				}
				qlp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if qlDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", qlTokname(qltoken))
			}
			if qltoken == qlEofCode {
				goto ret1
			}
			qlrcvr.char = -1
			qltoken = -1
			goto qlnewstate /* try again in the same state */
		}
	}

	/* reduction by production qln */
	if qlDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", qln, qlStatname(qlstate))
	}

	qlnt := qln
	qlpt := qlp
	_ = qlpt // guard against "declared and not used"

	qlp -= qlR2[qln]
	// qlp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if qlp+1 >= len(qlS) {
		nyys := make([]qlSymType, len(qlS)*2)
		copy(nyys, qlS)
		qlS = nyys
	}
	qlVAL = qlS[qlp+1]

	/* consult goto table to find next state */
	qln = qlR1[qln]
	qlg := qlPgo[qln]
	qlj := qlg + qlS[qlp].yys + 1

	if qlj >= qlLast {
		qlstate = qlAct[qlg]
	} else {
		qlstate = qlAct[qlj]
		if qlChk[qlstate] != -qln {
			qlstate = qlAct[qlg]
		}
	}
	// dummy call; replaced with literal code
	switch qlnt {

	case 1:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:74
		{
			nodes := make([]Node, 0, len(qlDollar[1].Statements))
			for _, node := range qlDollar[1].Statements {
				nodes = append(nodes, node)
			}
			qllex.(*parseContext).setParsed(nodes)
		}
	case 2:
		qlDollar = qlS[qlpt-0 : qlpt+1]
//line ql.y:84
		{
			qlVAL.Statements = nil
		}
	case 3:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:88
		{
			qlVAL.Statements = qlDollar[1].Statements
		}
	case 4:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:94
		{
			if qlDollar[1].ControlFlowExpr != nil {
				qlVAL.Statements = append(qlVAL.Statements, qlDollar[1].ControlFlowExpr)
			}
		}
	case 5:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:99
		{
			if qlDollar[2].ControlFlowExpr != nil {
				qlVAL.Statements = append(qlDollar[1].Statements, qlDollar[2].ControlFlowExpr)
			}
		}
	case 8:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:112
		{
			qlVAL.ControlFlowExpr = nil
		}
	case 9:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:115
		{
		}
	case 10:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:120
		{
			qlVAL.ControlFlowExpr = &Noop{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 11:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:126
		{
			qlVAL.ControlFlowExpr = qlDollar[1].ControlFlowExpr
		}
	case 12:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:129
		{
			qlVAL.ControlFlowExpr = qlDollar[1].ControlFlowExpr
		}
	case 13:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:132
		{
			qlVAL.ControlFlowExpr = qlDollar[1].ControlFlowExpr
		}
	case 14:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:135
		{
			qlVAL.ControlFlowExpr = qlDollar[1].ControlFlowExpr
		}
	case 15:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:142
		{
			qlVAL.ControlFlowExpr = &ExprBlock{
				Location: Location{
					Filename: qlDollar[1].Token.Loc().Filename,
					Start:    qlDollar[1].Token.Loc().Start,
					End:      qlDollar[3].Token.Loc().End,
				},
				LBrace:     qlDollar[1].Token,
				Statements: qlDollar[2].Statements,
				RBrace:     qlDollar[3].Token,
			}
		}
	case 16:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:157
		{
			qlVAL.ControlFlowExpr = &ConditionalExpr{
				Location: Location{
					Filename: qlDollar[1].Token.Loc().Filename,
					Start:    qlDollar[1].Token.Loc().Start,
					End:      qlDollar[3].ControlFlowExpr.Loc().End,
				},
				If:         qlDollar[1].Token,
				Predicate:  qlDollar[2].Expr,
				TrueClause: qlDollar[3].ControlFlowExpr,
			}
		}
	case 17:
		qlDollar = qlS[qlpt-4 : qlpt+1]
//line ql.y:169
		{
			qlVAL.ControlFlowExpr = &ConditionalExpr{
				Location: Location{
					Filename: qlDollar[1].Token.Loc().Filename,
					Start:    qlDollar[1].Token.Loc().Start,
					End:      qlDollar[4].ControlFlowExpr.Loc().End,
				},
				If:         qlDollar[1].Token,
				Predicate:  qlDollar[2].Expr,
				TrueClause: qlDollar[4].ControlFlowExpr,
			}
		}
	case 18:
		qlDollar = qlS[qlpt-5 : qlpt+1]
//line ql.y:181
		{
			qlVAL.ControlFlowExpr = &ConditionalExpr{
				Location: Location{
					Filename: qlDollar[1].Token.Loc().Filename,
					Start:    qlDollar[1].Token.Loc().Start,
					End:      qlDollar[5].ControlFlowExpr.Loc().End,
				},
				If:          qlDollar[1].Token,
				Predicate:   qlDollar[2].Expr,
				TrueClause:  qlDollar[3].ControlFlowExpr,
				Else:        qlDollar[4].Token,
				FalseClause: qlDollar[5].ControlFlowExpr,
			}
		}
	case 19:
		qlDollar = qlS[qlpt-6 : qlpt+1]
//line ql.y:195
		{
			qlVAL.ControlFlowExpr = &ConditionalExpr{
				Location: Location{
					Filename: qlDollar[1].Token.Loc().Filename,
					Start:    qlDollar[1].Token.Loc().Start,
					End:      qlDollar[6].ControlFlowExpr.Loc().End,
				},
				If:          qlDollar[1].Token,
				Predicate:   qlDollar[2].Expr,
				TrueClause:  qlDollar[4].ControlFlowExpr,
				Else:        qlDollar[5].Token,
				FalseClause: qlDollar[6].ControlFlowExpr,
			}
		}
	case 20:
		qlDollar = qlS[qlpt-5 : qlpt+1]
//line ql.y:209
		{
			qlVAL.ControlFlowExpr = &ConditionalExpr{
				Location: Location{
					Filename: qlDollar[1].Token.Loc().Filename,
					Start:    qlDollar[1].Token.Loc().Start,
					End:      qlDollar[5].ControlFlowExpr.Loc().End,
				},
				If:          qlDollar[1].Token,
				Predicate:   qlDollar[2].Expr,
				TrueClause:  qlDollar[3].ControlFlowExpr,
				Else:        qlDollar[4].Token,
				FalseClause: qlDollar[5].ControlFlowExpr,
			}
		}
	case 21:
		qlDollar = qlS[qlpt-6 : qlpt+1]
//line ql.y:223
		{
			qlVAL.ControlFlowExpr = &ConditionalExpr{
				Location: Location{
					Filename: qlDollar[1].Token.Loc().Filename,
					Start:    qlDollar[1].Token.Loc().Start,
					End:      qlDollar[6].ControlFlowExpr.Loc().End,
				},
				If:          qlDollar[1].Token,
				Predicate:   qlDollar[2].Expr,
				TrueClause:  qlDollar[4].ControlFlowExpr,
				Else:        qlDollar[5].Token,
				FalseClause: qlDollar[6].ControlFlowExpr,
			}
		}
	case 22:
		qlDollar = qlS[qlpt-4 : qlpt+1]
//line ql.y:240
		{
			qlVAL.ControlFlowExpr = &AssignExpr{
				Location: Location{
					Filename: qlDollar[1].Token.Loc().Filename,
					Start:    qlDollar[1].Token.Loc().Start,
					End:      qlDollar[4].Expr.Loc().End,
				},
				Let:        qlDollar[1].Token,
				Name:       qlDollar[2].Token,
				Assign:     qlDollar[3].Token,
				Expression: qlDollar[4].Expr,
			}
		}
	case 23:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:257
		{
			qlVAL.ControlFlowExpr = &ReturnExpr{
				Location: Location{
					Filename: qlDollar[1].Token.Loc().Filename,
					Start:    qlDollar[1].Token.Loc().Start,
					End:      qlDollar[2].Expr.Loc().End,
				},
				Return:     qlDollar[1].Token,
				Expression: qlDollar[2].Expr,
			}
		}
	case 24:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:271
		{
			qlVAL.Expr = qlDollar[1].Expr
		}
	case 25:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:274
		{
			qlVAL.Expr = qlDollar[1].ControlFlowExpr
		}
	case 26:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:281
		{
			qlVAL.Expr = &Identifier{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 27:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:287
		{
			qlVAL.Expr = &Literal{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 28:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:293
		{
			qlVAL.Expr = &Literal{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 29:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:299
		{
			qlVAL.Expr = &Literal{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 30:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:305
		{
			qlVAL.Expr = &Literal{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 31:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:311
		{
			qlVAL.Expr = &Literal{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 32:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:317
		{
			qlVAL.Expr = qlDollar[1].ControlFlowExpr
		}
	case 33:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:320
		{
			qlVAL.Expr = &Accessor{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Token.Loc().End,
				},
				PrimaryExpr: qlDollar[1].Expr,
				Dot:         qlDollar[2].Token,
				Name:        qlDollar[3].Token,
			}
		}
	case 34:
		qlDollar = qlS[qlpt-4 : qlpt+1]
//line ql.y:332
		{
		}
	case 35:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:334
		{
			qlVAL.Expr = &EvalOrderExpr{
				Location: Location{
					Filename: qlDollar[1].Token.Loc().Filename,
					Start:    qlDollar[1].Token.Loc().Start,
					End:      qlDollar[3].Token.Loc().End,
				},
				LParen:     qlDollar[1].Token,
				Expression: qlDollar[2].Expr,
				RParen:     qlDollar[3].Token,
			}
		}
	case 36:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:350
		{
			qlVAL.Expr = qlDollar[1].Expr
		}
	case 37:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:353
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 38:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:365
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 39:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:377
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 40:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:389
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 41:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:401
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 42:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:413
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 43:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:425
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 44:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:437
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 45:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:449
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 46:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:461
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 47:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:473
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 48:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:485
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 49:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:497
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 50:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:509
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 51:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:521
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 52:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:533
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 53:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:545
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 54:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:557
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 55:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:569
		{
			qlVAL.Expr = &UnaryExpr{
				Location: Location{
					Filename: qlDollar[1].Token.Loc().Filename,
					Start:    qlDollar[1].Token.Loc().Start,
					End:      qlDollar[2].Expr.Loc().End,
				},
				Op:         qlDollar[1].Token,
				Expression: qlDollar[2].Expr,
			}
		}
	case 56:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:580
		{
			qlVAL.Expr = &UnaryExpr{
				Location: Location{
					Filename: qlDollar[1].Token.Loc().Filename,
					Start:    qlDollar[1].Token.Loc().Start,
					End:      qlDollar[2].Expr.Loc().End,
				},
				Op:         qlDollar[1].Token,
				Expression: qlDollar[2].Expr,
			}
		}
	case 57:
		qlDollar = qlS[qlpt-0 : qlpt+1]
//line ql.y:594
		{ // empty
		}
	case 58:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:596
		{
		}
	case 59:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:601
		{
		}
	case 60:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:603
		{
		}
	}
	goto qlstack /* stack new state and value */
}
