package benchmarks

import (
	"bytes"
	"testing"
	"yolk/vm"
)

func BenchmarkPushNum(b *testing.B) {
	// push on an off the stack b.N times
	push, err := vm.ParseInstruction(`PUSH_NUM 0`)
	if err != nil {
		b.Fatalf("Unexpected error parsing push instruction: %v", err)
	}

	machine := vm.NewVM()
	for n := 0; n < b.N; n++ {
		if err := push.Perform(&machine); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkPushStr(b *testing.B) {
	// push on an off the stack b.N times
	push, err := vm.ParseInstruction(`PUSH_STR "hello"`)
	if err != nil {
		b.Fatalf("Unexpected error parsing push instruction: %v", err)
	}

	machine := vm.NewVM()
	for n := 0; n < b.N; n++ {
		if err := push.Perform(&machine); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkPushPop(b *testing.B) {
	// push on an off the stack b.N times
	var output_buffer bytes.Buffer

	push, err := vm.ParseInstruction(`PUSH_NUM 0`)
	if err != nil {
		b.Fatalf("Unexpected error parsing push instruction: %v", err)
	}

	printer, err := vm.ParseInstruction(`PRINT`)
	if err != nil {
		b.Fatalf("Unexpected error parsing push instruction: %v", err)
	}

	machine := vm.NewDebugVM(&output_buffer)
	for n := 0; n < b.N; n++ {
		if err := push.Perform(&machine); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
		if err := printer.Perform(&machine); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
		if err := push.Perform(&machine); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
		if err := printer.Perform(&machine); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
		if err := push.Perform(&machine); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
		if err := printer.Perform(&machine); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
		if err := push.Perform(&machine); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
		if err := printer.Perform(&machine); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}
