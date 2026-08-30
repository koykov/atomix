package atomix

import "unsafe"

// ============================================================
// LOAD operations (Relaxed, Acquire, SeqCst)
// ============================================================

// Relaxed

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

// Acquire

//go:noescape
//go:nosplit
//go:noinline
func LoadInt32Acquire(addr *int32) int32

//go:noescape
//go:nosplit
//go:noinline
func LoadUint32Acquire(addr *uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func LoadInt64Acquire(addr *int64) int64

//go:noescape
//go:nosplit
//go:noinline
func LoadUint64Acquire(addr *uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func LoadUintptrAcquire(addr *uintptr) uintptr

//go:noescape
//go:nosplit
//go:noinline
func LoadPointerAcquire(addr *unsafe.Pointer) unsafe.Pointer

// SeqCst

//go:noescape
//go:nosplit
//go:noinline
func LoadInt32SeqCst(addr *int32) int32

//go:noescape
//go:nosplit
//go:noinline
func LoadUint32SeqCst(addr *uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func LoadInt64SeqCst(addr *int64) int64

//go:noescape
//go:nosplit
//go:noinline
func LoadUint64SeqCst(addr *uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func LoadUintptrSeqCst(addr *uintptr) uintptr

//go:noescape
//go:nosplit
//go:noinline
func LoadPointerSeqCst(addr *unsafe.Pointer) unsafe.Pointer

// ============================================================
// STORE operations (Relaxed, Release, SeqCst)
// ============================================================

// Relaxed

//go:noescape
//go:nosplit
//go:noinline
func StoreInt32Relaxed(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func StoreUint32Relaxed(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func StoreInt64Relaxed(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func StoreUint64Relaxed(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func StoreUintptrRelaxed(addr *uintptr, val uintptr)

// Release

//go:noescape
//go:nosplit
//go:noinline
func StoreInt32Release(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func StoreUint32Release(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func StoreInt64Release(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func StoreUint64Release(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func StoreUintptrRelease(addr *uintptr, val uintptr)

// SeqCst

//go:noescape
//go:nosplit
//go:noinline
func StoreInt32SeqCst(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func StoreUint32SeqCst(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func StoreInt64SeqCst(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func StoreUint64SeqCst(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func StoreUintptrSeqCst(addr *uintptr, val uintptr)

// ============================================================
// CAS operations (Relaxed, Acquire, Release, AcqRel, SeqCst)
// ============================================================

// Relaxed

//go:noescape
//go:nosplit
//go:noinline
func CasInt32Relaxed(addr *int32, old, new int32) bool

//go:noescape
//go:nosplit
//go:noinline
func CasUint32Relaxed(addr *uint32, old, new uint32) bool

//go:noescape
//go:nosplit
//go:noinline
func CasInt64Relaxed(addr *int64, old, new int64) bool

//go:noescape
//go:nosplit
//go:noinline
func CasUint64Relaxed(addr *uint64, old, new uint64) bool

//go:noescape
//go:nosplit
//go:noinline
func CasUintptrRelaxed(addr *uintptr, old, new uintptr) bool

// Acquire

//go:noescape
//go:nosplit
//go:noinline
func CasInt32Acquire(addr *int32, old, new int32) bool

//go:noescape
//go:nosplit
//go:noinline
func CasUint32Acquire(addr *uint32, old, new uint32) bool

//go:noescape
//go:nosplit
//go:noinline
func CasInt64Acquire(addr *int64, old, new int64) bool

//go:noescape
//go:nosplit
//go:noinline
func CasUint64Acquire(addr *uint64, old, new uint64) bool

//go:noescape
//go:nosplit
//go:noinline
func CasUintptrAcquire(addr *uintptr, old, new uintptr) bool

// Release

//go:noescape
//go:nosplit
//go:noinline
func CasInt32Release(addr *int32, old, new int32) bool

//go:noescape
//go:nosplit
//go:noinline
func CasUint32Release(addr *uint32, old, new uint32) bool

//go:noescape
//go:nosplit
//go:noinline
func CasInt64Release(addr *int64, old, new int64) bool

//go:noescape
//go:nosplit
//go:noinline
func CasUint64Release(addr *uint64, old, new uint64) bool

//go:noescape
//go:nosplit
//go:noinline
func CasUintptrRelease(addr *uintptr, old, new uintptr) bool

// AcqRel

//go:noescape
//go:nosplit
//go:noinline
func CasInt32AcqRel(addr *int32, old, new int32) bool

//go:noescape
//go:nosplit
//go:noinline
func CasUint32AcqRel(addr *uint32, old, new uint32) bool

//go:noescape
//go:nosplit
//go:noinline
func CasInt64AcqRel(addr *int64, old, new int64) bool

//go:noescape
//go:nosplit
//go:noinline
func CasUint64AcqRel(addr *uint64, old, new uint64) bool

//go:noescape
//go:nosplit
//go:noinline
func CasUintptrAcqRel(addr *uintptr, old, new uintptr) bool

// SeqCst

//go:noescape
//go:nosplit
//go:noinline
func CasInt32SeqCst(addr *int32, old, new int32) bool

//go:noescape
//go:nosplit
//go:noinline
func CasUint32SeqCst(addr *uint32, old, new uint32) bool

//go:noescape
//go:nosplit
//go:noinline
func CasInt64SeqCst(addr *int64, old, new int64) bool

//go:noescape
//go:nosplit
//go:noinline
func CasUint64SeqCst(addr *uint64, old, new uint64) bool

//go:noescape
//go:nosplit
//go:noinline
func CasUintptrSeqCst(addr *uintptr, old, new uintptr) bool

// ============================================================
// ADD operations (Relaxed, Acquire, Release, AcqRel, SeqCst)
// ============================================================

// Relaxed

//go:noescape
//go:nosplit
//go:noinline
func AddInt32Relaxed(addr *int32, delta int32) int32

//go:noescape
//go:nosplit
//go:noinline
func AddUint32Relaxed(addr *uint32, delta uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func AddInt64Relaxed(addr *int64, delta int64) int64

//go:noescape
//go:nosplit
//go:noinline
func AddUint64Relaxed(addr *uint64, delta uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func AddUintptrRelaxed(addr *uintptr, delta uintptr) uintptr

// Acquire

//go:noescape
//go:nosplit
//go:noinline
func AddInt32Acquire(addr *int32, delta int32) int32

//go:noescape
//go:nosplit
//go:noinline
func AddUint32Acquire(addr *uint32, delta uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func AddInt64Acquire(addr *int64, delta int64) int64

//go:noescape
//go:nosplit
//go:noinline
func AddUint64Acquire(addr *uint64, delta uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func AddUintptrAcquire(addr *uintptr, delta uintptr) uintptr

// Release

//go:noescape
//go:nosplit
//go:noinline
func AddInt32Release(addr *int32, delta int32) int32

//go:noescape
//go:nosplit
//go:noinline
func AddUint32Release(addr *uint32, delta uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func AddInt64Release(addr *int64, delta int64) int64

//go:noescape
//go:nosplit
//go:noinline
func AddUint64Release(addr *uint64, delta uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func AddUintptrRelease(addr *uintptr, delta uintptr) uintptr

// AcqRel

//go:noescape
//go:nosplit
//go:noinline
func AddInt32AcqRel(addr *int32, delta int32) int32

//go:noescape
//go:nosplit
//go:noinline
func AddUint32AcqRel(addr *uint32, delta uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func AddInt64AcqRel(addr *int64, delta int64) int64

//go:noescape
//go:nosplit
//go:noinline
func AddUint64AcqRel(addr *uint64, delta uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func AddUintptrAcqRel(addr *uintptr, delta uintptr) uintptr

// SeqCst

//go:noescape
//go:nosplit
//go:noinline
func AddInt32SeqCst(addr *int32, delta int32) int32

//go:noescape
//go:nosplit
//go:noinline
func AddUint32SeqCst(addr *uint32, delta uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func AddInt64SeqCst(addr *int64, delta int64) int64

//go:noescape
//go:nosplit
//go:noinline
func AddUint64SeqCst(addr *uint64, delta uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func AddUintptrSeqCst(addr *uintptr, delta uintptr) uintptr

// ============================================================
// SWAP operations (Relaxed, Acquire, Release, AcqRel, SeqCst)
// ============================================================

// Relaxed

//go:noescape
//go:nosplit
//go:noinline
func SwapInt32Relaxed(addr *int32, new int32) int32

//go:noescape
//go:nosplit
//go:noinline
func SwapUint32Relaxed(addr *uint32, new uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func SwapInt64Relaxed(addr *int64, new int64) int64

//go:noescape
//go:nosplit
//go:noinline
func SwapUint64Relaxed(addr *uint64, new uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func SwapUintptrRelaxed(addr *uintptr, new uintptr) uintptr

// Acquire

//go:noescape
//go:nosplit
//go:noinline
func SwapInt32Acquire(addr *int32, new int32) int32

//go:noescape
//go:nosplit
//go:noinline
func SwapUint32Acquire(addr *uint32, new uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func SwapInt64Acquire(addr *int64, new int64) int64

//go:noescape
//go:nosplit
//go:noinline
func SwapUint64Acquire(addr *uint64, new uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func SwapUintptrAcquire(addr *uintptr, new uintptr) uintptr

// Release

//go:noescape
//go:nosplit
//go:noinline
func SwapInt32Release(addr *int32, new int32) int32

//go:noescape
//go:nosplit
//go:noinline
func SwapUint32Release(addr *uint32, new uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func SwapInt64Release(addr *int64, new int64) int64

//go:noescape
//go:nosplit
//go:noinline
func SwapUint64Release(addr *uint64, new uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func SwapUintptrRelease(addr *uintptr, new uintptr) uintptr

// AcqRel

//go:noescape
//go:nosplit
//go:noinline
func SwapInt32AcqRel(addr *int32, new int32) int32

//go:noescape
//go:nosplit
//go:noinline
func SwapUint32AcqRel(addr *uint32, new uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func SwapInt64AcqRel(addr *int64, new int64) int64

//go:noescape
//go:nosplit
//go:noinline
func SwapUint64AcqRel(addr *uint64, new uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func SwapUintptrAcqRel(addr *uintptr, new uintptr) uintptr

// SeqCst

//go:noescape
//go:nosplit
//go:noinline
func SwapInt32SeqCst(addr *int32, new int32) int32

//go:noescape
//go:nosplit
//go:noinline
func SwapUint32SeqCst(addr *uint32, new uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func SwapInt64SeqCst(addr *int64, new int64) int64

//go:noescape
//go:nosplit
//go:noinline
func SwapUint64SeqCst(addr *uint64, new uint64) uint64

//go:noescape
//go:nosplit
//go:noinline
func SwapUintptrSeqCst(addr *uintptr, new uintptr) uintptr

// ============================================================
// BITWISE OR operations
// ============================================================

// Or (no return): Relaxed, Acquire, Release, SeqCst
// Relaxed

//go:noescape
//go:nosplit
//go:noinline
func OrInt32Relaxed(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func OrUint32Relaxed(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func OrInt64Relaxed(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func OrUint64Relaxed(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func OrUintptrRelaxed(addr *uintptr, val uintptr)

// Acquire

//go:noescape
//go:nosplit
//go:noinline
func OrInt32Acquire(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func OrUint32Acquire(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func OrInt64Acquire(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func OrUint64Acquire(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func OrUintptrAcquire(addr *uintptr, val uintptr)

// Release

//go:noescape
//go:nosplit
//go:noinline
func OrInt32Release(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func OrUint32Release(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func OrInt64Release(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func OrUint64Release(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func OrUintptrRelease(addr *uintptr, val uintptr)

// SeqCst

//go:noescape
//go:nosplit
//go:noinline
func OrInt32SeqCst(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func OrUint32SeqCst(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func OrInt64SeqCst(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func OrUint64SeqCst(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func OrUintptrSeqCst(addr *uintptr, val uintptr)

// Or return: AcqRel (full barrier)

//go:noescape
//go:nosplit
//go:noinline
func Or32AcqRel(addr *uint32, val uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func Or64AcqRel(addr *uint64, val uint64) uint64

// ============================================================
// BITWISE AND operations
// ============================================================

// And (no return value): Relaxed, Acquire, Release, SeqCst
// Relaxed

//go:noescape
//go:nosplit
//go:noinline
func AndInt32Relaxed(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func AndUint32Relaxed(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func AndInt64Relaxed(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func AndUint64Relaxed(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func AndUintptrRelaxed(addr *uintptr, val uintptr)

// Acquire

//go:noescape
//go:nosplit
//go:noinline
func AndInt32Acquire(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func AndUint32Acquire(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func AndInt64Acquire(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func AndUint64Acquire(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func AndUintptrAcquire(addr *uintptr, val uintptr)

// Release

//go:noescape
//go:nosplit
//go:noinline
func AndInt32Release(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func AndUint32Release(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func AndInt64Release(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func AndUint64Release(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func AndUintptrRelease(addr *uintptr, val uintptr)

// SeqCst

//go:noescape
//go:nosplit
//go:noinline
func AndInt32SeqCst(addr *int32, val int32)

//go:noescape
//go:nosplit
//go:noinline
func AndUint32SeqCst(addr *uint32, val uint32)

//go:noescape
//go:nosplit
//go:noinline
func AndInt64SeqCst(addr *int64, val int64)

//go:noescape
//go:nosplit
//go:noinline
func AndUint64SeqCst(addr *uint64, val uint64)

//go:noescape
//go:nosplit
//go:noinline
func AndUintptrSeqCst(addr *uintptr, val uintptr)

// And with return value: AcqRel (full barrier)

//go:noescape
//go:nosplit
//go:noinline
func And32AcqRel(addr *uint32, val uint32) uint32

//go:noescape
//go:nosplit
//go:noinline
func And64AcqRel(addr *uint64, val uint64) uint64
