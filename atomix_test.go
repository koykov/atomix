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
		testPointers := func() []unsafe.Pointer {
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
