package atomix

import "unsafe"

func SwapInt32(addr *int32, new int32, order MemoryOrder) (old int32) {
	// todo implement me
	return
}

func SwapUint32(addr *uint32, new uint32, order MemoryOrder) (old uint32) {
	// todo implement me
	return
}

func SwapUintptr(addr *uintptr, new uintptr, order MemoryOrder) (old uintptr) {
	// todo implement me
	return
}

func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer, order MemoryOrder) (old unsafe.Pointer) {
	// todo implement me
	return
}

func CompareAndSwapInt32(addr *int32, old, new int32, order MemoryOrder) (swapped bool) {
	// todo implement me
	return
}

func CompareAndSwapUint32(addr *uint32, old, new uint32, order MemoryOrder) (swapped bool) {
	// todo implement me
	return
}

func CompareAndSwapUintptr(addr *uintptr, old, new uintptr, order MemoryOrder) (swapped bool) {
	// todo implement me
	return
}

func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer, order MemoryOrder) (swapped bool) {
	// todo implement me
	return
}

func AddInt32(addr *int32, delta int32, order MemoryOrder) (new int32) {
	// todo implement me
	return
}

func AddUint32(addr *uint32, delta uint32, order MemoryOrder) (new uint32) {
	// todo implement me
	return
}

func AddUintptr(addr *uintptr, delta uintptr, order MemoryOrder) (new uintptr) {
	// todo implement me
	return
}

func AndInt32(addr *int32, mask int32, order MemoryOrder) (old int32) {
	// todo implement me
	return
}

func AndUint32(addr *uint32, mask uint32, order MemoryOrder) (old uint32) {
	// todo implement me
	return
}

func AndUintptr(addr *uintptr, mask uintptr, order MemoryOrder) (old uintptr) {
	// todo implement me
	return
}

func OrInt32(addr *int32, mask int32, order MemoryOrder) (old int32) {
	// todo implement me
	return
}

func OrUint32(addr *uint32, mask uint32, order MemoryOrder) (old uint32) {
	// todo implement me
	return
}

func OrUintptr(addr *uintptr, mask uintptr, order MemoryOrder) (old uintptr) {
	// todo implement me
	return
}

func LoadInt32(addr *int32, order MemoryOrder) (val int32) {
	// todo implement me
	return
}

func LoadUint32(addr *uint32, order MemoryOrder) (val uint32) {
	// todo implement me
	return
}

func LoadUintptr(addr *uintptr, order MemoryOrder) (val uintptr) {
	// todo implement me
	return
}

func LoadPointer(addr *unsafe.Pointer, order MemoryOrder) (val unsafe.Pointer) {
	// todo implement me
	return
}

func StoreInt32(addr *int32, val int32, order MemoryOrder) {
	// todo implement me
}

func StoreUint32(addr *uint32, val uint32, order MemoryOrder) {
	// todo implement me
}

func StoreUintptr(addr *uintptr, val uintptr, order MemoryOrder) {
	// todo implement me
}

func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer, order MemoryOrder) {
	// todo implement me
}
