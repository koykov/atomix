package atomix

import (
	"testing"
	"unsafe"
)

func BenchmarkLoad(b *testing.B) {
	types := []struct {
		name string
		fn   func(addr interface{}, order MemoryOrderLoad) interface{}
		addr interface{}
	}{
		{"int32", func(addr interface{}, order MemoryOrderLoad) interface{} {
			return LoadInt32(addr.(*int32), order)
		}, new(int32)},
		{"uint32", func(addr interface{}, order MemoryOrderLoad) interface{} {
			return LoadUint32(addr.(*uint32), order)
		}, new(uint32)},
		{"int64", func(addr interface{}, order MemoryOrderLoad) interface{} {
			return LoadInt64(addr.(*int64), order)
		}, new(int64)},
		{"uint64", func(addr interface{}, order MemoryOrderLoad) interface{} {
			return LoadUint64(addr.(*uint64), order)
		}, new(uint64)},
		{"uintptr", func(addr interface{}, order MemoryOrderLoad) interface{} {
			return LoadUintptr(addr.(*uintptr), order)
		}, new(uintptr)},
		{"pointer", func(addr interface{}, order MemoryOrderLoad) interface{} {
			return LoadPointer(addr.(*unsafe.Pointer), order)
		}, new(unsafe.Pointer)},
	}

	orders := []struct {
		name  string
		order MemoryOrderLoad
	}{
		{"relaxed", MemoryOrderRelaxed},
		{"acquire", MemoryOrderAcquire},
		{"seq cst", MemoryOrderSeqCst},
	}

	for _, typ := range types {
		b.Run(typ.name, func(b *testing.B) {
			for _, ord := range orders {
				b.Run(ord.name, func(b *testing.B) {
					b.ReportAllocs()
					for i := 0; i < b.N; i++ {
						typ.fn(typ.addr, ord.order)
					}
				})
			}
		})
	}
}

