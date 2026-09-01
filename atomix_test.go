package atomix

import (
	"testing"
	"unsafe"
)

const (
	magic32 = 0xdedbeef
	magic64 = 0xdeddeadbeefbeef
)

var global [1024]byte

func testPointers() []unsafe.Pointer {
	var pointers []unsafe.Pointer
	// globals
	for i := 0; i < 10; i++ {
		pointers = append(pointers, unsafe.Pointer(&global[1<<i-1]))
	}
	// heap
	pointers = append(pointers, unsafe.Pointer(new(byte)))
	// nil
	pointers = append(pointers, nil)
	return pointers
}

func TestLoad(t *testing.T) {
	t.Run("int32", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderLoad) {
			var x struct {
				before int32
				i      int32
				after  int32
			}
			x.before = magic32
			x.after = magic32
			for delta := int32(1); delta+delta > delta; delta += delta {
				k := LoadInt32(&x.i, order)
				if k != x.i {
					t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
				}
				x.i += delta
			}
			if x.before != magic32 || x.after != magic32 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("uint32", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderLoad) {
			var x struct {
				before uint32
				i      uint32
				after  uint32
			}
			x.before = magic32
			x.after = magic32
			for delta := uint32(1); delta+delta > delta; delta += delta {
				k := LoadUint32(&x.i, order)
				if k != x.i {
					t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
				}
				x.i += delta
			}
			if x.before != magic32 || x.after != magic32 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("int64", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderLoad) {
			var x struct {
				before int64
				i      int64
				after  int64
			}
			x.before = magic64
			x.after = magic64
			for delta := int64(1); delta+delta > delta; delta += delta {
				k := LoadInt64(&x.i, order)
				if k != x.i {
					t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
				}
				x.i += delta
			}
			if x.before != magic64 || x.after != magic64 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("uint64", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderLoad) {
			var x struct {
				before uint64
				i      uint64
				after  uint64
			}
			x.before = magic64
			x.after = magic64
			for delta := uint64(1); delta+delta > delta; delta += delta {
				k := LoadUint64(&x.i, order)
				if k != x.i {
					t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
				}
				x.i += delta
			}
			if x.before != magic64 || x.after != magic64 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("uintptr", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderLoad) {
			var x struct {
				before uintptr
				i      uintptr
				after  uintptr
			}
			var m uint64 = magic64
			magicptr := uintptr(m)
			x.before = magicptr
			x.after = magicptr
			for delta := uintptr(1); delta+delta > delta; delta += delta {
				k := LoadUintptr(&x.i, order)
				if k != x.i {
					t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
				}
				x.i += delta
			}
			if x.before != magicptr || x.after != magicptr {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("pointer", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderLoad) {
			var x struct {
				before uintptr
				i      unsafe.Pointer
				after  uintptr
			}
			var m uint64 = magic64
			magicptr := uintptr(m)
			x.before = magicptr
			x.after = magicptr
			for _, p := range testPointers() {
				x.i = p
				k := LoadPointer(&x.i, order)
				if k != p {
					t.Fatalf("p=%x k=%x", p, k)
				}
			}
			if x.before != magicptr || x.after != magicptr {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
}

func TestStore(t *testing.T) {
	t.Run("int32", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderStore) {
			var x struct {
				before int32
				i      int32
				after  int32
			}
			x.before = magic32
			x.after = magic32
			v := int32(0)
			for delta := int32(1); delta+delta > delta; delta += delta {
				StoreInt32(&x.i, v, order)
				if x.i != v {
					t.Fatalf("delta=%d i=%d v=%d", delta, x.i, v)
				}
				v += delta
			}
			if x.before != magic32 || x.after != magic32 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("uint32", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderStore) {
			var x struct {
				before uint32
				i      uint32
				after  uint32
			}
			x.before = magic32
			x.after = magic32
			v := uint32(0)
			for delta := uint32(1); delta+delta > delta; delta += delta {
				StoreUint32(&x.i, v, order)
				if x.i != v {
					t.Fatalf("delta=%d i=%d v=%d", delta, x.i, v)
				}
				v += delta
			}
			if x.before != magic32 || x.after != magic32 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("int64", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderStore) {
			var x struct {
				before int64
				i      int64
				after  int64
			}
			magic64 := int64(magic64)
			x.before = magic64
			x.after = magic64
			v := int64(0)
			for delta := int64(1); delta+delta > delta; delta += delta {
				StoreInt64(&x.i, v, order)
				if x.i != v {
					t.Fatalf("delta=%d i=%d v=%d", delta, x.i, v)
				}
				v += delta
			}
			if x.before != magic64 || x.after != magic64 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("uint64", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderStore) {
			var x struct {
				before uint64
				i      uint64
				after  uint64
			}
			magic64 := uint64(magic64)
			x.before = magic64
			x.after = magic64
			v := uint64(0)
			for delta := uint64(1); delta+delta > delta; delta += delta {
				StoreUint64(&x.i, v, order)
				if x.i != v {
					t.Fatalf("delta=%d i=%d v=%d", delta, x.i, v)
				}
				v += delta
			}
			if x.before != magic64 || x.after != magic64 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("uintptr", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderStore) {
			var x struct {
				before uintptr
				i      uintptr
				after  uintptr
			}
			var m uint64 = magic64
			magicptr := uintptr(m)
			x.before = magicptr
			x.after = magicptr
			v := uintptr(0)
			for delta := uintptr(1); delta+delta > delta; delta += delta {
				StoreUintptr(&x.i, v, order)
				if x.i != v {
					t.Fatalf("delta=%d i=%d v=%d", delta, x.i, v)
				}
				v += delta
			}
			if x.before != magicptr || x.after != magicptr {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("pointer", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderStore) {
			var x struct {
				before uintptr
				i      unsafe.Pointer
				after  uintptr
			}
			var m uint64 = magic64
			magicptr := uintptr(m)
			x.before = magicptr
			x.after = magicptr
			for _, p := range testPointers() {
				StorePointer(&x.i, p, order)
				if x.i != p {
					t.Fatalf("x.i=%p p=%p", x.i, p)
				}
			}
			if x.before != magicptr || x.after != magicptr {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
}

func TestSwap(t *testing.T) {
	t.Run("int32", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderAll) {
			var x struct {
				before int32
				i      int32
				after  int32
			}
			x.before = magic32
			x.after = magic32
			var j int32
			for delta := int32(1); delta+delta > delta; delta += delta {
				k := SwapInt32(&x.i, delta, order)
				if x.i != delta || k != j {
					t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
				}
				j = delta
			}
			if x.before != magic32 || x.after != magic32 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("acq rel", func(t *testing.T) { testfn(t, MemoryOrderAcqRel) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("uint32", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderAll) {
			var x struct {
				before uint32
				i      uint32
				after  uint32
			}
			x.before = magic32
			x.after = magic32
			var j uint32
			for delta := uint32(1); delta+delta > delta; delta += delta {
				k := SwapUint32(&x.i, delta, order)
				if x.i != delta || k != j {
					t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
				}
				j = delta
			}
			if x.before != magic32 || x.after != magic32 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("acq rel", func(t *testing.T) { testfn(t, MemoryOrderAcqRel) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("int64", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderAll) {
			var x struct {
				before int64
				i      int64
				after  int64
			}
			magic64 := int64(magic64)
			x.before = magic64
			x.after = magic64
			var j int64
			for delta := int64(1); delta+delta > delta; delta += delta {
				k := SwapInt64(&x.i, delta, order)
				if x.i != delta || k != j {
					t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
				}
				j = delta
			}
			if x.before != magic64 || x.after != magic64 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("acq rel", func(t *testing.T) { testfn(t, MemoryOrderAcqRel) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("uint64", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderAll) {
			var x struct {
				before uint64
				i      uint64
				after  uint64
			}
			magic64 := uint64(magic64)
			x.before = magic64
			x.after = magic64
			var j uint64
			for delta := uint64(1); delta+delta > delta; delta += delta {
				k := SwapUint64(&x.i, delta, order)
				if x.i != delta || k != j {
					t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
				}
				j = delta
			}
			if x.before != magic64 || x.after != magic64 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("acq rel", func(t *testing.T) { testfn(t, MemoryOrderAcqRel) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("uintptr", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderAll) {
			var x struct {
				before uintptr
				i      uintptr
				after  uintptr
			}
			var m uint64 = magic64
			magicptr := uintptr(m)
			x.before = magicptr
			x.after = magicptr
			var j uintptr
			for delta := uintptr(1); delta+delta > delta; delta += delta {
				k := SwapUintptr(&x.i, delta, order)
				if x.i != delta || k != j {
					t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
				}
				j = delta
			}
			if x.before != magicptr || x.after != magicptr {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("acq rel", func(t *testing.T) { testfn(t, MemoryOrderAcqRel) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("pointer", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderAll) {
			var x struct {
				before uintptr
				i      unsafe.Pointer
				after  uintptr
			}
			var m uint64 = magic64
			magicptr := uintptr(m)
			x.before = magicptr
			x.after = magicptr
			var j unsafe.Pointer

			for _, p := range testPointers() {
				k := SwapPointer(&x.i, p, order)
				if x.i != p || k != j {
					t.Fatalf("p=%p i=%p j=%p k=%p", p, x.i, j, k)
				}
				j = p
			}
			if x.before != magicptr || x.after != magicptr {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("acq rel", func(t *testing.T) { testfn(t, MemoryOrderAcqRel) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
}

func TestCAS(t *testing.T) {
	t.Run("int32", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderAll) {
			var x struct {
				before int32
				i      int32
				after  int32
			}
			x.before = magic32
			x.after = magic32
			for val := int32(1); val+val > val; val += val {
				x.i = val
				if !CompareAndSwapInt32(&x.i, val, val+1, order) {
					t.Fatalf("should have swapped %#x %#x", val, val+1)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
				x.i = val + 1
				if CompareAndSwapInt32(&x.i, val, val+2, order) {
					t.Fatalf("should not have swapped %#x %#x", val, val+2)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
			}
			if x.before != magic32 || x.after != magic32 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("acq rel", func(t *testing.T) { testfn(t, MemoryOrderAcqRel) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("uint32", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderAll) {
			var x struct {
				before uint32
				i      uint32
				after  uint32
			}
			x.before = magic32
			x.after = magic32
			for val := uint32(1); val+val > val; val += val {
				x.i = val
				if !CompareAndSwapUint32(&x.i, val, val+1, order) {
					t.Fatalf("should have swapped %#x %#x", val, val+1)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
				x.i = val + 1
				if CompareAndSwapUint32(&x.i, val, val+2, order) {
					t.Fatalf("should not have swapped %#x %#x", val, val+2)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
			}
			if x.before != magic32 || x.after != magic32 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("acq rel", func(t *testing.T) { testfn(t, MemoryOrderAcqRel) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("int64", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderAll) {
			var x struct {
				before int64
				i      int64
				after  int64
			}
			magic64 := int64(magic64)
			x.before = magic64
			x.after = magic64
			for val := int64(1); val+val > val; val += val {
				x.i = val
				if !CompareAndSwapInt64(&x.i, val, val+1, order) {
					t.Fatalf("should have swapped %#x %#x", val, val+1)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
				x.i = val + 1
				if CompareAndSwapInt64(&x.i, val, val+2, order) {
					t.Fatalf("should not have swapped %#x %#x", val, val+2)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
			}
			if x.before != magic64 || x.after != magic64 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("acq rel", func(t *testing.T) { testfn(t, MemoryOrderAcqRel) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("uint64", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderAll) {
			var x struct {
				before uint64
				i      uint64
				after  uint64
			}
			magic64 := uint64(magic64)
			x.before = magic64
			x.after = magic64
			for val := uint64(1); val+val > val; val += val {
				x.i = val
				if !CompareAndSwapUint64(&x.i, val, val+1, order) {
					t.Fatalf("should have swapped %#x %#x", val, val+1)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
				x.i = val + 1
				if CompareAndSwapUint64(&x.i, val, val+2, order) {
					t.Fatalf("should not have swapped %#x %#x", val, val+2)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
			}
			if x.before != magic64 || x.after != magic64 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("acq rel", func(t *testing.T) { testfn(t, MemoryOrderAcqRel) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("uintptr", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderAll) {
			var x struct {
				before uintptr
				i      uintptr
				after  uintptr
			}
			var m uint64 = magic64
			magicptr := uintptr(m)
			x.before = magicptr
			x.after = magicptr
			for val := uintptr(1); val+val > val; val += val {
				x.i = val
				if !CompareAndSwapUintptr(&x.i, val, val+1, order) {
					t.Fatalf("should have swapped %#x %#x", val, val+1)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
				x.i = val + 1
				if CompareAndSwapUintptr(&x.i, val, val+2, order) {
					t.Fatalf("should not have swapped %#x %#x", val, val+2)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
			}
			if x.before != magicptr || x.after != magicptr {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("acq rel", func(t *testing.T) { testfn(t, MemoryOrderAcqRel) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
	t.Run("pointer", func(t *testing.T) {
		testfn := func(t *testing.T, order MemoryOrderAll) {
			var x struct {
				before uintptr
				i      unsafe.Pointer
				after  uintptr
			}
			var m uint64 = magic64
			magicptr := uintptr(m)
			x.before = magicptr
			x.after = magicptr
			q := unsafe.Pointer(new(byte))
			for _, p := range testPointers() {
				x.i = p
				if !CompareAndSwapPointer(&x.i, p, q, order) {
					t.Fatalf("should have swapped %p %p", p, q)
				}
				if x.i != q {
					t.Fatalf("wrong x.i after swap: x.i=%p want %p", x.i, q)
				}
				if CompareAndSwapPointer(&x.i, p, nil, order) {
					t.Fatalf("should not have swapped %p nil", p)
				}
				if x.i != q {
					t.Fatalf("wrong x.i after swap: x.i=%p want %p", x.i, q)
				}
			}
			if x.before != magicptr || x.after != magicptr {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
			}
		}
		t.Run("relaxed", func(t *testing.T) { testfn(t, MemoryOrderRelaxed) })
		t.Run("acquire", func(t *testing.T) { testfn(t, MemoryOrderAcquire) })
		t.Run("release", func(t *testing.T) { testfn(t, MemoryOrderRelease) })
		t.Run("acq rel", func(t *testing.T) { testfn(t, MemoryOrderAcqRel) })
		t.Run("seq cst", func(t *testing.T) { testfn(t, MemoryOrderSeqCst) })
	})
}
