package atomix

type (
	MemoryOrderLoad interface {
		load()
	}
	MemoryOrderStore interface {
		store()
	}
	MemoryOrderAll interface {
		all()
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

func (relaxed) all() {}
func (acquire) all() {}
func (release) all() {}
func (acqRel) all()  {}
func (seqCst) all()  {}

var (
	MemoryOrderRelaxed = relaxed{}
	MemoryOrderAcquire = acquire{}
	MemoryOrderRelease = release{}
	MemoryOrderAcqRel  = acqRel{}
	MemoryOrderSeqCst  = seqCst{}
)
