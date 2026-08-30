package atomix

import "testing"

const (
	magic32 = 0xdedbeef
	magic64 = 0xdeddeadbeefbeef
)

func TestAtomix(t *testing.T) {
	t.Run("load", func(t *testing.T) {
		t.Run("int32", func(t *testing.T) {
			t.Run("relaxed", func(t *testing.T) {
				var x struct {
					before int32
					i      int32
					after  int32
				}
				x.before = magic32
				x.after = magic32
				for delta := int32(1); delta+delta > delta; delta += delta {
					k := LoadInt32(&x.i, MemoryOrderRelaxed)
					if k != x.i {
						t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
					}
					x.i += delta
				}
				if x.before != magic32 || x.after != magic32 {
					t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
				}
			})
		})
	})
}
