package atomix

import (
	"sync/atomic"
	"unsafe"
)

func SwapInt32(addr *int32, new int32, order MemoryOrderAll) int32 {
	switch order {
	case relaxed{}:
		return swapInt32Relaxed(addr, new)
	case acquire{}:
		return swapInt32Acquire(addr, new)
	case release{}:
		return swapInt32Release(addr, new)
	case acqRel{}:
		return swapInt32AcqRel(addr, new)
	case seqCst{}:
		fallthrough
	default:
		return swapInt32SeqCst(addr, new)
	}
}

func SwapUint32(addr *uint32, new uint32, order MemoryOrderAll) uint32 {
	switch order {
	case relaxed{}:
		return swapUint32Relaxed(addr, new)
	case acquire{}:
		return swapUint32Acquire(addr, new)
	case release{}:
		return swapUint32Release(addr, new)
	case acqRel{}:
		return swapUint32AcqRel(addr, new)
	case seqCst{}:
		fallthrough
	default:
		return swapUint32SeqCst(addr, new)
	}
}

func SwapInt64(addr *int64, new int64, order MemoryOrderAll) int64 {
	switch order {
	case relaxed{}:
		return swapInt64Relaxed(addr, new)
	case acquire{}:
		return swapInt64Acquire(addr, new)
	case release{}:
		return swapInt64Release(addr, new)
	case acqRel{}:
		return swapInt64AcqRel(addr, new)
	case seqCst{}:
		fallthrough
	default:
		return swapInt64SeqCst(addr, new)
	}
}

func SwapUint64(addr *uint64, new uint64, order MemoryOrderAll) uint64 {
	switch order {
	case relaxed{}:
		return swapUint64Relaxed(addr, new)
	case acquire{}:
		return swapUint64Acquire(addr, new)
	case release{}:
		return swapUint64Release(addr, new)
	case acqRel{}:
		return swapUint64AcqRel(addr, new)
	case seqCst{}:
		fallthrough
	default:
		return swapUint64SeqCst(addr, new)
	}
}

func SwapUintptr(addr *uintptr, new uintptr, order MemoryOrderAll) uintptr {
	switch order {
	case relaxed{}:
		return swapUintptrRelaxed(addr, new)
	case acquire{}:
		return swapUintptrAcquire(addr, new)
	case release{}:
		return swapUintptrRelease(addr, new)
	case acqRel{}:
		return swapUintptrAcqRel(addr, new)
	case seqCst{}:
		fallthrough
	default:
		return swapUintptrSeqCst(addr, new)
	}
}

func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer, _ MemoryOrderAll) unsafe.Pointer {
	// GC requires write barrier for storing pointers there is no way to check in non-runtime code is GC active or not.
	// Thus use native atomic.StorePointer without considering order.
	return atomic.SwapPointer(addr, new)
}

func CompareAndSwapInt32(addr *int32, old, new int32, order MemoryOrderAll) bool {
	switch order {
	case relaxed{}:
		return casInt32Relaxed(addr, old, new)
	case acquire{}:
		return casInt32Acquire(addr, old, new)
	case release{}:
		return casInt32Release(addr, old, new)
	case acqRel{}:
		return casInt32AcqRel(addr, old, new)
	case seqCst{}:
		fallthrough
	default:
		return casInt32SeqCst(addr, old, new)
	}
}

func CompareAndSwapUint32(addr *uint32, old, new uint32, order MemoryOrderAll) (swapped bool) {
	switch order {
	case relaxed{}:
		return casUint32Relaxed(addr, old, new)
	case acquire{}:
		return casUint32Acquire(addr, old, new)
	case release{}:
		return casUint32Release(addr, old, new)
	case acqRel{}:
		return casUint32AcqRel(addr, old, new)
	case seqCst{}:
		fallthrough
	default:
		return casUint32SeqCst(addr, old, new)
	}
}

func CompareAndSwapInt64(addr *int64, old, new int64, order MemoryOrderAll) bool {
	switch order {
	case relaxed{}:
		return casInt64Relaxed(addr, old, new)
	case acquire{}:
		return casInt64Acquire(addr, old, new)
	case release{}:
		return casInt64Release(addr, old, new)
	case acqRel{}:
		return casInt64AcqRel(addr, old, new)
	case seqCst{}:
		fallthrough
	default:
		return casInt64SeqCst(addr, old, new)
	}
}

