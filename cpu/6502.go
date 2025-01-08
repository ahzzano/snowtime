package cpu

// Handles the logic of the CPU. Stores the acc, x,y, flg, registers along with
// the program counter and stack pointer.
// For more details, kindly read
//
// https://www.nesdev.org/6502.txt
//

type CPU struct {
	acc uint8
	x   uint8
	y   uint8
	flg uint8
	pc  uint8
	sp  uint8
}

func (s *CPU) tick() {
	s.pc += 1
}