func BenchmarkStore(b *testing.B) {
	types := []struct {
		name string
		fn   func(addr interface{}, val interface{}, order MemoryOrderStore)
		addr interface{}
		val  interface{}
	}{
		{"int32", func(addr interface{}, val interface{}, order MemoryOrderStore) {
			StoreInt32(addr.(*int32), val.(int32), order)
		}, new(int32), int32(0)},
		{"uint32", func(addr interface{}, val interface{}, order MemoryOrderStore) {
			StoreUint32(addr.(*uint32), val.(uint32), order)
		}, new(uint32), uint32(0)},
		{"int64", func(addr interface{}, val interface{}, order MemoryOrderStore) {
			StoreInt64(addr.(*int64), val.(int64), order)
		}, new(int64), int64(0)},
		{"uint64", func(addr interface{}, val interface{}, order MemoryOrderStore) {
			StoreUint64(addr.(*uint64), val.(uint64), order)
		}, new(uint64), uint64(0)},
		{"uintptr", func(addr interface{}, val interface{}, order MemoryOrderStore) {
			StoreUintptr(addr.(*uintptr), val.(uintptr), order)
		}, new(uintptr), uintptr(0)},
		{"pointer", func(addr interface{}, val interface{}, order MemoryOrderStore) {
			StorePointer(addr.(*unsafe.Pointer), val.(unsafe.Pointer), order)
		}, new(unsafe.Pointer), unsafe.Pointer(nil)},
	}

	orders := []struct {
		name  string
		order MemoryOrderStore
	}{
		{"relaxed", MemoryOrderRelaxed},
		{"release", MemoryOrderRelease},
		{"seq cst", MemoryOrderSeqCst},
	}

	for _, typ := range types {
		b.Run(typ.name, func(b *testing.B) {
			for _, ord := range orders {
				b.Run(ord.name, func(b *testing.B) {
					b.ReportAllocs()
					for i := 0; i < b.N; i++ {
						typ.fn(typ.addr, typ.val, ord.order)
					}
				})
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	types := []struct {
		name  string
		fn    func(addr interface{}, delta interface{}, order MemoryOrderAll) interface{}
		addr  interface{}
		delta interface{}
	}{
		{"int32", func(addr interface{}, delta interface{}, order MemoryOrderAll) interface{} {
			return AddInt32(addr.(*int32), delta.(int32), order)
		}, new(int32), int32(1)},
		{"uint32", func(addr interface{}, delta interface{}, order MemoryOrderAll) interface{} {
			return AddUint32(addr.(*uint32), delta.(uint32), order)
		}, new(uint32), uint32(1)},
		{"int64", func(addr interface{}, delta interface{}, order MemoryOrderAll) interface{} {
			return AddInt64(addr.(*int64), delta.(int64), order)
		}, new(int64), int64(1)},
		{"uint64", func(addr interface{}, delta interface{}, order MemoryOrderAll) interface{} {
			return AddUint64(addr.(*uint64), delta.(uint64), order)
		}, new(uint64), uint64(1)},
		{"uintptr", func(addr interface{}, delta interface{}, order MemoryOrderAll) interface{} {
			return AddUintptr(addr.(*uintptr), delta.(uintptr), order)
		}, new(uintptr), uintptr(1)},
	}

	orders := []struct {
		name  string
		order MemoryOrderAll
	}{
		{"relaxed", MemoryOrderRelaxed},
		{"acquire", MemoryOrderAcquire},
		{"release", MemoryOrderRelease},
		{"acq rel", MemoryOrderAcqRel},
		{"seq cst", MemoryOrderSeqCst},
	}

	for _, typ := range types {
		b.Run(typ.name, func(b *testing.B) {
			for _, ord := range orders {
				b.Run(ord.name, func(b *testing.B) {
					b.ReportAllocs()
					for i := 0; i < b.N; i++ {
						typ.fn(typ.addr, typ.delta, ord.order)
					}
				})
			}
		})
	}
}

func BenchmarkSwap(b *testing.B) {
	types := []struct {
		name string
		fn   func(addr interface{}, new interface{}, order MemoryOrderAll) interface{}
		addr interface{}
		new  interface{}
	}{
		{"int32", func(addr interface{}, new interface{}, order MemoryOrderAll) interface{} {
			return SwapInt32(addr.(*int32), new.(int32), order)
		}, new(int32), int32(0)},
		{"uint32", func(addr interface{}, new interface{}, order MemoryOrderAll) interface{} {
			return SwapUint32(addr.(*uint32), new.(uint32), order)
		}, new(uint32), uint32(0)},
		{"int64", func(addr interface{}, new interface{}, order MemoryOrderAll) interface{} {
			return SwapInt64(addr.(*int64), new.(int64), order)
		}, new(int64), int64(0)},
		{"uint64", func(addr interface{}, new interface{}, order MemoryOrderAll) interface{} {
			return SwapUint64(addr.(*uint64), new.(uint64), order)
		}, new(uint64), uint64(0)},
		{"uintptr", func(addr interface{}, new interface{}, order MemoryOrderAll) interface{} {
			return SwapUintptr(addr.(*uintptr), new.(uintptr), order)
		}, new(uintptr), uintptr(0)},
		{"pointer", func(addr interface{}, new interface{}, order MemoryOrderAll) interface{} {
			return SwapPointer(addr.(*unsafe.Pointer), new.(unsafe.Pointer), order)
		}, new(unsafe.Pointer), unsafe.Pointer(nil)},
	}

	orders := []struct {
		name  string
		order MemoryOrderAll
	}{
		{"relaxed", MemoryOrderRelaxed},
		{"acquire", MemoryOrderAcquire},
		{"release", MemoryOrderRelease},
		{"acq rel", MemoryOrderAcqRel},
		{"seq cst", MemoryOrderSeqCst},
	}

	for _, typ := range types {
		b.Run(typ.name, func(b *testing.B) {
			for _, ord := range orders {
				b.Run(ord.name, func(b *testing.B) {
					b.ReportAllocs()
					for i := 0; i < b.N; i++ {
						typ.fn(typ.addr, typ.new, ord.order)
					}
				})
			}
		})
	}
}

func BenchmarkCAS(b *testing.B) {
	types := []struct {
		name string
		fn   func(addr interface{}, old, new interface{}, order MemoryOrderAll) bool
		addr interface{}
		old  interface{}
		new  interface{}
	}{
		{"int32", func(addr interface{}, old, new interface{}, order MemoryOrderAll) bool {
			return CompareAndSwapInt32(addr.(*int32), old.(int32), new.(int32), order)
		}, new(int32), int32(0), int32(1)},
		{"uint32", func(addr interface{}, old, new interface{}, order MemoryOrderAll) bool {
			return CompareAndSwapUint32(addr.(*uint32), old.(uint32), new.(uint32), order)
		}, new(uint32), uint32(0), uint32(1)},
		{"int64", func(addr interface{}, old, new interface{}, order MemoryOrderAll) bool {
			return CompareAndSwapInt64(addr.(*int64), old.(int64), new.(int64), order)
		}, new(int64), int64(0), int64(1)},
		{"uint64", func(addr interface{}, old, new interface{}, order MemoryOrderAll) bool {
			return CompareAndSwapUint64(addr.(*uint64), old.(uint64), new.(uint64), order)
		}, new(uint64), uint64(0), uint64(1)},
		{"uintptr", func(addr interface{}, old, new interface{}, order MemoryOrderAll) bool {
			return CompareAndSwapUintptr(addr.(*uintptr), old.(uintptr), new.(uintptr), order)
		}, new(uintptr), uintptr(0), uintptr(1)},
		{"pointer", func(addr interface{}, old, new interface{}, order MemoryOrderAll) bool {
			return CompareAndSwapPointer(addr.(*unsafe.Pointer), old.(unsafe.Pointer), new.(unsafe.Pointer), order)
		}, new(unsafe.Pointer), unsafe.Pointer(nil), unsafe.Pointer(new(byte))},
	}

	orders := []struct {
		name  string
		order MemoryOrderAll
	}{
		{"relaxed", MemoryOrderRelaxed},
		{"acquire", MemoryOrderAcquire},
		{"release", MemoryOrderRelease},
		{"acq rel", MemoryOrderAcqRel},
		{"seq cst", MemoryOrderSeqCst},
	}

	for _, typ := range types {
		b.Run(typ.name, func(b *testing.B) {
			for _, ord := range orders {
				b.Run(ord.name, func(b *testing.B) {
					b.ReportAllocs()
					for i := 0; i < b.N; i++ {
						typ.fn(typ.addr, typ.old, typ.new, ord.order)
					}
				})
			}
		})
	}
}

func BenchmarkAnd(b *testing.B) {
	types := []struct {
		name string
		fn   func(addr interface{}, val interface{}, order MemoryOrderAll) interface{}
		addr interface{}
		val  interface{}
	}{
		{"int32", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return AndInt32(addr.(*int32), val.(int32), order)
		}, func() *int32 { x := int32(-1); return &x }(), int32(^1)},
		{"uint32", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return AndUint32(addr.(*uint32), val.(uint32), order)
		}, func() *uint32 { x := uint32(0xffffffff); return &x }(), uint32(^uint32(1))},
		{"int64", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return AndInt64(addr.(*int64), val.(int64), order)
		}, func() *int64 { x := int64(-1); return &x }(), int64(^1)},
		{"uint64", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return AndUint64(addr.(*uint64), val.(uint64), order)
		}, func() *uint64 { x := uint64(0xffffffffffffffff); return &x }(), uint64(^uint64(1))},
		{"uintptr", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return AndUintptr(addr.(*uintptr), val.(uintptr), order)
		}, func() *uintptr { x := ^uintptr(0); return &x }(), ^uintptr(1)},
	}

	orders := []struct {
		name  string
		order MemoryOrderAll
	}{
		{"relaxed", MemoryOrderRelaxed},
		{"acquire", MemoryOrderAcquire},
		{"release", MemoryOrderRelease},
		{"acq rel", MemoryOrderAcqRel},
		{"seq cst", MemoryOrderSeqCst},
	}

	for _, typ := range types {
		b.Run(typ.name, func(b *testing.B) {
			for _, ord := range orders {
				b.Run(ord.name, func(b *testing.B) {
					b.ReportAllocs()
					for i := 0; i < b.N; i++ {
						typ.fn(typ.addr, typ.val, ord.order)
					}
				})
			}
		})
	}
}