func CompareAndSwapUint64(addr *uint64, old, new uint64, order MemoryOrderAll) (swapped bool) {
	switch order {
	case relaxed{}:
		return casUint64Relaxed(addr, old, new)
	case acquire{}:
		return casUint64Acquire(addr, old, new)
	case release{}:
		return casUint64Release(addr, old, new)
	case acqRel{}:
		return casUint64AcqRel(addr, old, new)
	case seqCst{}:
		fallthrough
	default:
		return casUint64SeqCst(addr, old, new)
	}
}

func CompareAndSwapUintptr(addr *uintptr, old, new uintptr, order MemoryOrderAll) (swapped bool) {
	switch order {
	case relaxed{}:
		return casUintptrRelaxed(addr, old, new)
	case acquire{}:
		return casUintptrAcquire(addr, old, new)
	case release{}:
		return casUintptrRelease(addr, old, new)
	case acqRel{}:
		return casUintptrAcqRel(addr, old, new)
	case seqCst{}:
		fallthrough
	default:
		return casUintptrSeqCst(addr, old, new)
	}
}

func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer, _ MemoryOrderAll) (swapped bool) {
	// GC requires write barrier for storing pointers there is no way to check in non-runtime code is GC active or not.
	// Thus use native atomic.StorePointer without considering order.
	return atomic.CompareAndSwapPointer(addr, old, new)
}

func AddInt32(addr *int32, delta int32, order MemoryOrderAll) int32 {
	switch order {
	case relaxed{}:
		return addInt32Relaxed(addr, delta)
	case acquire{}:
		return addInt32Acquire(addr, delta)
	case release{}:
		return addInt32Release(addr, delta)
	case acqRel{}:
		return addInt32AcqRel(addr, delta)
	case seqCst{}:
		fallthrough
	default:
		return addInt32SeqCst(addr, delta)
	}
}

func AddUint32(addr *uint32, delta uint32, order MemoryOrderAll) uint32 {
	switch order {
	case relaxed{}:
		return addUint32Relaxed(addr, delta)
	case acquire{}:
		return addUint32Acquire(addr, delta)
	case release{}:
		return addUint32Release(addr, delta)
	case acqRel{}:
		return addUint32AcqRel(addr, delta)
	case seqCst{}:
		fallthrough
	default:
		return addUint32SeqCst(addr, delta)
	}
}

func AddInt64(addr *int64, delta int64, order MemoryOrderAll) int64 {
	switch order {
	case relaxed{}:
		return addInt64Relaxed(addr, delta)
	case acquire{}:
		return addInt64Acquire(addr, delta)
	case release{}:
		return addInt64Release(addr, delta)
	case acqRel{}:
		return addInt64AcqRel(addr, delta)
	case seqCst{}:
		fallthrough
	default:
		return addInt64SeqCst(addr, delta)
	}
}

func AddUint64(addr *uint64, delta uint64, order MemoryOrderAll) uint64 {
	switch order {
	case relaxed{}:
		return addUint64Relaxed(addr, delta)
	case acquire{}:
		return addUint64Acquire(addr, delta)
	case release{}:
		return addUint64Release(addr, delta)
	case acqRel{}:
		return addUint64AcqRel(addr, delta)
	case seqCst{}:
		fallthrough
	default:
		return addUint64SeqCst(addr, delta)
	}
}

func AddUintptr(addr *uintptr, delta uintptr, order MemoryOrderAll) uintptr {
	switch order {
	case relaxed{}:
		return addUintptrRelaxed(addr, delta)
	case acquire{}:
		return addUintptrAcquire(addr, delta)
	case release{}:
		return addUintptrRelease(addr, delta)
	case acqRel{}:
		return addUintptrAcqRel(addr, delta)
	case seqCst{}:
		fallthrough
	default:
		return addUintptrSeqCst(addr, delta)
	}
}

func AndInt32(addr *int32, mask int32, order MemoryOrderAll) int32 {
	switch order {
	case relaxed{}:
		return andInt32RelaxedReturn(addr, mask)
	case acquire{}:
		return andInt32AcquireReturn(addr, mask)
	case release{}:
		return andInt32ReleaseReturn(addr, mask)
	case acqRel{}:
		return andInt32ReleaseReturn(addr, mask)
	case seqCst{}:
		fallthrough
	default:
		return andInt32SeqCstReturn(addr, mask)
	}
}

