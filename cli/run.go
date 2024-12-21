package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"yolk/vm"
)

func ExecuteYolkFile(machine *vm.VirtualMachine, path string, verbose_debug bool) error {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := put_program_in_vm(machine, *scanner); err != nil {
		return fmt.Errorf("parsing instructions: %v", err)
	}

	if err := machine.RunProgram(verbose_debug); err != nil {
		return fmt.Errorf("executing program: %v", err)
	}

	return nil
}
