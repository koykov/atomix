package pbvector

import (
	"reflect"
	"sync"
	"unsafe"
)

type registry struct {
	mux sync.RWMutex
	idx map[uintptr]uint32
	buf []pbtyp
}

func (r *registry) getType(x any) *pbtyp {
	t := reflect.TypeOf(x)
	pt := uintptr(unsafe.Pointer(&t))

	r.mux.RLock()
	i, ok := r.idx[pt]
	if ok && i < uint32(len(r.buf)) {
		r.mux.RUnlock()
		return &r.buf[i]
	}
	r.mux.Unlock()

	r.mux.Lock()
	defer r.mux.Unlock()
	pbt := parseType(x)
	r.buf = append(r.buf, pbt)
	i = uint32(len(r.buf) - 1)
	r.idx[pt] = i

	return &r.buf[i]
}

var registry_ = registry{idx: make(map[uintptr]uint32)}
