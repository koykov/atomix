package atomix

type (
	MemoryOrderLoad interface {
		load()
	}
	MemoryOrderStore interface {
		store()
	}
	MemoryOrderFull interface {
		MemoryOrderLoad
		MemoryOrderStore
	}
	MemoryOrderNoReturn interface {
		noret()
	}
	MemoryOrderReturn interface {
		ret()
	}
)

type (
	relaxed struct{}
	acquire struct{}
	release struct{}
	acqRel  struct{}
	seqCst  struct{}
)

func (relaxed) load() {}
func (acquire) load() {}
func (seqCst) load()  {}

func (relaxed) store() {}
func (release) store() {}
func (seqCst) store()  {}

func (acqRel) noret() {}
func (acqRel) ret()   {}

var (
	MemoryOrderRelaxed = relaxed{}
	MemoryOrderAcquire = acquire{}
	MemoryOrderRelease = release{}
	MemoryOrderAcqRel  = acqRel{}
	MemoryOrderSeqCst  = seqCst{}
)
