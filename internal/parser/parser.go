package parser

import (
	"os"
)

var MAXSTACK int = 1700
var MAXLEVEL int = 200
var MAXINDENT int = 100

type perrdetail struct {
    err int
    filename *PyObject
    lineno int
    offset int
    text string                 /* UTF-8-encoded string */
    token int
    expected int
}

type stackentry struct {
	s_state int
	s_dfa *dfa
	s_parent *node
}

type stack struct {
	s_top *stackentry
	s_base []stackentry
}

type parser_state struct {
	p_stack stack
	p_grammar *grammar
	p_tree *node
	// p_flags int
}

type items struct {
	lineno int
	comment string
}

type growable_comment_array struct {
	items_list []items
	size int
	num_items int
}

/* Tokenizer state */
type tok_state struct {
    /* Input state; buf <= cur <= inp <= end */
    /* NB an entire line is held in the buffer */
    buf string          /* Input buffer, or NULL; malloc'ed if fp != NULL */
    cur string          /* Next character in buffer */
    inp string          /* End of data in buffer */
    end string          /* End of input buffer if buf != NULL */
    start string        /* Start of current token if not NULL */
    done int           /* E_OK normally, E_EOF at EOF, otherwise error code */
    /* NB If done != E_OK, cur must be == inp!!! */
    fp *os.File           /* Rest of input; NULL if tokenizing a string */
    tabsize int        /* Tab spacing */
    indent int         /* Current indentation index */
    indstack []int            /* Stack of indents */
    atbol int          /* Nonzero if at begin of new line */
    pendin int         /* Pending indents (if > 0) or dedents (if < 0) */
	prompt string
	nextprompt string          /* For interactive prompting */
    lineno int         /* Current line number */
    first_lineno int   /* First line of a single line or multi line string
                           expression (cf. issue 16806) */
    level int          /* () [] {} Parentheses nesting level */
            /* Used to allow free continuations inside them */
    parenstack string
    parenlinenostack []int
    filename *PyObject
    /* Stuff for checking on different tab sizes */
    altindstack []int         /* Stack of alternate indents */
    /* Stuff for PEP 0263 */
    //enum decoding_state decoding_state;
    decoding_erred int         /* whether erred in decoding  */
    read_coding_spec int       /* whether 'coding:...' has been read  */
    encoding string         /* Source encoding. */
    cont_line int          /* whether we are in a continuation line. */
    line_start string     /* pointer to start of current line */
    multi_line_start string /* pointer to start of first line of
                                     a single line or multi line string
                                     expression (cf. issue 16806) */
    decoding_readline *PyObject /* open(...).readline */
    decoding_buffer *PyObject
    enc string        /* Encoding for the current str. */
    str string
    input string /* Tokenizer's newline translated copy of the string. */

    type_comments int      /* Whether to look for type comments */

    /* async/await related fields (still needed depending on feature_version) */
    async_hacks int     /* =1 if async/await aren't always keywords */
    async_def int        /* =1 if tokens are inside an 'async def' body. */
    async_def_indent int /* Indentation level of the outermost 'async def'. */
    async_def_nl int     /* =1 if the outermost 'async def' had at least one
                             NEWLINE token after it. */
};

func growable_comment_array_init(arr *growable_comment_array, initial_size int) int {
	var listitems []items
	arr.items_list = listitems
	return arr.size
}

func growable_comment_array_add(arr *growable_comment_array, lineno int, comment string) int {
	if arr.num_items >= arr.size {
		arr.size *= 2
		if arr.items_list == nil {
			return 0
		}
	}
	arr.items_list[arr.num_items].lineno = lineno
	arr.items_list[arr.num_items].comment = comment
	return 1
}

func parsetok(tok *tok_state, g *grammar, start int, err_ret *perrdetail, flags []int) *node {

}
