package atomix

type MemoryOrder uint8

const (
	MemoryOrderRelaxed MemoryOrder = iota
	MemoryOrderAcquire
	MemoryOrderRelease
	MemoryOrderAcqRel
	MemoryOrderSeqCnt
	// MemoryOrderConsume - deprecated, not implement
)
