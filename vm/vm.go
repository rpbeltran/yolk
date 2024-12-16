package vm

import "yolk/types"

type VirtualMachine struct {
	instructions        []Instruction
	instruction_pointer int
	stack               Stack[types.Primitive]
}