func BenchmarkOr(b *testing.B) {
	types := []struct {
		name string
		fn   func(addr interface{}, val interface{}, order MemoryOrderAll) interface{}
		addr interface{}
		val  interface{}
	}{
		{"int32", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return OrInt32(addr.(*int32), val.(int32), order)
		}, new(int32), int32(1)},
		{"uint32", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return OrUint32(addr.(*uint32), val.(uint32), order)
		}, new(uint32), uint32(1)},
		{"int64", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return OrInt64(addr.(*int64), val.(int64), order)
		}, new(int64), int64(1)},
		{"uint64", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return OrUint64(addr.(*uint64), val.(uint64), order)
		}, new(uint64), uint64(1)},
		{"uintptr", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return OrUintptr(addr.(*uintptr), val.(uintptr), order)
		}, new(uintptr), uintptr(1)},
	}

	orders := []struct {
		name  string
		order MemoryOrderAll
	}{
		{"relaxed", MemoryOrderRelaxed},
		{"acquire", MemoryOrderAcquire},
		{"release", MemoryOrderRelease},
		{"acq rel", MemoryOrderAcqRel},
		{"seq cst", MemoryOrderSeqCst},
	}

	for _, typ := range types {
		b.Run(typ.name, func(b *testing.B) {
			for _, ord := range orders {
				b.Run(ord.name, func(b *testing.B) {
					b.ReportAllocs()
					for i := 0; i < b.N; i++ {
						typ.fn(typ.addr, typ.val, ord.order)
					}
				})
			}
		})
	}
}

