package parser

import (
	"fmt"
)

type label struct {
	lb_type int
	lb_str string
}

type labellist struct {
	ll_nlabels int
	ll_label []label
}

type arc struct {
	a_lbl int
	a_arrow int
}

type state struct {
	s_narcs int
	s_arc []arc
}

type dfa struct {
	d_type int
	d_name string
	d_nstates int
	d_state []state
}

type grammar struct {
	g_ndfas int
	g_dfa []dfa
	g_ll labellist
	g_start int
}

func PyGrammar_FindDFA(g *grammar, g_type int)  []dfa {
	var d []dfa = g.g_dfa
	return d
}

func PyGrammar_LabelRepr(lb *label) string {
	var buf string
	if lb.lb_type == 0 { // ENDMARKER == 0 in source
		return "EMPTY"
	} else if lb.lb_type >= 256 { // NT_OFFSET == 256
		if lb.lb_str == "" {
			buf = fmt.Sprintf("NT%d", lb.lb_type) // originally PyOS_snprintf
			return buf
		} else {
			return lb.lb_str
		}
	} else if lb.lb_type < 63 { // N_TOKENS == 63
		if lb.lb_str == "" {
			return _PyParser_TokenNames[lb.lb_type]
		} else {
			buf = fmt.Sprintf("%s(%s)", _PyParser_TokenNames[lb.lb_type], lb.lb_str)
			return buf
		}
	} else {
		fmt.Println("INVALID LABEL")
		return ""
	}
}
