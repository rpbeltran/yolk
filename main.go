package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"yolk/instructions"
	"yolk/vm"
)

var runFlag = flag.String("run", "", ".yolk file to exec")

func execute_yolk(path string) error {
	var machine vm.VirtualMachine

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if instruction, err := instructions.ParseInstruction(scanner.Text()); err != nil {
			return err
		} else if instruction != nil {
			if err := instruction.Perform(&machine); err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func main() {
	flag.Parse()

	if len(*runFlag) == 0 {
		log.Fatal("--run flag not spcified")
	}

	if err := execute_yolk(*runFlag); err != nil {
		log.Fatal(err)
	}
}
