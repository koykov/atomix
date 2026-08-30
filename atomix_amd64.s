//go:build !race

#include "textflag.h"

// ============================================================
// Memory Order definitions (matching C++ enum)
// ============================================================
// 0: Relaxed
// 1: Consume (skipped - deprecated)
// 2: Acquire
// 3: Release
// 4: AcqRel
// 5: SeqCst

// ============================================================
// LOAD operations
// ============================================================

// LoadRelaxed - no barriers, just load
TEXT ·LoadInt32Relaxed(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    (BX), AX
    MOVL    AX, ret+8(FP)
    RET

TEXT ·LoadUint32Relaxed(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    (BX), AX
    MOVL    AX, ret+8(FP)
    RET

TEXT ·LoadInt64Relaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    (BX), AX
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·LoadUint64Relaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    (BX), AX
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·LoadUintptrRelaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    (BX), AX
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·LoadPointerRelaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    (BX), AX
    MOVQ    AX, ret+8(FP)
    RET

// LoadAcquire - LFENCE guarantees next reads will not reorder till load
TEXT ·LoadInt32Acquire(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    (BX), AX
    LFENCE
    MOVL    AX, ret+8(FP)
    RET

TEXT ·LoadUint32Acquire(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    (BX), AX
    LFENCE
    MOVL    AX, ret+8(FP)
    RET

TEXT ·LoadInt64Acquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    (BX), AX
    LFENCE
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·LoadUint64Acquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    (BX), AX
    LFENCE
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·LoadUintptrAcquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    (BX), AX
    LFENCE
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·LoadPointerAcquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    (BX), AX
    LFENCE
    MOVQ    AX, ret+8(FP)
    RET

// LoadSeqCst - full barrier
TEXT ·LoadInt32SeqCst(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MFENCE
    MOVL    (BX), AX
    MFENCE
    MOVL    AX, ret+8(FP)
    RET

TEXT ·LoadUint32SeqCst(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MFENCE
    MOVL    (BX), AX
    MFENCE
    MOVL    AX, ret+8(FP)
    RET

TEXT ·LoadInt64SeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MFENCE
    MOVQ    (BX), AX
    MFENCE
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·LoadUint64SeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MFENCE
    MOVQ    (BX), AX
    MFENCE
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·LoadUintptrSeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MFENCE
    MOVQ    (BX), AX
    MFENCE
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·LoadPointerSeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MFENCE
    MOVQ    (BX), AX
    MFENCE
    MOVQ    AX, ret+8(FP)
    RET

// ============================================================
// STORE operations
// ============================================================

// StoreRelaxed - no barriers
TEXT ·StoreInt32Relaxed(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    val+8(FP), AX
    MOVL    AX, (BX)
    RET

TEXT ·StoreUint32Relaxed(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    val+8(FP), AX
    MOVL    AX, (BX)
    RET

TEXT ·StoreInt64Relaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    MOVQ    AX, (BX)
    RET

TEXT ·StoreUint64Relaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    MOVQ    AX, (BX)
    RET

TEXT ·StoreUintptrRelaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    MOVQ    AX, (BX)
    RET

// StoreRelease - SFENCE guarantees previous writes visible till call
TEXT ·StoreInt32Release(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    val+8(FP), AX
    SFENCE
    MOVL    AX, (BX)
    RET

TEXT ·StoreUint32Release(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    val+8(FP), AX
    SFENCE
    MOVL    AX, (BX)
    RET

TEXT ·StoreInt64Release(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    SFENCE
    MOVQ    AX, (BX)
    RET

TEXT ·StoreUint64Release(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    SFENCE
    MOVQ    AX, (BX)
    RET

TEXT ·StoreUintptrRelease(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    SFENCE
    MOVQ    AX, (BX)
    RET

// StoreSeqCst - full barrier
TEXT ·StoreInt32SeqCst(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    val+8(FP), AX
    MFENCE
    MOVL    AX, (BX)
    MFENCE
    RET

TEXT ·StoreUint32SeqCst(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    val+8(FP), AX
    MFENCE
    MOVL    AX, (BX)
    MFENCE
    RET

TEXT ·StoreInt64SeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    MFENCE
    MOVQ    AX, (BX)
    MFENCE
    RET

TEXT ·StoreUint64SeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    MFENCE
    MOVQ    AX, (BX)
    MFENCE
    RET

TEXT ·StoreUintptrSeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    MFENCE
    MOVQ    AX, (BX)
    MFENCE
    RET

// ============================================================
// CAS (Compare And Swap) operations
// ============================================================

// CasRelaxed - no barriers
TEXT ·CasInt32Relaxed(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    LOCK
    CMPXCHGL    CX, (BX)
    SETEQ   ret+16(FP)
    RET

TEXT ·CasUint32Relaxed(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    LOCK
    CMPXCHGL    CX, (BX)
    SETEQ   ret+16(FP)
    RET

TEXT ·CasInt64Relaxed(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    LOCK
    CMPXCHGQ    CX, (BX)
    SETEQ   ret+24(FP)
    RET

TEXT ·CasUint64Relaxed(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    LOCK
    CMPXCHGQ    CX, (BX)
    SETEQ   ret+24(FP)
    RET

TEXT ·CasUintptrRelaxed(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    LOCK
    CMPXCHGQ    CX, (BX)
    SETEQ   ret+24(FP)
    RET

// CasAcquire - CAS with acquire semantic(barrier after sucessful operation)
TEXT ·CasInt32Acquire(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    LOCK
    CMPXCHGL    CX, (BX)
    LFENCE
    SETEQ   ret+16(FP)
    RET

TEXT ·CasUint32Acquire(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    LOCK
    CMPXCHGL    CX, (BX)
    LFENCE
    SETEQ   ret+16(FP)
    RET

TEXT ·CasInt64Acquire(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    LOCK
    CMPXCHGQ    CX, (BX)
    LFENCE
    SETEQ   ret+24(FP)
    RET

TEXT ·CasUint64Acquire(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    LOCK
    CMPXCHGQ    CX, (BX)
    LFENCE
    SETEQ   ret+24(FP)
    RET

TEXT ·CasUintptrAcquire(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    LOCK
    CMPXCHGQ    CX, (BX)
    LFENCE
    SETEQ   ret+24(FP)
    RET

// CasRelease - CAS with release seantic (barrier before operation)
TEXT ·CasInt32Release(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    SFENCE
    LOCK
    CMPXCHGL    CX, (BX)
    SETEQ   ret+16(FP)
    RET

TEXT ·CasUint32Release(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    SFENCE
    LOCK
    CMPXCHGL    CX, (BX)
    SETEQ   ret+16(FP)
    RET

TEXT ·CasInt64Release(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    SFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    SETEQ   ret+24(FP)
    RET

TEXT ·CasUint64Release(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    SFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    SETEQ   ret+24(FP)
    RET

TEXT ·CasUintptrRelease(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    SFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    SETEQ   ret+24(FP)
    RET

// CasAcqRel - full barrier (both direction)
TEXT ·CasInt32AcqRel(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    MFENCE
    LOCK
    CMPXCHGL    CX, (BX)
    MFENCE
    SETEQ   ret+16(FP)
    RET

TEXT ·CasUint32AcqRel(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    MFENCE
    LOCK
    CMPXCHGL    CX, (BX)
    MFENCE
    SETEQ   ret+16(FP)
    RET

TEXT ·CasInt64AcqRel(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    MFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    MFENCE
    SETEQ   ret+24(FP)
    RET

TEXT ·CasUint64AcqRel(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    MFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    MFENCE
    SETEQ   ret+24(FP)
    RET

TEXT ·CasUintptrAcqRel(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    MFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    MFENCE
    SETEQ   ret+24(FP)
    RET

// CasSeqCst - full barrier + global order
TEXT ·CasInt32SeqCst(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    MFENCE
    LOCK
    CMPXCHGL    CX, (BX)
    MFENCE
    SETEQ   ret+16(FP)
    RET

TEXT ·CasUint32SeqCst(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    MFENCE
    LOCK
    CMPXCHGL    CX, (BX)
    MFENCE
    SETEQ   ret+16(FP)
    RET

TEXT ·CasInt64SeqCst(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    MFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    MFENCE
    SETEQ   ret+24(FP)
    RET

TEXT ·CasUint64SeqCst(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    MFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    MFENCE
    SETEQ   ret+24(FP)
    RET

TEXT ·CasUintptrSeqCst(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    MFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    MFENCE
    SETEQ   ret+24(FP)
    RET

// ============================================================
// ADD operations
// ============================================================

// XaddRelaxed - no barrier
TEXT ·AddInt32Relaxed(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    MOVL    AX, CX
    LOCK
    XADDL   AX, (BX)
    ADDL    CX, AX
    MOVL    AX, ret+16(FP)
    RET

TEXT ·AddUint32Relaxed(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    MOVL    AX, CX
    LOCK
    XADDL   AX, (BX)
    ADDL    CX, AX
    MOVL    AX, ret+16(FP)
    RET

TEXT ·AddInt64Relaxed(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·AddUint64Relaxed(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·AddUintptrRelaxed(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MOVQ    AX, ret+16(FP)
    RET

// XaddAcquire - XADD + LFENCE (acquire)
TEXT ·AddInt32Acquire(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    MOVL    AX, CX
    LOCK
    XADDL   AX, (BX)
    ADDL    CX, AX
    LFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·AddUint32Acquire(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    MOVL    AX, CX
    LOCK
    XADDL   AX, (BX)
    ADDL    CX, AX
    LFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·AddInt64Acquire(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    LFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·AddUint64Acquire(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    LFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·AddUintptrAcquire(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    LFENCE
    MOVQ    AX, ret+16(FP)
    RET

// XaddRelease - SFENCE + XADD (release)
TEXT ·AddInt32Release(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    MOVL    AX, CX
    SFENCE
    LOCK
    XADDL   AX, (BX)
    ADDL    CX, AX
    MOVL    AX, ret+16(FP)
    RET

TEXT ·AddUint32Release(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    MOVL    AX, CX
    SFENCE
    LOCK
    XADDL   AX, (BX)
    ADDL    CX, AX
    MOVL    AX, ret+16(FP)
    RET

TEXT ·AddInt64Release(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    SFENCE
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·AddUint64Release(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    SFENCE
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·AddUintptrRelease(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    SFENCE
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MOVQ    AX, ret+16(FP)
    RET

// XaddAcqRel - full barrier
TEXT ·AddInt32AcqRel(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    MOVL    AX, CX
    MFENCE
    LOCK
    XADDL   AX, (BX)
    ADDL    CX, AX
    MFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·AddUint32AcqRel(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    MOVL    AX, CX
    MFENCE
    LOCK
    XADDL   AX, (BX)
    ADDL    CX, AX
    MFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·AddInt64AcqRel(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    MFENCE
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·AddUint64AcqRel(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    MFENCE
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·AddUintptrAcqRel(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    MFENCE
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

// XaddSeqCst - full barrier + global order
TEXT ·AddInt32SeqCst(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    MOVL    AX, CX
    MFENCE
    LOCK
    XADDL   AX, (BX)
    ADDL    CX, AX
    MFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·AddUint32SeqCst(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    MOVL    AX, CX
    MFENCE
    LOCK
    XADDL   AX, (BX)
    ADDL    CX, AX
    MFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·AddInt64SeqCst(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    MFENCE
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·AddUint64SeqCst(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    MFENCE
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·AddUintptrSeqCst(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    MFENCE
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

// ============================================================
// SWAP (Xchg) operations
// ============================================================

// XchgRelaxed - no barrier
TEXT ·SwapInt32Relaxed(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    XCHGL   AX, (BX)
    MOVL    AX, ret+16(FP)
    RET

TEXT ·SwapUint32Relaxed(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    XCHGL   AX, (BX)
    MOVL    AX, ret+16(FP)
    RET

TEXT ·SwapInt64Relaxed(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    XCHGQ   AX, (BX)
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·SwapUint64Relaxed(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    XCHGQ   AX, (BX)
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·SwapUintptrRelaxed(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    XCHGQ   AX, (BX)
    MOVQ    AX, ret+16(FP)
    RET

// XchgAcquire - XCHG + LFENCE
TEXT ·SwapInt32Acquire(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    XCHGL   AX, (BX)
    LFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·SwapUint32Acquire(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    XCHGL   AX, (BX)
    LFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·SwapInt64Acquire(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    XCHGQ   AX, (BX)
    LFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·SwapUint64Acquire(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    XCHGQ   AX, (BX)
    LFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·SwapUintptrAcquire(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    XCHGQ   AX, (BX)
    LFENCE
    MOVQ    AX, ret+16(FP)
    RET

// XchgRelease - SFENCE + XCHG
TEXT ·SwapInt32Release(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    SFENCE
    XCHGL   AX, (BX)
    MOVL    AX, ret+16(FP)
    RET

TEXT ·SwapUint32Release(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    SFENCE
    XCHGL   AX, (BX)
    MOVL    AX, ret+16(FP)
    RET

TEXT ·SwapInt64Release(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    SFENCE
    XCHGQ   AX, (BX)
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·SwapUint64Release(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    SFENCE
    XCHGQ   AX, (BX)
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·SwapUintptrRelease(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    SFENCE
    XCHGQ   AX, (BX)
    MOVQ    AX, ret+16(FP)
    RET

// XchgAcqRel - full barrier
TEXT ·SwapInt32AcqRel(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    MFENCE
    XCHGL   AX, (BX)
    MFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·SwapUint32AcqRel(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    MFENCE
    XCHGL   AX, (BX)
    MFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·SwapInt64AcqRel(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    MFENCE
    XCHGQ   AX, (BX)
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·SwapUint64AcqRel(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    MFENCE
    XCHGQ   AX, (BX)
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·SwapUintptrAcqRel(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    MFENCE
    XCHGQ   AX, (BX)
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

// XchgSeqCst - fill barrier + global order
TEXT ·SwapInt32SeqCst(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    MFENCE
    XCHGL   AX, (BX)
    MFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·SwapUint32SeqCst(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    MFENCE
    XCHGL   AX, (BX)
    MFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·SwapInt64SeqCst(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    MFENCE
    XCHGQ   AX, (BX)
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·SwapUint64SeqCst(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    MFENCE
    XCHGQ   AX, (BX)
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·SwapUintptrSeqCst(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    MFENCE
    XCHGQ   AX, (BX)
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

// ============================================================
// BITWISE OR operations
// ============================================================

// OrRelaxed - no barrier
TEXT ·OrInt32Relaxed(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    LOCK
    ORL     BX, (AX)
    RET

TEXT ·OrUint32Relaxed(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    LOCK
    ORL     BX, (AX)
    RET

TEXT ·OrInt64Relaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ORQ     BX, (AX)
    RET

TEXT ·OrUint64Relaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ORQ     BX, (AX)
    RET

TEXT ·OrUintptrRelaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ORQ     BX, (AX)
    RET

// OrAcquire - OR + LFENCE
TEXT ·OrInt32Acquire(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    LOCK
    ORL     BX, (AX)
    LFENCE
    RET

TEXT ·OrUint32Acquire(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    LOCK
    ORL     BX, (AX)
    LFENCE
    RET

TEXT ·OrInt64Acquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ORQ     BX, (AX)
    LFENCE
    RET

TEXT ·OrUint64Acquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ORQ     BX, (AX)
    LFENCE
    RET

TEXT ·OrUintptrAcquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ORQ     BX, (AX)
    LFENCE
    RET

// OrRelease - SFENCE + OR
TEXT ·OrInt32Release(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    SFENCE
    LOCK
    ORL     BX, (AX)
    RET

TEXT ·OrUint32Release(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    SFENCE
    LOCK
    ORL     BX, (AX)
    RET

TEXT ·OrInt64Release(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    SFENCE
    LOCK
    ORQ     BX, (AX)
    RET

TEXT ·OrUint64Release(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    SFENCE
    LOCK
    ORQ     BX, (AX)
    RET

TEXT ·OrUintptrRelease(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    SFENCE
    LOCK
    ORQ     BX, (AX)
    RET

// OrSeqCst - full barrier
TEXT ·OrInt32SeqCst(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    MFENCE
    LOCK
    ORL     BX, (AX)
    MFENCE
    RET

TEXT ·OrUint32SeqCst(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    MFENCE
    LOCK
    ORL     BX, (AX)
    MFENCE
    RET

TEXT ·OrInt64SeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    MFENCE
    LOCK
    ORQ     BX, (AX)
    MFENCE
    RET

TEXT ·OrUint64SeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    MFENCE
    LOCK
    ORQ     BX, (AX)
    MFENCE
    RET

TEXT ·OrUintptrSeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    MFENCE
    LOCK
    ORQ     BX, (AX)
    MFENCE
    RET

// Or32AcqRel full barrier with old value return
TEXT ·Or32AcqRel(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    val+8(FP), CX
    MFENCE
casloop_or32_acqrel:
    MOVL    CX, DX
    MOVL    (BX), AX
    ORL     AX, DX
    LOCK
    CMPXCHGL    DX, (BX)
    JNZ     casloop_or32_acqrel
    MFENCE
    MOVL    AX, ret+16(FP)
    RET

// Or64AcqRel full barrier with old value return
TEXT ·Or64AcqRel(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), CX
    MFENCE
casloop_or64_acqrel:
    MOVQ    CX, DX
    MOVQ    (BX), AX
    ORQ     AX, DX
    LOCK
    CMPXCHGQ    DX, (BX)
    JNZ     casloop_or64_acqrel
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

// ============================================================
// BITWISE AND operations
// ============================================================

// AndRelaxed - no barrier
TEXT ·AndInt32Relaxed(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    LOCK
    ANDL    BX, (AX)
    RET

TEXT ·AndUint32Relaxed(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    LOCK
    ANDL    BX, (AX)
    RET

TEXT ·AndInt64Relaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ANDQ    BX, (AX)
    RET

TEXT ·AndUint64Relaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ANDQ    BX, (AX)
    RET

TEXT ·AndUintptrRelaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ANDQ    BX, (AX)
    RET

// AndAcquire - AND + LFENCE
TEXT ·AndInt32Acquire(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    LOCK
    ANDL    BX, (AX)
    LFENCE
    RET

TEXT ·AndUint32Acquire(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    LOCK
    ANDL    BX, (AX)
    LFENCE
    RET

TEXT ·AndInt64Acquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ANDQ    BX, (AX)
    LFENCE
    RET

TEXT ·AndUint64Acquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ANDQ    BX, (AX)
    LFENCE
    RET

TEXT ·AndUintptrAcquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ANDQ    BX, (AX)
    LFENCE
    RET

// AndRelease - SFENCE + AND
TEXT ·AndInt32Release(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    SFENCE
    LOCK
    ANDL    BX, (AX)
    RET

TEXT ·AndUint32Release(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    SFENCE
    LOCK
    ANDL    BX, (AX)
    RET

TEXT ·AndInt64Release(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    SFENCE
    LOCK
    ANDQ    BX, (AX)
    RET

TEXT ·AndUint64Release(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    SFENCE
    LOCK
    ANDQ    BX, (AX)
    RET

TEXT ·AndUintptrRelease(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    SFENCE
    LOCK
    ANDQ    BX, (AX)
    RET

// AndSeqCst - full barrier
TEXT ·AndInt32SeqCst(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    MFENCE
    LOCK
    ANDL    BX, (AX)
    MFENCE
    RET

TEXT ·AndUint32SeqCst(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    MFENCE
    LOCK
    ANDL    BX, (AX)
    MFENCE
    RET

TEXT ·AndInt64SeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    MFENCE
    LOCK
    ANDQ    BX, (AX)
    MFENCE
    RET

TEXT ·AndUint64SeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    MFENCE
    LOCK
    ANDQ    BX, (AX)
    MFENCE
    RET

TEXT ·AndUintptrSeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    MFENCE
    LOCK
    ANDQ    BX, (AX)
    MFENCE
    RET

// And32AcqRel full brarrier with old value return
TEXT ·And32AcqRel(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    val+8(FP), CX
    MFENCE
casloop_and32_acqrel:
    MOVL    CX, DX
    MOVL    (BX), AX
    ANDL    AX, DX
    LOCK
    CMPXCHGL    DX, (BX)
    JNZ     casloop_and32_acqrel
    MFENCE
    MOVL    AX, ret+16(FP)
    RET

// And64AcqRel full brarrier with old value return
TEXT ·And64AcqRel(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), CX
    MFENCE
casloop_and64_acqrel:
    MOVQ    CX, DX
    MOVQ    (BX), AX
    ANDQ    AX, DX
    LOCK
    CMPXCHGQ    DX, (BX)
    JNZ     casloop_and64_acqrel
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET
