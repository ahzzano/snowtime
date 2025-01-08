package cpu

type addr_mode uint8

// The addressing mode of each instruction represented with the
// addr_mode enum
const (
	implicit addr_mode = iota
	immediate
	zeropage
	zeropagex
	zeropagey
	relative
	absolute
	absolutex
	absolutey
	indrect
	indexed_indirect
	indirect_indexed
)
