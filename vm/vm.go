package vm

import (
	"yolk/types"
	"yolk/utils"
)

type VirtualMachine struct {
	instructions        []Instruction
	instruction_pointer int
	stack               utils.Stack[types.Primitive]
	output_buffer       utils.Queue[string]
}
