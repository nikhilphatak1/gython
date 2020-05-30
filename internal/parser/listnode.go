package parser

import (
	"os"
	"fmt"
	"io"
)

type node struct {
	n_type int
	n_str string
	n_lineno int
	n_col_offset int
	n_nchildren int
	n_child []*node
	n_end_lineno int
	n_end_col_offset int
}

var level int
var atbol int
var MAX_INT32 int = 1<<31 - 1

func NCH(n *node) int { return n.n_nchildren}
func CHILD(n *node, i int) *node { return n.n_child[i]}
func RCHILD(n *node, i int) *node {
	return CHILD(n, NCH(n) + i)
}
func TYPE(n *node) int { return n.n_type}
func STR(n *node) string { return n.n_str}
func LINENO(n *node) int { return n.n_lineno}

func PyNode_ListTree(n *node) {
	listnode(os.Stdout, n)
}

func listnode(w io.Writer, n *node) {
	level = 0
	atbol = 1
	list1node(w, n)
}

func list1node(w io.Writer, n *node) {
	if n == nil {
		return
	}
	if TYPE(n) >= 256 { // >= NT_OFFSET
		var i int
		for i = 0; i < NCH(n); i++ {
			list1node(w, CHILD(n, i))
		}
	} else if TYPE(n) < 256 {
		switch TYPE(n) {
		case 5: level++
		case 6: level--
		default:
			if atbol == 0 {
				var i int
				for i = 0; i < level; i++ {
					fmt.Fprint(w, "\t")
				}
				atbol = 0
			}
			if TYPE(n) == 4 { //NEWLINE
				if STR(n) != "" {
					fmt.Fprintf(w, "%s", STR(n))
				}
				fmt.Fprint(w, "\n")
				atbol = 1
			} else {
				fmt.Fprintf(w, "%s ", STR(n))
			}
		}
	} else {
		fmt.Fprint(w, "? ")
	}
}

func PyNode_New(n_type int) *node {
	var n *node
	n.n_type = n_type
	n.n_str = ""
	n.n_lineno = 0
	n.n_end_lineno = 0
	n.n_end_col_offset = -1
	n.n_nchildren = 0
	n.n_child = nil
	return n
}

func fancy_roundup(n int) int {
	var result int = 256
	for result < n {
		result <<= 1
		if result <= 0 {
			return -1
		}
	}
	return result
}

func _Py_SIZE_ROUND_UP(n, a int) int {
	return (n + a - 1) & ^(a - 1)
}

func XXXROUNDUP(n int) int {
	if n <= 1 {
		return 1
	} else {
		if n <= 128 {
			return _Py_SIZE_ROUND_UP(n, 4)
		} else {
			return fancy_roundup(n)
		}
	}
}

func _PyNode_FinalizeEndPos(n *node) {
	var nch int = NCH(n)
	var last *node
	if nch == 0 {
		return
	}
	last = CHILD(n, nch - 1)
	_PyNode_FinalizeEndPos(last)
	n.n_end_lineno = last.n_end_lineno
	n.n_end_col_offset = last.n_end_col_offset
}

func PyNode_AddChild(n1 *node, n_type int, str string,
					 lineno int, col_offset int, end_lineno int,
					 enc_col_offset int) int {
	var nch int = n1.n_nchildren
	var current_capacity int
	var required_capacity int
	var n *node

	if nch > 0 {
		_PyNode_FinalizeEndPos(CHILD(n1, nch - 1))
	}

	if nch == MAX_INT32 || nch < 0 {
		return 19 // E_OVERFLOW
	}
	if current_capacity < required_capacity {
		/*
		if required_capacity > SIZE_MAX / sizeof(node) {
			return E_NOMEM
		}
		*/
		n = n1.n_child[0] // n = n1->n_child beginning of array
	}
	n = n1.n_child[n1.n_nchildren]
	n1.n_nchildren++
	n.n_type = n_type
	n.n_str = str
	n.n_lineno = lineno
	n.n_col_offset = col_offset
	n.n_end_lineno = end_lineno
	n.n_end_col_offset = enc_col_offset
	n.n_nchildren = 0
	n.n_child = nil
	return 0
}
