package atomix

import (
	"testing"
	"unsafe"
)

func BenchmarkLoad(b *testing.B) {
	b.Run("int32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderLoad) {
			var v int32 = 123
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = LoadInt32(&v, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderLoad) {
			var v uint32 = 123
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = LoadUint32(&v, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("int64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderLoad) {
			var v int64 = 123
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = LoadInt64(&v, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderLoad) {
			var v uint64 = 123
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = LoadUint64(&v, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uintptr", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderLoad) {
			var v uintptr = 123
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = LoadUintptr(&v, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("pointer", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderLoad) {
			var v uintptr = 123
			p := unsafe.Pointer(&v)
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = LoadPointer(&p, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
}

func BenchmarkStore(b *testing.B) {
	b.Run("int32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderStore) {
			var v int32
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				StoreInt32(&v, int32(i), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderStore) {
			var v uint32
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				StoreUint32(&v, uint32(i), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("int64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderStore) {
			var v int64
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				StoreInt64(&v, int64(i), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderStore) {
			var v uint64
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				StoreUint64(&v, uint64(i), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uintptr", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderStore) {
			var v uintptr
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				StoreUintptr(&v, uintptr(i), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("pointer", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderStore) {
			var p unsafe.Pointer
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				StorePointer(&p, unsafe.Pointer(&i), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
}

func BenchmarkAdd(b *testing.B) {
	b.Run("int32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int32
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				AddInt32(&v, 1, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint32
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				AddUint32(&v, 1, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("int64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int64
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				AddInt64(&v, 1, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint64
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				AddUint64(&v, 1, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uintptr", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uintptr
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				AddUintptr(&v, 1, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
}

func BenchmarkSwap(b *testing.B) {
	b.Run("int32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int32
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				SwapInt32(&v, int32(i), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint32
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				SwapUint32(&v, uint32(i), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("int64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int64
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				SwapInt64(&v, int64(i), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint64
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				SwapUint64(&v, uint64(i), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uintptr", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uintptr
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				SwapUintptr(&v, uintptr(i), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("pointer", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var p unsafe.Pointer
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				SwapPointer(&p, unsafe.Pointer(&i), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
}

func BenchmarkCAS(b *testing.B) {
	b.Run("int32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int32
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				old := int32(i)
				new_ := old + 1
				CompareAndSwapInt32(&v, old, new_, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint32
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				old := uint32(i)
				new_ := old + 1
				CompareAndSwapUint32(&v, old, new_, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("int64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int64
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				old := int64(i)
				new_ := old + 1
				CompareAndSwapInt64(&v, old, new_, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint64
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				old := uint64(i)
				new_ := old + 1
				CompareAndSwapUint64(&v, old, new_, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uintptr", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uintptr
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				old := uintptr(i)
				new_ := old + 1
				CompareAndSwapUintptr(&v, old, new_, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("pointer", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var p, old, new_ unsafe.Pointer
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				j := i + 1
				old = unsafe.Pointer(&i)
				new_ = unsafe.Pointer(&j)
				CompareAndSwapPointer(&p, old, new_, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
}

func BenchmarkAnd(b *testing.B) {
	b.Run("int32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int32 = -1
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				AndInt32(&v, ^int32(1), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint32 = 0xffffffff
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				AndUint32(&v, ^uint32(1), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("int64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int64 = -1
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				AndInt64(&v, ^int64(1), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint64 = 0xffffffffffffffff
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				AndUint64(&v, ^uint64(1), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uintptr", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uintptr = ^uintptr(0)
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				AndUintptr(&v, ^uintptr(1), order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
}

func BenchmarkOr(b *testing.B) {
	b.Run("int32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int32
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				OrInt32(&v, 1, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint32
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				OrUint32(&v, 1, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("int64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int64
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				OrInt64(&v, 1, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint64
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				OrUint64(&v, 1, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uintptr", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uintptr
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				OrUintptr(&v, 1, order)
			}
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
}

func BenchmarkParallelLoad(b *testing.B) {
	b.Run("int32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderLoad) {
			var v int32 = 123
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = LoadInt32(&v, order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderLoad) {
			var v uint32 = 123
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = LoadUint32(&v, order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("int64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderLoad) {
			var v int64 = 123
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = LoadInt64(&v, order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderLoad) {
			var v uint64 = 123
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = LoadUint64(&v, order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uintptr", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderLoad) {
			var v uintptr = 123
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = LoadUintptr(&v, order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("pointer", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderLoad) {
			var v uintptr = 123
			p := unsafe.Pointer(&v)
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = LoadPointer(&p, order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
}

func BenchmarkParallelStore(b *testing.B) {
	b.Run("int32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderStore) {
			var v int32
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := int32(0)
				for pb.Next() {
					StoreInt32(&v, i, order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderStore) {
			var v uint32
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := uint32(0)
				for pb.Next() {
					StoreUint32(&v, i, order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("int64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderStore) {
			var v int64
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := int64(0)
				for pb.Next() {
					StoreInt64(&v, i, order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderStore) {
			var v uint64
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := uint64(0)
				for pb.Next() {
					StoreUint64(&v, i, order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uintptr", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderStore) {
			var v uintptr
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := uintptr(0)
				for pb.Next() {
					StoreUintptr(&v, i, order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("pointer", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderStore) {
			var p unsafe.Pointer
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := uintptr(0)
				for pb.Next() {
					StorePointer(&p, unsafe.Pointer(&i), order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
}

func BenchmarkParallelAdd(b *testing.B) {
	b.Run("int32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int32
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					AddInt32(&v, 1, order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint32
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					AddUint32(&v, 1, order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("int64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int64
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					AddInt64(&v, 1, order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint64
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					AddUint64(&v, 1, order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uintptr", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uintptr
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					AddUintptr(&v, 1, order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
}

func BenchmarkParallelSwap(b *testing.B) {
	b.Run("int32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int32
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := int32(0)
				for pb.Next() {
					SwapInt32(&v, i, order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint32
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := uint32(0)
				for pb.Next() {
					SwapUint32(&v, i, order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("int64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int64
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := int64(0)
				for pb.Next() {
					SwapInt64(&v, i, order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint64
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := uint64(0)
				for pb.Next() {
					SwapUint64(&v, i, order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uintptr", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uintptr
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := uintptr(0)
				for pb.Next() {
					SwapUintptr(&v, i, order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("pointer", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var p unsafe.Pointer
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := uintptr(0)
				for pb.Next() {
					SwapPointer(&p, unsafe.Pointer(&i), order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
}

func BenchmarkParallelCAS(b *testing.B) {
	b.Run("int32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int32
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := int32(0)
				for pb.Next() {
					old := i
					new_ := old + 1
					CompareAndSwapInt32(&v, old, new_, order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint32
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := uint32(0)
				for pb.Next() {
					old := i
					new_ := old + 1
					CompareAndSwapUint32(&v, old, new_, order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("int64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int64
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := int64(0)
				for pb.Next() {
					old := i
					new_ := old + 1
					CompareAndSwapInt64(&v, old, new_, order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint64
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := uint64(0)
				for pb.Next() {
					old := i
					new_ := old + 1
					CompareAndSwapUint64(&v, old, new_, order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uintptr", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uintptr
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := uintptr(0)
				for pb.Next() {
					old := i
					new_ := old + 1
					CompareAndSwapUintptr(&v, old, new_, order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("pointer", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var p unsafe.Pointer
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := uintptr(0)
				for pb.Next() {
					j := i + 1
					old := unsafe.Pointer(&i)
					new_ := unsafe.Pointer(&j)
					CompareAndSwapPointer(&p, old, new_, order)
					i++
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
}

func BenchmarkParallelAnd(b *testing.B) {
	b.Run("int32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int32 = -1
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					AndInt32(&v, ^int32(1), order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint32 = 0xffffffff
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					AndUint32(&v, ^uint32(1), order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("int64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int64 = -1
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					AndInt64(&v, ^int64(1), order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint64 = 0xffffffffffffffff
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					AndUint64(&v, ^uint64(1), order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uintptr", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v = ^uintptr(0)
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					AndUintptr(&v, ^uintptr(1), order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
}

func BenchmarkParallelOr(b *testing.B) {
	b.Run("int32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int32
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					OrInt32(&v, 1, order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint32", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint32
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					OrUint32(&v, 1, order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("int64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v int64
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					OrInt64(&v, 1, order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uint64", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uint64
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					OrUint64(&v, 1, order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
	b.Run("uintptr", func(b *testing.B) {
		benchfn := func(b *testing.B, order MemoryOrderAll) {
			var v uintptr
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					OrUintptr(&v, 1, order)
				}
			})
		}
		b.Run("relaxed", func(b *testing.B) { benchfn(b, MemoryOrderRelaxed) })
		b.Run("acquire", func(b *testing.B) { benchfn(b, MemoryOrderAcquire) })
		b.Run("release", func(b *testing.B) { benchfn(b, MemoryOrderRelease) })
		b.Run("acq rel", func(b *testing.B) { benchfn(b, MemoryOrderAcqRel) })
		b.Run("seq cst", func(b *testing.B) { benchfn(b, MemoryOrderSeqCst) })
	})
}
