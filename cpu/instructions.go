package cpu

import "fmt"

type addressing_mode uint8

// The addressing mode of each instruction represented with the
// addr_mode enum
const (
	implicit   addressing_mode = iota
	immediate                  // Get the immediate value from PC + 1
	zeropage                   // Get the memory addr from PC + 1  thats only two nibbles
	zeropage_x                 // Get the memory addr from (PC + 1) + x thats only two nibbles
	zeropage_y                 // Get the memory addr from (PC + 1) + y thats only two nibbles
	relative                   // Offset is added to the PC
	absolute                   // Get the memory addr from (PC+1) and (PC+2)
	absolute_x                 //  Get the memory addr from (PC + 1 + x) and (PC + 2 + x)
	absolute_y                 // Get the memory addr from (PC + 1 + y)  and (PC + 2 + y)
	indirect
	indexed_indirect
	indirect_indexed
)

//	ADC
//
// Add Memory to the accumulator with carry
type flag uint8

// the flag bits
const (
	carry     flag = 0b00000001
	zero      flag = 0b00000010
	interrupt flag = 0b00000100
	decimal   flag = 0b00001000
	brk       flag = 0b00010000
	overflow  flag = 0b01000000
	sign      flag = 0b10000000
)

func (s *CPU) ADC(addrm addressing_mode) {
	if addrm == immediate {
		target_addr := s.pc + 1
		s.acc = s.acc + s.getCarry() + s.mem.Read(target_addr)
		s.pc += 2
	} else if addrm == zeropage {
		target_addr := uint16(s.mem.Read(s.pc+1) & 0xFF)
		s.acc = s.acc + s.getCarry() + s.mem.Read(target_addr)
		s.pc += 2
	} else if addrm == zeropage_x {
		target_addr := (uint16(s.x) + s.pc + 1) & 0xFF
		s.acc = s.acc + s.getCarry() + s.mem.Read(target_addr)
		s.pc += 2
	} else if addrm == absolute {
		target_addr := s.mem.Read16Alt(s.pc+2, s.pc+1)
		s.acc = s.acc + s.getCarry() + s.mem.Read(target_addr)
		s.pc += 3
	} else if addrm == absolute_x {
		target_addr := s.mem.Read16Alt(s.pc+2+uint16(s.x),
			s.pc+1+uint16(s.x))
		s.acc += s.getCarry() + s.mem.Read(target_addr)
		s.pc += 3
	} else if addrm == absolute_y {
		target_addr := s.mem.Read16Alt(s.pc+2+uint16(s.y),
			s.pc+1+uint16(s.y))
		s.acc += s.getCarry() + s.mem.Read(target_addr)
		s.pc += 3
	} else if addrm == indexed_indirect {
		target_addr := (uint16(s.x) + s.pc + 1) & 0xFF
		target_addr = s.mem.Read16(target_addr)
		s.acc = s.acc + s.getCarry() + s.mem.Read(target_addr)
		s.pc += 2
	} else if addrm == indirect_indexed {
		target_addr := (uint16(s.y) + s.pc + 1) & 0xFF
		s.acc = s.acc + s.getCarry() + s.mem.Read(target_addr)
		s.pc += 2
	} else {
		fmt.Errorf("This is currently unimplemented")
	}
}
