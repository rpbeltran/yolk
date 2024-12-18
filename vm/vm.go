package vm

import (
	"yolk/types"
	"yolk/utils"
)

type VirtualMachine struct {
	instructions        []Instruction
	instruction_pointer int
	stack               utils.Stack[types.Primitive] //todo: benchmark this being a pointer
	output_buffer       utils.Queue[string]
	pipeline_states     utils.Stack[*types.Primitive]
}
