package main

import (
	"fmt"
	"os"
	"io"

	"gython/internal/util"
)

// Status enum
type Status int

const (
	OK Status = iota
	ERROR
	NOT_RUN
	NO_FILE
)

// Action enum
type Action int

const (
	RUN Action = iota
	ERROR_OUT
	HELP
	VERSION
)


type CommandLineOptions struct {
	Action Action
	Message string
	Command string
	Filename string
	Module string
	Help bool
	Version bool
	Jar bool
	Verbosity int
	// TODO Properties (basically a map)
	Argv []string
	WarnOptions []string
}

func NewCommandLineOptions() CommandLineOptions {
	c := CommandLineOptions{
		Action: RUN,
		Message: "",
		Help: false,
		Version: false,
		Jar: false,
		Verbosity: 0,
	}
	return c
}

func Parse() CommandLineOptions {
	var clo CommandLineOptions
	return clo
}


func Menu(status Status) Status {
	var fullHelp bool = (status == OK)
	var f io.Writer
	if fullHelp {
		f = os.Stdout
	} else {
		f = os.Stderr
	}
	fmt.Fprint(f, util.HowToUseGython)
	if fullHelp {
		fmt.Fprint(f, "~use `gython -h` for help~")
	}
	return status
}


/*
MAIN main mAiN
*/
func main() {
	fmt.Println("Welcome to Gython")
	var opts CommandLineOptions = Parse()

	switch opts.Action {
	case VERSION :
		fmt.Fprint(os.Stderr, "Gython", version.PY_VERSION)
		os.Exit(OK)
	case HELP:
		os.Exit(Menu(OK))
	case ERROR_OUT:
		fmt.Fprint(os.Stderr, opts.Message)
	case RUN:
	}

	


}
