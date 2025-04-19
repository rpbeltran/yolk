package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"yolk/cli"
	"yolk/vm"
)

var runFlag = flag.String("run", "", ".yolk file to exec")
var interactiveFlag = flag.Bool("interactive", false, "execute instructions interactively via stdin.")

var profilerFlag = flag.String("profiler", "", "start profiler and write the profile to the given path")
var debugFlag = flag.Bool("debug", false, "Show vm state after each instruction")

func main() {
	flag.Parse()

	if len(*profilerFlag) != 0 {
		if f, err := os.Create(*profilerFlag); err != nil {
			log.Fatalf("starting profiler: %v", err)
		} else {
			pprof.StartCPUProfile(f)
		}
	}

	machine := vm.NewVM()
	machine.MockExecutions = true

	if *interactiveFlag {
		if err := cli.ExecuteInteractive(&machine, *debugFlag); err != nil {
			log.Fatal(err)
		}
	} else if len(*runFlag) == 0 {
		log.Fatal("either --interactive or --run must be specified")
	} else if err := cli.ExecuteYolkFile(&machine, *runFlag, *debugFlag); err != nil {
		log.Fatal(err)
	}

	if len(*profilerFlag) != 0 {
		pprof.StopCPUProfile()
	}
}