func BenchmarkParallelLoad(b *testing.B) {
	types := []struct {
		name string
		fn   func(addr interface{}, order MemoryOrderLoad) interface{}
		addr interface{}
	}{
		{"int32", func(addr interface{}, order MemoryOrderLoad) interface{} {
			return LoadInt32(addr.(*int32), order)
		}, new(int32)},
		{"uint32", func(addr interface{}, order MemoryOrderLoad) interface{} {
			return LoadUint32(addr.(*uint32), order)
		}, new(uint32)},
		{"int64", func(addr interface{}, order MemoryOrderLoad) interface{} {
			return LoadInt64(addr.(*int64), order)
		}, new(int64)},
		{"uint64", func(addr interface{}, order MemoryOrderLoad) interface{} {
			return LoadUint64(addr.(*uint64), order)
		}, new(uint64)},
		{"uintptr", func(addr interface{}, order MemoryOrderLoad) interface{} {
			return LoadUintptr(addr.(*uintptr), order)
		}, new(uintptr)},
		{"pointer", func(addr interface{}, order MemoryOrderLoad) interface{} {
			return LoadPointer(addr.(*unsafe.Pointer), order)
		}, new(unsafe.Pointer)},
	}

	orders := []struct {
		name  string
		order MemoryOrderLoad
	}{
		{"relaxed", MemoryOrderRelaxed},
		{"acquire", MemoryOrderAcquire},
		{"seq cst", MemoryOrderSeqCst},
	}

	for _, typ := range types {
		b.Run(typ.name, func(b *testing.B) {
			for _, ord := range orders {
				b.Run(ord.name, func(b *testing.B) {
					b.ReportAllocs()
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							typ.fn(typ.addr, ord.order)
						}
					})
				})
			}
		})
	}
}

