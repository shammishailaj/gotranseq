package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/feliixx/gotranseq/transeq"
	"github.com/jessevdk/go-flags"
)

const (
	version  = "0.1"
	toolName = "gotranseq"
)

func printErrorAndExit(err error) {
	fmt.Printf("error: %v\n", err)
	os.Exit(1)
}

func main() {

	var options transeq.Options
	p := flags.NewParser(&options, flags.Default&^flags.HelpFlag)
	_, err := p.Parse()
	if err != nil {
		fmt.Printf("wrong args: %v, try %s --help for more informations\n", err, toolName)
		os.Exit(1)
	}
	if options.Help {
		fmt.Printf("%s version %s\n\n", toolName, version)
		p.WriteHelp(os.Stdout)
		os.Exit(0)
	}
	if options.Version {
		fmt.Printf("%s version version %s\n", toolName, version)
		os.Exit(0)
	}
	if options.Sequence == "" {
		printErrorAndExit(fmt.Errorf("missing required parameter -s | -sequence, try %s --help for details", toolName))
	}
	if options.Outseq == "" {
		printErrorAndExit(fmt.Errorf("missing required parameter -o | -outseq, try %s --help for details", toolName))
	}

	if options.NumWorker == 0 {
		options.NumWorker = runtime.NumCPU()
	}
	in, err := os.Open(options.Sequence)
	if err != nil {
		printErrorAndExit(fmt.Errorf("Could not read from input file %v: %v", options.Sequence, err))
	}
	defer in.Close()
	out, err := os.Create(options.Outseq)
	if err != nil {
		printErrorAndExit(fmt.Errorf("Could not write to output file file %v: %v", options.Outseq, err))
	}
	defer out.Close()

	ioHandler := transeq.IOHandler{
		In:  in,
		Out: out,
	}

	err = ioHandler.ReadSequenceAndTranslate(options)
	if err != nil {
		printErrorAndExit(err)
	}
}
