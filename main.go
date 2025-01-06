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

	if len(*runFlag) == 0 {
		log.Fatal("--run flag not spcified")
	}

	if err := cli.ExecuteYolkFile(&machine, *runFlag, *debugFlag); err != nil {
		log.Fatal(err)
	}

	if len(*profilerFlag) != 0 {
		pprof.StopCPUProfile()
	}
}