func BenchmarkParallelStore(b *testing.B) {
	types := []struct {
		name string
		fn   func(addr interface{}, val interface{}, order MemoryOrderStore)
		addr interface{}
		val  interface{}
	}{
		{"int32", func(addr interface{}, val interface{}, order MemoryOrderStore) {
			StoreInt32(addr.(*int32), val.(int32), order)
		}, new(int32), int32(0)},
		{"uint32", func(addr interface{}, val interface{}, order MemoryOrderStore) {
			StoreUint32(addr.(*uint32), val.(uint32), order)
		}, new(uint32), uint32(0)},
		{"int64", func(addr interface{}, val interface{}, order MemoryOrderStore) {
			StoreInt64(addr.(*int64), val.(int64), order)
		}, new(int64), int64(0)},
		{"uint64", func(addr interface{}, val interface{}, order MemoryOrderStore) {
			StoreUint64(addr.(*uint64), val.(uint64), order)
		}, new(uint64), uint64(0)},
		{"uintptr", func(addr interface{}, val interface{}, order MemoryOrderStore) {
			StoreUintptr(addr.(*uintptr), val.(uintptr), order)
		}, new(uintptr), uintptr(0)},
		{"pointer", func(addr interface{}, val interface{}, order MemoryOrderStore) {
			StorePointer(addr.(*unsafe.Pointer), val.(unsafe.Pointer), order)
		}, new(unsafe.Pointer), unsafe.Pointer(nil)},
	}

	orders := []struct {
		name  string
		order MemoryOrderStore
	}{
		{"relaxed", MemoryOrderRelaxed},
		{"release", MemoryOrderRelease},
		{"seq cst", MemoryOrderSeqCst},
	}

	for _, typ := range types {
		b.Run(typ.name, func(b *testing.B) {
			for _, ord := range orders {
				b.Run(ord.name, func(b *testing.B) {
					b.ReportAllocs()
					b.RunParallel(func(pb *testing.PB) {
						i := 0
						for pb.Next() {
							typ.fn(typ.addr, typ.val, ord.order)
							i++
						}
					})
				})
			}
		})
	}
}

func BenchmarkParallelAdd(b *testing.B) {
	types := []struct {
		name  string
		fn    func(addr interface{}, delta interface{}, order MemoryOrderAll) interface{}
		addr  interface{}
		delta interface{}
	}{
		{"int32", func(addr interface{}, delta interface{}, order MemoryOrderAll) interface{} {
			return AddInt32(addr.(*int32), delta.(int32), order)
		}, new(int32), int32(1)},
		{"uint32", func(addr interface{}, delta interface{}, order MemoryOrderAll) interface{} {
			return AddUint32(addr.(*uint32), delta.(uint32), order)
		}, new(uint32), uint32(1)},
		{"int64", func(addr interface{}, delta interface{}, order MemoryOrderAll) interface{} {
			return AddInt64(addr.(*int64), delta.(int64), order)
		}, new(int64), int64(1)},
		{"uint64", func(addr interface{}, delta interface{}, order MemoryOrderAll) interface{} {
			return AddUint64(addr.(*uint64), delta.(uint64), order)
		}, new(uint64), uint64(1)},
		{"uintptr", func(addr interface{}, delta interface{}, order MemoryOrderAll) interface{} {
			return AddUintptr(addr.(*uintptr), delta.(uintptr), order)
		}, new(uintptr), uintptr(1)},
	}

	orders := []struct {
		name  string
		order MemoryOrderAll
	}{
		{"relaxed", MemoryOrderRelaxed},
		{"acquire", MemoryOrderAcquire},
		{"release", MemoryOrderRelease},
		{"acq rel", MemoryOrderAcqRel},
		{"seq cst", MemoryOrderSeqCst},
	}

	for _, typ := range types {
		b.Run(typ.name, func(b *testing.B) {
			for _, ord := range orders {
				b.Run(ord.name, func(b *testing.B) {
					b.ReportAllocs()
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							typ.fn(typ.addr, typ.delta, ord.order)
						}
					})
				})
			}
		})
	}
}

