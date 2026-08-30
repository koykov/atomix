package atomix

import "unsafe"

//go:noescape
//go:nosplit
//go:noinline
func LoadInt32Relaxed(addr *int32) int32

//go:noescape
//go:nosplit
//go:noinline
func LoadUint32Relaxed(addr *uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func LoadInt64Relaxed(addr *int64) int64

//go:noescape
//go:nosplit
//go:noinline
func LoadUint64Relaxed(addr *uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func LoadUintptrRelaxed(addr *uintptr) uintptr

//go:noescape
//go:nosplit
//go:noinline
func LoadPointerRelaxed(addr *unsafe.Pointer) unsafe.Pointer

// ...
