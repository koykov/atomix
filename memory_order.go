package atomix

type MemoryOrder uint8

const (
	MemoryOrderRelaxed MemoryOrder = iota
	MemoryOrderAcquire
	MemoryOrderConsume // deprecated, fallthrough to MemoryOrderSeqCnt
	MemoryOrderRelease
	MemoryOrderAcqRel
	MemoryOrderSeqCnt
)