func BenchmarkParallelSwap(b *testing.B) {
	types := []struct {
		name string
		fn   func(addr interface{}, new interface{}, order MemoryOrderAll) interface{}
		addr interface{}
		new  interface{}
	}{
		{"int32", func(addr interface{}, new interface{}, order MemoryOrderAll) interface{} {
			return SwapInt32(addr.(*int32), new.(int32), order)
		}, new(int32), int32(0)},
		{"uint32", func(addr interface{}, new interface{}, order MemoryOrderAll) interface{} {
			return SwapUint32(addr.(*uint32), new.(uint32), order)
		}, new(uint32), uint32(0)},
		{"int64", func(addr interface{}, new interface{}, order MemoryOrderAll) interface{} {
			return SwapInt64(addr.(*int64), new.(int64), order)
		}, new(int64), int64(0)},
		{"uint64", func(addr interface{}, new interface{}, order MemoryOrderAll) interface{} {
			return SwapUint64(addr.(*uint64), new.(uint64), order)
		}, new(uint64), uint64(0)},
		{"uintptr", func(addr interface{}, new interface{}, order MemoryOrderAll) interface{} {
			return SwapUintptr(addr.(*uintptr), new.(uintptr), order)
		}, new(uintptr), uintptr(0)},
		{"pointer", func(addr interface{}, new interface{}, order MemoryOrderAll) interface{} {
			return SwapPointer(addr.(*unsafe.Pointer), new.(unsafe.Pointer), order)
		}, new(unsafe.Pointer), unsafe.Pointer(nil)},
	}

	orders := []struct {
		name  string
		order MemoryOrderAll
	}{
		{"relaxed", MemoryOrderRelaxed},
		{"acquire", MemoryOrderAcquire},
		{"release", MemoryOrderRelease},
		{"acq rel", MemoryOrderAcqRel},
		{"seq cst", MemoryOrderSeqCst},
	}

	for _, typ := range types {
		b.Run(typ.name, func(b *testing.B) {
			for _, ord := range orders {
				b.Run(ord.name, func(b *testing.B) {
					b.ReportAllocs()
					b.RunParallel(func(pb *testing.PB) {
						i := 0
						for pb.Next() {
							typ.fn(typ.addr, typ.new, ord.order)
							i++
						}
					})
				})
			}
		})
	}
}

func BenchmarkParallelCAS(b *testing.B) {
	types := []struct {
		name string
		fn   func(addr interface{}, old, new interface{}, order MemoryOrderAll) bool
		addr interface{}
		old  interface{}
		new  interface{}
	}{
		{"int32", func(addr interface{}, old, new interface{}, order MemoryOrderAll) bool {
			return CompareAndSwapInt32(addr.(*int32), old.(int32), new.(int32), order)
		}, new(int32), int32(0), int32(1)},
		{"uint32", func(addr interface{}, old, new interface{}, order MemoryOrderAll) bool {
			return CompareAndSwapUint32(addr.(*uint32), old.(uint32), new.(uint32), order)
		}, new(uint32), uint32(0), uint32(1)},
		{"int64", func(addr interface{}, old, new interface{}, order MemoryOrderAll) bool {
			return CompareAndSwapInt64(addr.(*int64), old.(int64), new.(int64), order)
		}, new(int64), int64(0), int64(1)},
		{"uint64", func(addr interface{}, old, new interface{}, order MemoryOrderAll) bool {
			return CompareAndSwapUint64(addr.(*uint64), old.(uint64), new.(uint64), order)
		}, new(uint64), uint64(0), uint64(1)},
		{"uintptr", func(addr interface{}, old, new interface{}, order MemoryOrderAll) bool {
			return CompareAndSwapUintptr(addr.(*uintptr), old.(uintptr), new.(uintptr), order)
		}, new(uintptr), uintptr(0), uintptr(1)},
		{"pointer", func(addr interface{}, old, new interface{}, order MemoryOrderAll) bool {
			return CompareAndSwapPointer(addr.(*unsafe.Pointer), old.(unsafe.Pointer), new.(unsafe.Pointer), order)
		}, new(unsafe.Pointer), unsafe.Pointer(nil), unsafe.Pointer(new(byte))},
	}

	orders := []struct {
		name  string
		order MemoryOrderAll
	}{
		{"relaxed", MemoryOrderRelaxed},
		{"acquire", MemoryOrderAcquire},
		{"release", MemoryOrderRelease},
		{"acq rel", MemoryOrderAcqRel},
		{"seq cst", MemoryOrderSeqCst},
	}

	for _, typ := range types {
		b.Run(typ.name, func(b *testing.B) {
			for _, ord := range orders {
				b.Run(ord.name, func(b *testing.B) {
					b.ReportAllocs()
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							typ.fn(typ.addr, typ.old, typ.new, ord.order)
						}
					})
				})
			}
		})
	}
}

