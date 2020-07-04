package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"io"
	"flag"
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


func Menu(status int) int {
	var fullHelp bool = (status == 0)
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

func evalLineByLine() {
	return
}

/*
MAIN main mAiN
*/
func main() {
	fmt.Println("Welcome to Gython")
	cliFlag := flag.Bool("cli", false, "Run Python CLI prompt.")
	pythonFileFlag := flag.String("file", "", "Python *.py file to run.")
	flag.Parse()

	if *cliFlag == true {
		evalLineByLine()
		// eventually (hopefully) this should just do the same thing as the file
		// i.e. treat each new line as being appended to the "file" so we can define a function
		// with multiple lines or something instead of each line being fully atomic
	} else if (true) { // if file exists(*pythonFileFlag)
		// check that file is a .py file (and first that it exists)
		// read contents to this variable and build AST from that
		var fileContents, err = ioutil.ReadFile("")
		if err != nil {
			fmt.Println("Error reading file ", *pythonFileFlag, ". Exiting.")
			os.Exit(2)
		}
		fmt.Println(fileContents)
	}

}
