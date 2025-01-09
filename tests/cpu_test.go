package cpu_test

import "snowtime/cpu"
import "testing"

func TestADCInstructionImmediate(t *testing.T) {
	_6502 := cpu.NewCPU()
	_6502_memory := _6502.GetMemory()

	var target_value uint8 = 0x01

	_6502_memory.Write(0, 0x69)         // Write ADC
	_6502_memory.Write(1, target_value) // Write the Immdeiate value
	_6502_memory.Write(2, 0x69)         // Write ADC
	_6502_memory.Write(3, 0x01)         // Write the Immediate Value

	_6502.Tick()

	if _6502.GetAcc() != target_value {
		t.Fatalf(`FAILED: Accumulator should be set to %d. Found: %d`,
			target_value,
			_6502.GetAcc())
	}

	_6502.Tick()

	if _6502.GetAcc() != target_value+0x01 {
		t.Fatalf(`FAILED: Accumulator should be set to %d. Found: %d`,
			target_value+1,
			_6502.GetAcc())
	}
}

func TestADCInstructionZeroPage(t *testing.T) {
	_6502 := cpu.NewCPU()
	_6502_memory := _6502.GetMemory()

	// Add the value stored at 0x00 into acc
	_6502_memory.Write(0, 0x65)
	_6502.Tick()

	if _6502.GetAcc() != 0x65 {
		t.Fatalf(`FAILED: Zero Paging did not work as expected`)
	}
}