func BenchmarkParallelAnd(b *testing.B) {
	types := []struct {
		name string
		fn   func(addr interface{}, val interface{}, order MemoryOrderAll) interface{}
		addr interface{}
		val  interface{}
	}{
		{"int32", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return AndInt32(addr.(*int32), val.(int32), order)
		}, func() *int32 { x := int32(-1); return &x }(), int32(^1)},
		{"uint32", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return AndUint32(addr.(*uint32), val.(uint32), order)
		}, func() *uint32 { x := uint32(0xffffffff); return &x }(), uint32(^uint32(1))},
		{"int64", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return AndInt64(addr.(*int64), val.(int64), order)
		}, func() *int64 { x := int64(-1); return &x }(), int64(^1)},
		{"uint64", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return AndUint64(addr.(*uint64), val.(uint64), order)
		}, func() *uint64 { x := uint64(0xffffffffffffffff); return &x }(), uint64(^uint64(1))},
		{"uintptr", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return AndUintptr(addr.(*uintptr), val.(uintptr), order)
		}, func() *uintptr { x := ^uintptr(0); return &x }(), ^uintptr(1)},
	}

	orders := []struct {
		name  string
		order MemoryOrderAll
	}{
		{"relaxed", MemoryOrderRelaxed},
		{"acquire", MemoryOrderAcquire},
		{"release", MemoryOrderRelease},
		{"acq rel", MemoryOrderAcqRel},
		{"seq cst", MemoryOrderSeqCst},
	}

	for _, typ := range types {
		b.Run(typ.name, func(b *testing.B) {
			for _, ord := range orders {
				b.Run(ord.name, func(b *testing.B) {
					b.ReportAllocs()
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							typ.fn(typ.addr, typ.val, ord.order)
						}
					})
				})
			}
		})
	}
}

func BenchmarkParallelOr(b *testing.B) {
	types := []struct {
		name string
		fn   func(addr interface{}, val interface{}, order MemoryOrderAll) interface{}
		addr interface{}
		val  interface{}
	}{
		{"int32", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return OrInt32(addr.(*int32), val.(int32), order)
		}, new(int32), int32(1)},
		{"uint32", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return OrUint32(addr.(*uint32), val.(uint32), order)
		}, new(uint32), uint32(1)},
		{"int64", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return OrInt64(addr.(*int64), val.(int64), order)
		}, new(int64), int64(1)},
		{"uint64", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return OrUint64(addr.(*uint64), val.(uint64), order)
		}, new(uint64), uint64(1)},
		{"uintptr", func(addr interface{}, val interface{}, order MemoryOrderAll) interface{} {
			return OrUintptr(addr.(*uintptr), val.(uintptr), order)
		}, new(uintptr), uintptr(1)},
	}

	orders := []struct {
		name  string
		order MemoryOrderAll
	}{
		{"relaxed", MemoryOrderRelaxed},
		{"acquire", MemoryOrderAcquire},
		{"release", MemoryOrderRelease},
		{"acq rel", MemoryOrderAcqRel},
		{"seq cst", MemoryOrderSeqCst},
	}

	for _, typ := range types {
		b.Run(typ.name, func(b *testing.B) {
			for _, ord := range orders {
				b.Run(ord.name, func(b *testing.B) {
					b.ReportAllocs()
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							typ.fn(typ.addr, typ.val, ord.order)
						}
					})
				})
			}
		})
	}
}
