package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"yolk/vm"
)

var runFlag = flag.String("run", "", ".yolk file to exec")

func execute_yolk(path string, verbose_debug bool) error {
	var machine vm.VirtualMachine
	machine.MockExecutions = true

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line_num := 0
	for scanner.Scan() {
		line_num += 1
		line := scanner.Text()
		if instruction, err := vm.ParseInstruction(line); err != nil {
			return fmt.Errorf("parsing line %d %q: %w", line_num, line, err)
		} else if instruction != nil {
			if verbose_debug {
				fmt.Printf("-- executing %d: %v\n", line_num, instruction)
			}
			if err := instruction.Perform(&machine); err != nil {
				return fmt.Errorf("executing line %d %q: %w", line_num, line, err)
			}
			if verbose_debug {
				dbg_state := machine.GetDebugState()
				fmt.Printf("   -- stack size: %d\n", dbg_state.StackSize)
				if dbg_state.StackSize != 0 {
					fmt.Printf("   -- top of stack: %q\n", dbg_state.TopOfStack)
				}
			}
		} else if verbose_debug {
			fmt.Printf("-- skipping  %d: %q\n", line_num, line)
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

	if err := execute_yolk(*runFlag, true); err != nil {
		log.Fatal(err)
	}
}
