package atomix

import "unsafe"

func SwapInt32(addr *int32, new int32, order MemoryOrderFull) (old int32) {
	// todo implement me
	return
}

func SwapUint32(addr *uint32, new uint32, order MemoryOrderFull) (old uint32) {
	// todo implement me
	return
}

func SwapUintptr(addr *uintptr, new uintptr, order MemoryOrderFull) (old uintptr) {
	// todo implement me
	return
}

func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer, order MemoryOrderFull) (old unsafe.Pointer) {
	// todo implement me
	return
}

func CompareAndSwapInt32(addr *int32, old, new int32, order MemoryOrderFull) (swapped bool) {
	// todo implement me
	return
}

func CompareAndSwapUint32(addr *uint32, old, new uint32, order MemoryOrderFull) (swapped bool) {
	// todo implement me
	return
}

func CompareAndSwapUintptr(addr *uintptr, old, new uintptr, order MemoryOrderFull) (swapped bool) {
	// todo implement me
	return
}

func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer, order MemoryOrderFull) (swapped bool) {
	// todo implement me
	return
}

func AddInt32(addr *int32, delta int32, order MemoryOrderFull) (new int32) {
	// todo implement me
	return
}

func AddUint32(addr *uint32, delta uint32, order MemoryOrderFull) (new uint32) {
	// todo implement me
	return
}

func AddUintptr(addr *uintptr, delta uintptr, order MemoryOrderFull) (new uintptr) {
	// todo implement me
	return
}

func AndInt32(addr *int32, mask int32, order MemoryOrderFull) (old int32) {
	// todo implement me
	return
}

func AndUint32(addr *uint32, mask uint32, order MemoryOrderFull) (old uint32) {
	// todo implement me
	return
}

func AndUintptr(addr *uintptr, mask uintptr, order MemoryOrderFull) (old uintptr) {
	// todo implement me
	return
}

func OrInt32(addr *int32, mask int32, order MemoryOrderFull) (old int32) {
	// todo implement me
	return
}

func OrUint32(addr *uint32, mask uint32, order MemoryOrderFull) (old uint32) {
	// todo implement me
	return
}

func OrUintptr(addr *uintptr, mask uintptr, order MemoryOrderFull) (old uintptr) {
	// todo implement me
	return
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
	// todo implement me
}

func StoreUint32(addr *uint32, val uint32, order MemoryOrderStore) {
	// todo implement me
}

func StoreInt64(addr *int64, val int64, order MemoryOrderStore) {
	// todo implement me
}

func StoreUint64(addr *uint64, val uint64, order MemoryOrderStore) {
	// todo implement me
}

func StoreUintptr(addr *uintptr, val uintptr, order MemoryOrderStore) {
	// todo implement me
}

func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer, order MemoryOrderStore) {
	// todo implement me
}
