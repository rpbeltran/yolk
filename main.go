package main

import (
	"flag"
	"log"
	"yolk/cli"
	"yolk/vm"
)

var runFlag = flag.String("run", "", ".yolk file to exec")
var debugFlag = flag.Bool("debug", false, "Show vm state after each instruction")

func main() {
	flag.Parse()
	machine := vm.NewVM()
	machine.MockExecutions = true

	if len(*runFlag) == 0 {
		log.Fatal("--run flag not spcified")
	}

	if err := cli.ExecuteYolkFile(&machine, *runFlag, *debugFlag); err != nil {
		log.Fatal(err)
	}
}
