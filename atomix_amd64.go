package atomix

import "unsafe"

// ============================================================
// LOAD operations (Relaxed, Acquire, SeqCst)
// ============================================================

// Relaxed

//go:noescape
//go:nosplit
//go:noinline
func loadInt32Relaxed(addr *int32) int32

//go:noescape
//go:nosplit
//go:noinline
func loadUint32Relaxed(addr *uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func loadInt64Relaxed(addr *int64) int64

//go:noescape
//go:nosplit
//go:noinline
func loadUint64Relaxed(addr *uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func loadUintptrRelaxed(addr *uintptr) uintptr

//go:noescape
//go:nosplit
//go:noinline
func loadPointerRelaxed(addr *unsafe.Pointer) unsafe.Pointer

// Acquire

//go:noescape
//go:nosplit
//go:noinline
func loadInt32Acquire(addr *int32) int32

//go:noescape
//go:nosplit
//go:noinline
func loadUint32Acquire(addr *uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func loadInt64Acquire(addr *int64) int64

//go:noescape
//go:nosplit
//go:noinline
func loadUint64Acquire(addr *uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func loadUintptrAcquire(addr *uintptr) uintptr

//go:noescape
//go:nosplit
//go:noinline
func loadPointerAcquire(addr *unsafe.Pointer) unsafe.Pointer

// SeqCst

//go:noescape
//go:nosplit
//go:noinline
func loadInt32SeqCst(addr *int32) int32

//go:noescape
//go:nosplit
//go:noinline
func loadUint32SeqCst(addr *uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func loadInt64SeqCst(addr *int64) int64

//go:noescape
//go:nosplit
//go:noinline
func loadUint64SeqCst(addr *uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func loadUintptrSeqCst(addr *uintptr) uintptr

//go:noescape
//go:nosplit
//go:noinline
func loadPointerSeqCst(addr *unsafe.Pointer) unsafe.Pointer

// ============================================================
// STORE operations (Relaxed, Release, SeqCst)
// ============================================================

// Relaxed

//go:noescape
//go:nosplit
//go:noinline
func storeInt32Relaxed(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func storeUint32Relaxed(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func storeInt64Relaxed(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func storeUint64Relaxed(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func storeUintptrRelaxed(addr *uintptr, val uintptr)

// Release

//go:noescape
//go:nosplit
//go:noinline
func storeInt32Release(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func storeUint32Release(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func storeInt64Release(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func storeUint64Release(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func storeUintptrRelease(addr *uintptr, val uintptr)

// SeqCst

//go:noescape
//go:nosplit
//go:noinline
func storeInt32SeqCst(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func storeUint32SeqCst(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func storeInt64SeqCst(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func storeUint64SeqCst(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func storeUintptrSeqCst(addr *uintptr, val uintptr)

// ============================================================
// CAS operations (Relaxed, Acquire, Release, AcqRel, SeqCst)
// ============================================================

// Relaxed

//go:noescape
//go:nosplit
//go:noinline
func casInt32Relaxed(addr *int32, old, new int32) bool

//go:noescape
//go:nosplit
//go:noinline
func casUint32Relaxed(addr *uint32, old, new uint32) bool

//go:noescape
//go:nosplit
//go:noinline
func casInt64Relaxed(addr *int64, old, new int64) bool

//go:noescape
//go:nosplit
//go:noinline
func casUint64Relaxed(addr *uint64, old, new uint64) bool

//go:noescape
//go:nosplit
//go:noinline
func casUintptrRelaxed(addr *uintptr, old, new uintptr) bool

// Acquire

//go:noescape
//go:nosplit
//go:noinline
func casInt32Acquire(addr *int32, old, new int32) bool

//go:noescape
//go:nosplit
//go:noinline
func casUint32Acquire(addr *uint32, old, new uint32) bool

//go:noescape
//go:nosplit
//go:noinline
func casInt64Acquire(addr *int64, old, new int64) bool

//go:noescape
//go:nosplit
//go:noinline
func casUint64Acquire(addr *uint64, old, new uint64) bool

//go:noescape
//go:nosplit
//go:noinline
func casUintptrAcquire(addr *uintptr, old, new uintptr) bool

// Release

//go:noescape
//go:nosplit
//go:noinline
func casInt32Release(addr *int32, old, new int32) bool

//go:noescape
//go:nosplit
//go:noinline
func casUint32Release(addr *uint32, old, new uint32) bool

//go:noescape
//go:nosplit
//go:noinline
func casInt64Release(addr *int64, old, new int64) bool

//go:noescape
//go:nosplit
//go:noinline
func casUint64Release(addr *uint64, old, new uint64) bool

//go:noescape
//go:nosplit
//go:noinline
func casUintptrRelease(addr *uintptr, old, new uintptr) bool

// AcqRel

//go:noescape
//go:nosplit
//go:noinline
func casInt32AcqRel(addr *int32, old, new int32) bool

//go:noescape
//go:nosplit
//go:noinline
func casUint32AcqRel(addr *uint32, old, new uint32) bool

//go:noescape
//go:nosplit
//go:noinline
func casInt64AcqRel(addr *int64, old, new int64) bool

//go:noescape
//go:nosplit
//go:noinline
func casUint64AcqRel(addr *uint64, old, new uint64) bool

//go:noescape
//go:nosplit
//go:noinline
func casUintptrAcqRel(addr *uintptr, old, new uintptr) bool

// SeqCst

//go:noescape
//go:nosplit
//go:noinline
func casInt32SeqCst(addr *int32, old, new int32) bool

//go:noescape
//go:nosplit
//go:noinline
func casUint32SeqCst(addr *uint32, old, new uint32) bool

//go:noescape
//go:nosplit
//go:noinline
func casInt64SeqCst(addr *int64, old, new int64) bool

//go:noescape
//go:nosplit
//go:noinline
func casUint64SeqCst(addr *uint64, old, new uint64) bool

//go:noescape
//go:nosplit
//go:noinline
func casUintptrSeqCst(addr *uintptr, old, new uintptr) bool

// ============================================================
// ADD operations (Relaxed, Acquire, Release, AcqRel, SeqCst)
// ============================================================

// Relaxed

//go:noescape
//go:nosplit
//go:noinline
func addInt32Relaxed(addr *int32, delta int32) int32

//go:noescape
//go:nosplit
//go:noinline
func addUint32Relaxed(addr *uint32, delta uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func addInt64Relaxed(addr *int64, delta int64) int64

//go:noescape
//go:nosplit
//go:noinline
func addUint64Relaxed(addr *uint64, delta uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func addUintptrRelaxed(addr *uintptr, delta uintptr) uintptr

// Acquire

//go:noescape
//go:nosplit
//go:noinline
func addInt32Acquire(addr *int32, delta int32) int32

//go:noescape
//go:nosplit
//go:noinline
func addUint32Acquire(addr *uint32, delta uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func addInt64Acquire(addr *int64, delta int64) int64

//go:noescape
//go:nosplit
//go:noinline
func addUint64Acquire(addr *uint64, delta uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func addUintptrAcquire(addr *uintptr, delta uintptr) uintptr

// Release

//go:noescape
//go:nosplit
//go:noinline
func addInt32Release(addr *int32, delta int32) int32

//go:noescape
//go:nosplit
//go:noinline
func addUint32Release(addr *uint32, delta uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func addInt64Release(addr *int64, delta int64) int64

//go:noescape
//go:nosplit
//go:noinline
func addUint64Release(addr *uint64, delta uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func addUintptrRelease(addr *uintptr, delta uintptr) uintptr

// AcqRel

//go:noescape
//go:nosplit
//go:noinline
func addInt32AcqRel(addr *int32, delta int32) int32

//go:noescape
//go:nosplit
//go:noinline
func addUint32AcqRel(addr *uint32, delta uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func addInt64AcqRel(addr *int64, delta int64) int64

//go:noescape
//go:nosplit
//go:noinline
func addUint64AcqRel(addr *uint64, delta uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func addUintptrAcqRel(addr *uintptr, delta uintptr) uintptr

// SeqCst

//go:noescape
//go:nosplit
//go:noinline
func addInt32SeqCst(addr *int32, delta int32) int32

//go:noescape
//go:nosplit
//go:noinline
func addUint32SeqCst(addr *uint32, delta uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func addInt64SeqCst(addr *int64, delta int64) int64

//go:noescape
//go:nosplit
//go:noinline
func addUint64SeqCst(addr *uint64, delta uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func addUintptrSeqCst(addr *uintptr, delta uintptr) uintptr

// ============================================================
// SWAP operations (Relaxed, Acquire, Release, AcqRel, SeqCst)
// ============================================================

// Relaxed

//go:noescape
//go:nosplit
//go:noinline
func swapInt32Relaxed(addr *int32, new int32) int32

//go:noescape
//go:nosplit
//go:noinline
func swapUint32Relaxed(addr *uint32, new uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func swapInt64Relaxed(addr *int64, new int64) int64

//go:noescape
//go:nosplit
//go:noinline
func swapUint64Relaxed(addr *uint64, new uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func swapUintptrRelaxed(addr *uintptr, new uintptr) uintptr

// Acquire

//go:noescape
//go:nosplit
//go:noinline
func swapInt32Acquire(addr *int32, new int32) int32

//go:noescape
//go:nosplit
//go:noinline
func swapUint32Acquire(addr *uint32, new uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func swapInt64Acquire(addr *int64, new int64) int64

//go:noescape
//go:nosplit
//go:noinline
func swapUint64Acquire(addr *uint64, new uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func swapUintptrAcquire(addr *uintptr, new uintptr) uintptr

// Release

//go:noescape
//go:nosplit
//go:noinline
func swapInt32Release(addr *int32, new int32) int32

//go:noescape
//go:nosplit
//go:noinline
func swapUint32Release(addr *uint32, new uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func swapInt64Release(addr *int64, new int64) int64

//go:noescape
//go:nosplit
//go:noinline
func swapUint64Release(addr *uint64, new uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func swapUintptrRelease(addr *uintptr, new uintptr) uintptr

// AcqRel

//go:noescape
//go:nosplit
//go:noinline
func swapInt32AcqRel(addr *int32, new int32) int32

//go:noescape
//go:nosplit
//go:noinline
func swapUint32AcqRel(addr *uint32, new uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func swapInt64AcqRel(addr *int64, new int64) int64

//go:noescape
//go:nosplit
//go:noinline
func swapUint64AcqRel(addr *uint64, new uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func swapUintptrAcqRel(addr *uintptr, new uintptr) uintptr

// SeqCst

//go:noescape
//go:nosplit
//go:noinline
func swapInt32SeqCst(addr *int32, new int32) int32

//go:noescape
//go:nosplit
//go:noinline
func swapUint32SeqCst(addr *uint32, new uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func swapInt64SeqCst(addr *int64, new int64) int64

//go:noescape
//go:nosplit
//go:noinline
func swapUint64SeqCst(addr *uint64, new uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func swapUintptrSeqCst(addr *uintptr, new uintptr) uintptr

// ============================================================
// BITWISE OR operations
// ============================================================

// Or (no return): Relaxed, Acquire, Release, SeqCst
// Relaxed

//go:noescape
//go:nosplit
//go:noinline
func orInt32Relaxed(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func orUint32Relaxed(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func orInt64Relaxed(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func orUint64Relaxed(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func orUintptrRelaxed(addr *uintptr, val uintptr)

// Acquire

//go:noescape
//go:nosplit
//go:noinline
func orInt32Acquire(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func orUint32Acquire(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func orInt64Acquire(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func orUint64Acquire(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func orUintptrAcquire(addr *uintptr, val uintptr)

// Release

//go:noescape
//go:nosplit
//go:noinline
func orInt32Release(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func orUint32Release(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func orInt64Release(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func orUint64Release(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func orUintptrRelease(addr *uintptr, val uintptr)

// SeqCst

//go:noescape
//go:nosplit
//go:noinline
func orInt32SeqCst(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func orUint32SeqCst(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func orInt64SeqCst(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func orUint64SeqCst(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func orUintptrSeqCst(addr *uintptr, val uintptr)

// or return: AcqRel (full barrier)

//go:noescape
//go:nosplit
//go:noinline
func or32AcqRel(addr *uint32, val uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func or64AcqRel(addr *uint64, val uint64) uint64

// ============================================================
// BITWISE AND operations
// ============================================================

// And (no return value): Relaxed, Acquire, Release, SeqCst
// Relaxed

//go:noescape
//go:nosplit
//go:noinline
func andInt32Relaxed(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func andUint32Relaxed(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func andInt64Relaxed(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func andUint64Relaxed(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func andUintptrRelaxed(addr *uintptr, val uintptr)

// Acquire

//go:noescape
//go:nosplit
//go:noinline
func andInt32Acquire(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func andUint32Acquire(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func andInt64Acquire(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func andUint64Acquire(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func andUintptrAcquire(addr *uintptr, val uintptr)

// Release

//go:noescape
//go:nosplit
//go:noinline
func andInt32Release(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func andUint32Release(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func andInt64Release(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func andUint64Release(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func andUintptrRelease(addr *uintptr, val uintptr)

// SeqCst

//go:noescape
//go:nosplit
//go:noinline
func andInt32SeqCst(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func andUint32SeqCst(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func andInt64SeqCst(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func andUint64SeqCst(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func andUintptrSeqCst(addr *uintptr, val uintptr)

// And with return value: AcqRel (full barrier)

//go:noescape
//go:nosplit
//go:noinline
func and32AcqRel(addr *uint32, val uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func and64AcqRel(addr *uint64, val uint64) uint64
