package cpu

// Handles the logic of the CPU. Stores the acc, x,y, flg, registers along with
// the program counter and stack pointer.
// For more details, kindly read
//
// https://www.nesdev.org/6502.txt
//

type CPU struct {
	// The general purpose registers
	acc uint8
	x   uint8
	y   uint8

	// An 8-bit unsigned integer that contains the NES's various
	// The 6502 uses 8 bits to represent the following flags
	//
	//  Bit No.     7  6  5  4  3  2  1  0
	//              S  V     B  D  I  Z  C
	// * C - Carry Flag : the c_out of the MSB in any arithmetic operation.
	//                    this flag will be cleared for subtraction operations
	// * Z - Zero Flag  : This is set to 1 when any arithmetic or logical operation
	//                    produces a zero result, is 0 if its non-zero
	// * I - Interrupt  : If set, then interrupts are disabled. Else, interrupts are
	//                    allowed
	// * D - Decimal    : if enabled, when add/carry/sub ops are executed, the source
	//                    are encoded in BCD
	// * B - Break      : Set when a software interrupt (like BRK) is executed
	// * V - Overflow   : If an arithmetic operation produces an overflow value,
	//                   this value is set
	// * S - Sign Flag  : 1 if the result of an operation is negative
	flg uint8

	// The Program Counter of the 6502
	pc uint16

	// The Stack Pointer of the NES
	sp uint8

	mem *Memory
}

// Where the program loaded in the memory gets executed per instruction
func (s *CPU) Tick() {
	current_inst := s.mem.Read(s.pc)
	if current_inst == 0x69 {
		s.ADC(immediate)
	} else if current_inst == 0x65 {
		s.ADC(zeropage)
	} else if current_inst == 0x75 {
		s.ADC(zeropagex)
	} else if current_inst == 0x6D {
		s.ADC(absolute)
	} else if current_inst == 0x7D {
		s.ADC(absolutex)
	} else if current_inst == 0x79 {
		s.ADC(absolutey)
	} else if current_inst == 0x7D {
		s.ADC(absolutex)
	} else if current_inst == 0x7D {
		s.ADC(absolutex)
	}

}

func (s *CPU) GetMemory() *Memory {
	return s.mem
}

func (s *CPU) GetAcc() uint8 {
	return s.acc
}

func (s *CPU) GetX() uint8 {
	return s.x
}

func (s *CPU) GetY() uint8 {
	return s.y
}

func (s *CPU) getFlag(flag flag) bool {
	return (s.flg & uint8(flag)) > 0
}

func (s *CPU) setFlagRegister(value uint8) {
	s.flg = value
}

func (s *CPU) getCarry() uint8 {
	return s.flg & 0b1
}

func NewCPU() *CPU {
	_6502 := new(CPU)
	_6502.mem = InitializeMemory()
	return _6502
}