func AndUint32(addr *uint32, mask uint32, order MemoryOrderAll) uint32 {
	switch order {
	case relaxed{}:
		return andUint32RelaxedReturn(addr, mask)
	case acquire{}:
		return andUint32AcquireReturn(addr, mask)
	case release{}:
		return andUint32ReleaseReturn(addr, mask)
	case acqRel{}:
		return andUint32ReleaseReturn(addr, mask)
	case seqCst{}:
		fallthrough
	default:
		return andUint32SeqCstReturn(addr, mask)
	}
}

func AndInt64(addr *int64, mask int64, order MemoryOrderAll) int64 {
	switch order {
	case relaxed{}:
		return andInt64RelaxedReturn(addr, mask)
	case acquire{}:
		return andInt64AcquireReturn(addr, mask)
	case release{}:
		return andInt64ReleaseReturn(addr, mask)
	case acqRel{}:
		return andInt64ReleaseReturn(addr, mask)
	case seqCst{}:
		fallthrough
	default:
		return andInt64SeqCstReturn(addr, mask)
	}
}

func AndUint64(addr *uint64, mask uint64, order MemoryOrderAll) uint64 {
	switch order {
	case relaxed{}:
		return andUint64RelaxedReturn(addr, mask)
	case acquire{}:
		return andUint64AcquireReturn(addr, mask)
	case release{}:
		return andUint64ReleaseReturn(addr, mask)
	case acqRel{}:
		return andUint64ReleaseReturn(addr, mask)
	case seqCst{}:
		fallthrough
	default:
		return andUint64SeqCstReturn(addr, mask)
	}
}

func AndUintptr(addr *uintptr, mask uintptr, order MemoryOrderAll) uintptr {
	switch order {
	case relaxed{}:
		return andUintptrRelaxedReturn(addr, mask)
	case acquire{}:
		return andUintptrAcquireReturn(addr, mask)
	case release{}:
		return andUintptrReleaseReturn(addr, mask)
	case acqRel{}:
		return andUintptrReleaseReturn(addr, mask)
	case seqCst{}:
		fallthrough
	default:
		return andUintptrSeqCstReturn(addr, mask)
	}
}

func OrInt32(addr *int32, mask int32, order MemoryOrderAll) int32 {
	switch order {
	case relaxed{}:
		return orInt32RelaxedReturn(addr, mask)
	case acquire{}:
		return orInt32AcquireReturn(addr, mask)
	case release{}:
		return orInt32ReleaseReturn(addr, mask)
	case acqRel{}:
		return orInt32ReleaseReturn(addr, mask)
	case seqCst{}:
		fallthrough
	default:
		return orInt32SeqCstReturn(addr, mask)
	}
}

func OrUint32(addr *uint32, mask uint32, order MemoryOrderAll) uint32 {
	switch order {
	case relaxed{}:
		return orUint32RelaxedReturn(addr, mask)
	case acquire{}:
		return orUint32AcquireReturn(addr, mask)
	case release{}:
		return orUint32ReleaseReturn(addr, mask)
	case acqRel{}:
		return orUint32ReleaseReturn(addr, mask)
	case seqCst{}:
		fallthrough
	default:
		return orUint32SeqCstReturn(addr, mask)
	}
}

func OrInt64(addr *int64, mask int64, order MemoryOrderAll) int64 {
	switch order {
	case relaxed{}:
		return orInt64RelaxedReturn(addr, mask)
	case acquire{}:
		return orInt64AcquireReturn(addr, mask)
	case release{}:
		return orInt64ReleaseReturn(addr, mask)
	case acqRel{}:
		return orInt64ReleaseReturn(addr, mask)
	case seqCst{}:
		fallthrough
	default:
		return orInt64SeqCstReturn(addr, mask)
	}
}

func OrUint64(addr *uint64, mask uint64, order MemoryOrderAll) uint64 {
	switch order {
	case relaxed{}:
		return orUint64RelaxedReturn(addr, mask)
	case acquire{}:
		return orUint64AcquireReturn(addr, mask)
	case release{}:
		return orUint64ReleaseReturn(addr, mask)
	case acqRel{}:
		return orUint64ReleaseReturn(addr, mask)
	case seqCst{}:
		fallthrough
	default:
		return orUint64SeqCstReturn(addr, mask)
	}
}

