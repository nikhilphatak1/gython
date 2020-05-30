package parser

var _PyParser_TokenNames = []string{
	"ENDMARKER",
    "NAME",
    "NUMBER",
    "STRING",
    "NEWLINE",
    "INDENT",
    "DEDENT",
    "LPAR",
    "RPAR",
    "LSQB",
    "RSQB",
    "COLON",
    "COMMA",
    "SEMI",
    "PLUS",
    "MINUS",
    "STAR",
    "SLASH",
    "VBAR",
    "AMPER",
    "LESS",
    "GREATER",
    "EQUAL",
    "DOT",
    "PERCENT",
    "LBRACE",
    "RBRACE",
    "EQEQUAL",
    "NOTEQUAL",
    "LESSEQUAL",
    "GREATEREQUAL",
    "TILDE",
    "CIRCUMFLEX",
    "LEFTSHIFT",
    "RIGHTSHIFT",
    "DOUBLESTAR",
    "PLUSEQUAL",
    "MINEQUAL",
    "STAREQUAL",
    "SLASHEQUAL",
    "PERCENTEQUAL",
    "AMPEREQUAL",
    "VBAREQUAL",
    "CIRCUMFLEXEQUAL",
    "LEFTSHIFTEQUAL",
    "RIGHTSHIFTEQUAL",
    "DOUBLESTAREQUAL",
    "DOUBLESLASH",
    "DOUBLESLASHEQUAL",
    "AT",
    "ATEQUAL",
    "RARROW",
    "ELLIPSIS",
    "COLONEQUAL",
    "OP",
    "AWAIT",
    "ASYNC",
    "TYPE_IGNORE",
    "TYPE_COMMENT",
    "<ERRORTOKEN>",
    "<COMMENT>",
    "<NL>",
    "<ENCODING>",
    "<N_TOKENS>",
}

func PyToken_OneChar(c1 int) int {
	switch c1 {
		case '%': return 24 //PERCENT;
		case '&': return 19 //AMPER;
		case '(': return 7 //LPAR;
		case ')': return 8 //RPAR;
		case '*': return 16 //STAR;
		case '+': return 14 //PLUS;
		case ',': return 12 //COMMA;
		case '-': return 15 //MINUS;
		case '.': return 23 //DOT;
		case '/': return 17 //SLASH;
		case ':': return 11 //COLON;
		case ';': return 13 //SEMI;
		case '<': return 20 //LESS;
		case '=': return 22 //EQUAL;
		case '>': return 21 //GREATER;
		case '@': return 49 //AT;
		case '[': return 9 //LSQB;
		case ']': return 10 //RSQB;
		case '^': return 32 //CIRCUMFLEX;
		case '{': return 25 //LBRACE;
		case '|': return 18 //VBAR;
		case '}': return 26 //RBRACE;
		case '~': return 31 //TILDE;
	default: return 54 //OP;
	}
}

func PyToken_TwoChars(c1, c2 int) int {
	switch c1 {
    case '!':
        switch c2 {
        case '=': return 28 //NOTEQUAL;
        }
    case '%':
        switch c2 {
        case '=': return 40 //PERCENTEQUAL;
        }
    case '&':
        switch c2 {
        case '=': return 41 //AMPEREQUAL;
        }
    case '*':
        switch c2 {
        case '*': return 35 //DOUBLESTAR;
        case '=': return 38 //STAREQUAL;
        }
    case '+':
        switch c2 {
        case '=': return 36 //PLUSEQUAL;
        }
    case '-':
        switch c2 {
        case '=': return 37 //MINEQUAL;
        case '>': return 51 //RARROW;
        }
    case '/':
        switch c2 {
        case '/': return 47 //DOUBLESLASH;
        case '=': return 39 //SLASHEQUAL;
        }
    case ':':
        switch c2 {
        case '=': return 53 //COLONEQUAL;
        }
    case '<':
        switch c2 {
        case '<': return 33 //LEFTSHIFT;
        case '=': return 29 //LESSEQUAL;
        case '>': return 28 //NOTEQUAL;
        }
    case '=':
        switch c2 {
        case '=': return 27 //EQEQUAL;
        }
    case '>':
        switch c2 {
        case '=': return 30 //GREATEREQUAL;
        case '>': return 34 //RIGHTSHIFT;
        }
    case '@':
        switch c2 {
        case '=': return 50 //ATEQUAL;
        }
    case '^':
        switch c2 {
        case '=': return 32 //CIRCUMFLEXEQUAL;
        }
    case '|':
        switch c2 {
        case '=': return 42 //VBAREQUAL;
        }
	}
	return 54 //OP;
}

func PyToken_ThreeChars(c1, c2, c3 int) int {
	switch c1 {
    case '*':
        switch c2 {
        case '*':
            switch c3 {
            case '=': return 46 //DOUBLESTAREQUAL;
            }
        }
    case '.':
        switch c2 {
        case '.':
            switch c3 {
            case '.': return 52 //ELLIPSIS;
            }
        }
    case '/':
        switch c2 {
        case '/':
            switch c3 {
            case '=': return 48 //DOUBLESLASHEQUAL;
            }
        }
    case '<':
        switch c2 {
        case '<':
            switch c3 {
            case '=': return 44 //LEFTSHIFTEQUAL;
            }
        }
    case '>':
        switch c2 {
        case '>':
            switch c3 {
            case '=': return 45 //RIGHTSHIFTEQUAL;
            }
        }
    }
	return 54 //OP;
}
