package pbvector

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestRegistry(t *testing.T) {
	t.Run("detect typeID", func(t *testing.T) {
		type (
			T0 struct {
				X int
				Y float64
			}
			T1 struct {
				Z string
				Q []byte
			}
		)
		var (
			t0, t01 T0
			t1, t02 T1
		)
		tt0 := reflect.TypeOf(t0)
		tt01 := reflect.TypeOf(t01)
		tt1 := reflect.TypeOf(t1)
		tt02 := reflect.TypeOf(t02)
		println(uintptr(unsafe.Pointer(&tt0)))
		println(uintptr(unsafe.Pointer(&tt01)))
		println(uintptr(unsafe.Pointer(&tt1)))
		println(uintptr(unsafe.Pointer(&tt02)))
		t.Logf("%#v", tt0)
		t.Logf("%#v", tt01)
		t.Logf("%#v", tt1)
		t.Logf("%#v", tt02)
	})
}
