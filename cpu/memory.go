package cpu

/*
Memory handles the memory needed for the 6502
*/
type Memory struct {
	contents [65536]uint8 // the contents of our memory
}

func InitializeMemory() *Memory {
	mem := new(Memory)
	mem.contents = [65536]uint8{0}
	return mem
}

func (m *Memory) Write(addr uint16, value uint8) {
	m.contents[addr] = value
}

func (m *Memory) Read(addr uint16) uint8 {
	return m.contents[addr]
}

func (m *Memory) Write16(addr uint16, value uint16) {
	lo := uint8(value | 0x00FF)
	hi := uint8((value | 0xFF00) >> 8)

	m.contents[addr] = hi
	m.contents[addr+1] = lo
}

func (m *Memory) Read16(addr uint16) uint16 {
	lo := uint16(addr)
	hi := uint16(addr) + 1
	return uint16(m.contents[hi])<<8 | uint16(m.contents[lo])
}

func (m *Memory) Read16Alt(hi, lo uint16) uint16 {
	hi_value := uint16(m.contents[hi]) << 8
	lo_value := uint16(m.contents[lo])
	return hi_value | lo_value
}