func OrUintptr(addr *uintptr, mask uintptr, order MemoryOrderAll) uintptr {
	switch order {
	case relaxed{}:
		return orUintptrRelaxedReturn(addr, mask)
	case acquire{}:
		return orUintptrAcquireReturn(addr, mask)
	case release{}:
		return orUintptrReleaseReturn(addr, mask)
	case acqRel{}:
		return orUintptrReleaseReturn(addr, mask)
	case seqCst{}:
		fallthrough
	default:
		return orUintptrSeqCstReturn(addr, mask)
	}
}

func LoadInt32(addr *int32, order MemoryOrderLoad) int32 {
	switch order {
	case relaxed{}:
		return loadInt32Relaxed(addr)
	case acquire{}:
		return loadInt32Acquire(addr)
	case seqCst{}:
		fallthrough
	default:
		return loadInt32SeqCst(addr)
	}
}

func LoadUint32(addr *uint32, order MemoryOrderLoad) uint32 {
	switch order {
	case relaxed{}:
		return loadUint32Relaxed(addr)
	case acquire{}:
		return loadUint32Acquire(addr)
	case seqCst{}:
		fallthrough
	default:
		return loadUint32SeqCst(addr)
	}
}

func LoadInt64(addr *int64, order MemoryOrderLoad) int64 {
	switch order {
	case relaxed{}:
		return loadInt64Relaxed(addr)
	case acquire{}:
		return loadInt64Acquire(addr)
	case seqCst{}:
		fallthrough
	default:
		return loadInt64SeqCst(addr)
	}
}

func LoadUint64(addr *uint64, order MemoryOrderLoad) uint64 {
	switch order {
	case relaxed{}:
		return loadUint64Relaxed(addr)
	case acquire{}:
		return loadUint64Acquire(addr)
	case seqCst{}:
		fallthrough
	default:
		return loadUint64SeqCst(addr)
	}
}

func LoadUintptr(addr *uintptr, order MemoryOrderLoad) uintptr {
	switch order {
	case relaxed{}:
		return loadUintptrRelaxed(addr)
	case acquire{}:
		return loadUintptrAcquire(addr)
	case seqCst{}:
		fallthrough
	default:
		return loadUintptrSeqCst(addr)
	}
}

func LoadPointer(addr *unsafe.Pointer, order MemoryOrderLoad) unsafe.Pointer {
	switch order {
	case relaxed{}:
		return loadPointerRelaxed(addr)
	case acquire{}:
		return loadPointerAcquire(addr)
	case seqCst{}:
		fallthrough
	default:
		return loadPointerSeqCst(addr)
	}
}

func StoreInt32(addr *int32, val int32, order MemoryOrderStore) {
	switch order {
	case relaxed{}:
		storeInt32Relaxed(addr, val)
	case release{}:
		storeInt32Release(addr, val)
	case seqCst{}:
		fallthrough
	default:
		storeInt32SeqCst(addr, val)
	}
}

func StoreUint32(addr *uint32, val uint32, order MemoryOrderStore) {
	switch order {
	case relaxed{}:
		storeUint32Relaxed(addr, val)
	case release{}:
		storeUint32Release(addr, val)
	case seqCst{}:
		fallthrough
	default:
		storeUint32SeqCst(addr, val)
	}
}

func StoreInt64(addr *int64, val int64, order MemoryOrderStore) {
	switch order {
	case relaxed{}:
		storeInt64Relaxed(addr, val)
	case release{}:
		storeInt64Release(addr, val)
	case seqCst{}:
		fallthrough
	default:
		storeInt64SeqCst(addr, val)
	}
}

func StoreUint64(addr *uint64, val uint64, order MemoryOrderStore) {
	switch order {
	case relaxed{}:
		storeUint64Relaxed(addr, val)
	case release{}:
		storeUint64Release(addr, val)
	case seqCst{}:
		fallthrough
	default:
		storeUint64SeqCst(addr, val)
	}
}

func StoreUintptr(addr *uintptr, val uintptr, order MemoryOrderStore) {
	switch order {
	case relaxed{}:
		storeUintptrRelaxed(addr, val)
	case release{}:
		storeUintptrRelease(addr, val)
	case seqCst{}:
		fallthrough
	default:
		storeUintptrSeqCst(addr, val)
	}
}

func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer, _ MemoryOrderStore) {
	// GC requires write barrier for storing pointers there is no way to check in non-runtime code is GC active or not.
	// Thus use native atomic.StorePointer without considering order.
	atomic.StorePointer(addr, val)
}
