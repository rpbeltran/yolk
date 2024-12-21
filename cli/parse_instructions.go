package cli

import (
	"bufio"
	"fmt"
	"yolk/vm"
)

func put_program_in_vm(machine *vm.VirtualMachine, scanner bufio.Scanner) error {
	machine.ClearProgram()

	line_num := 0
	for scanner.Scan() {
		line_num += 1
		line := scanner.Text()
		if instruction, err := vm.ParseInstruction(line); err != nil {
			return fmt.Errorf("parsing line %d %q: %w", line_num, line, err)
		} else if instruction != nil {
			machine.AddProgramInstruction(instruction)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanning yolk: %v", err)
	}
	return nil
}
