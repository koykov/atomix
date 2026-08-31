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

// loadRelaxed - no barriers, just load
TEXT ·loadInt32Relaxed(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    (BX), AX
    MOVL    AX, ret+8(FP)
    RET

TEXT ·loadUint32Relaxed(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    (BX), AX
    MOVL    AX, ret+8(FP)
    RET

TEXT ·loadInt64Relaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    (BX), AX
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·loadUint64Relaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    (BX), AX
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·loadUintptrRelaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    (BX), AX
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·loadPointerRelaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    (BX), AX
    MOVQ    AX, ret+8(FP)
    RET

// loadAcquire - LFENCE guarantees next reads will not reorder till load
TEXT ·loadInt32Acquire(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    (BX), AX
    LFENCE
    MOVL    AX, ret+8(FP)
    RET

TEXT ·loadUint32Acquire(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    (BX), AX
    LFENCE
    MOVL    AX, ret+8(FP)
    RET

TEXT ·loadInt64Acquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    (BX), AX
    LFENCE
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·loadUint64Acquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    (BX), AX
    LFENCE
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·loadUintptrAcquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    (BX), AX
    LFENCE
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·loadPointerAcquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    (BX), AX
    LFENCE
    MOVQ    AX, ret+8(FP)
    RET

// loadSeqCst - full barrier
TEXT ·loadInt32SeqCst(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MFENCE
    MOVL    (BX), AX
    MFENCE
    MOVL    AX, ret+8(FP)
    RET

TEXT ·loadUint32SeqCst(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MFENCE
    MOVL    (BX), AX
    MFENCE
    MOVL    AX, ret+8(FP)
    RET

TEXT ·loadInt64SeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MFENCE
    MOVQ    (BX), AX
    MFENCE
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·loadUint64SeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MFENCE
    MOVQ    (BX), AX
    MFENCE
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·loadUintptrSeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MFENCE
    MOVQ    (BX), AX
    MFENCE
    MOVQ    AX, ret+8(FP)
    RET

TEXT ·loadPointerSeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MFENCE
    MOVQ    (BX), AX
    MFENCE
    MOVQ    AX, ret+8(FP)
    RET

// ============================================================
// STORE operations
// ============================================================

// storeRelaxed - no barriers
TEXT ·storeInt32Relaxed(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    val+8(FP), AX
    MOVL    AX, (BX)
    RET

TEXT ·storeUint32Relaxed(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    val+8(FP), AX
    MOVL    AX, (BX)
    RET

TEXT ·storeInt64Relaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    MOVQ    AX, (BX)
    RET

TEXT ·storeUint64Relaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    MOVQ    AX, (BX)
    RET

TEXT ·storeUintptrRelaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    MOVQ    AX, (BX)
    RET

// storeRelease - SFENCE guarantees previous writes visible till call
TEXT ·storeInt32Release(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    val+8(FP), AX
    SFENCE
    MOVL    AX, (BX)
    RET

TEXT ·storeUint32Release(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    val+8(FP), AX
    SFENCE
    MOVL    AX, (BX)
    RET

TEXT ·storeInt64Release(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    SFENCE
    MOVQ    AX, (BX)
    RET

TEXT ·storeUint64Release(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    SFENCE
    MOVQ    AX, (BX)
    RET

TEXT ·storeUintptrRelease(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    SFENCE
    MOVQ    AX, (BX)
    RET

// storeSeqCst - full barrier
TEXT ·storeInt32SeqCst(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    val+8(FP), AX
    MFENCE
    MOVL    AX, (BX)
    MFENCE
    RET

TEXT ·storeUint32SeqCst(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), BX
    MOVL    val+8(FP), AX
    MFENCE
    MOVL    AX, (BX)
    MFENCE
    RET

TEXT ·storeInt64SeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    MFENCE
    MOVQ    AX, (BX)
    MFENCE
    RET

TEXT ·storeUint64SeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    MFENCE
    MOVQ    AX, (BX)
    MFENCE
    RET

TEXT ·storeUintptrSeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), BX
    MOVQ    val+8(FP), AX
    MFENCE
    MOVQ    AX, (BX)
    MFENCE
    RET

// ============================================================
// CAS (Compare And Swap) operations
// ============================================================

// casRelaxed - no barriers
TEXT ·casInt32Relaxed(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    LOCK
    CMPXCHGL    CX, (BX)
    SETEQ   ret+16(FP)
    RET

TEXT ·casUint32Relaxed(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    LOCK
    CMPXCHGL    CX, (BX)
    SETEQ   ret+16(FP)
    RET

TEXT ·casInt64Relaxed(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    LOCK
    CMPXCHGQ    CX, (BX)
    SETEQ   ret+24(FP)
    RET

TEXT ·casUint64Relaxed(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    LOCK
    CMPXCHGQ    CX, (BX)
    SETEQ   ret+24(FP)
    RET

TEXT ·casUintptrRelaxed(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    LOCK
    CMPXCHGQ    CX, (BX)
    SETEQ   ret+24(FP)
    RET

// casAcquire - CAS with acquire semantic(barrier after sucessful operation)
TEXT ·casInt32Acquire(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    LOCK
    CMPXCHGL    CX, (BX)
    LFENCE
    SETEQ   ret+16(FP)
    RET

TEXT ·casUint32Acquire(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    LOCK
    CMPXCHGL    CX, (BX)
    LFENCE
    SETEQ   ret+16(FP)
    RET

TEXT ·casInt64Acquire(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    LOCK
    CMPXCHGQ    CX, (BX)
    LFENCE
    SETEQ   ret+24(FP)
    RET

TEXT ·casUint64Acquire(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    LOCK
    CMPXCHGQ    CX, (BX)
    LFENCE
    SETEQ   ret+24(FP)
    RET

TEXT ·casUintptrAcquire(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    LOCK
    CMPXCHGQ    CX, (BX)
    LFENCE
    SETEQ   ret+24(FP)
    RET

// casRelease - CAS with release seantic (barrier before operation)
TEXT ·casInt32Release(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    SFENCE
    LOCK
    CMPXCHGL    CX, (BX)
    SETEQ   ret+16(FP)
    RET

TEXT ·casUint32Release(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    SFENCE
    LOCK
    CMPXCHGL    CX, (BX)
    SETEQ   ret+16(FP)
    RET

TEXT ·casInt64Release(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    SFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    SETEQ   ret+24(FP)
    RET

TEXT ·casUint64Release(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    SFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    SETEQ   ret+24(FP)
    RET

TEXT ·casUintptrRelease(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    SFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    SETEQ   ret+24(FP)
    RET

// casAcqRel - full barrier (both direction)
TEXT ·casInt32AcqRel(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    MFENCE
    LOCK
    CMPXCHGL    CX, (BX)
    MFENCE
    SETEQ   ret+16(FP)
    RET

TEXT ·casUint32AcqRel(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    MFENCE
    LOCK
    CMPXCHGL    CX, (BX)
    MFENCE
    SETEQ   ret+16(FP)
    RET

TEXT ·casInt64AcqRel(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    MFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    MFENCE
    SETEQ   ret+24(FP)
    RET

TEXT ·casUint64AcqRel(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    MFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    MFENCE
    SETEQ   ret+24(FP)
    RET

TEXT ·casUintptrAcqRel(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    MFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    MFENCE
    SETEQ   ret+24(FP)
    RET

// casSeqCst - full barrier + global order
TEXT ·casInt32SeqCst(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    MFENCE
    LOCK
    CMPXCHGL    CX, (BX)
    MFENCE
    SETEQ   ret+16(FP)
    RET

TEXT ·casUint32SeqCst(SB), NOSPLIT, $0-17
    MOVQ    ptr+0(FP), BX
    MOVL    old+8(FP), AX
    MOVL    new+12(FP), CX
    MFENCE
    LOCK
    CMPXCHGL    CX, (BX)
    MFENCE
    SETEQ   ret+16(FP)
    RET

TEXT ·casInt64SeqCst(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    MFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    MFENCE
    SETEQ   ret+24(FP)
    RET

TEXT ·casUint64SeqCst(SB), NOSPLIT, $0-25
    MOVQ    ptr+0(FP), BX
    MOVQ    old+8(FP), AX
    MOVQ    new+16(FP), CX
    MFENCE
    LOCK
    CMPXCHGQ    CX, (BX)
    MFENCE
    SETEQ   ret+24(FP)
    RET

TEXT ·casUintptrSeqCst(SB), NOSPLIT, $0-25
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
TEXT ·addInt32Relaxed(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    MOVL    AX, CX
    LOCK
    XADDL   AX, (BX)
    ADDL    CX, AX
    MOVL    AX, ret+16(FP)
    RET

TEXT ·addUint32Relaxed(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    MOVL    AX, CX
    LOCK
    XADDL   AX, (BX)
    ADDL    CX, AX
    MOVL    AX, ret+16(FP)
    RET

TEXT ·addInt64Relaxed(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·addUint64Relaxed(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·addUintptrRelaxed(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MOVQ    AX, ret+16(FP)
    RET

// XaddAcquire - XADD + LFENCE (acquire)
TEXT ·addInt32Acquire(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    MOVL    AX, CX
    LOCK
    XADDL   AX, (BX)
    ADDL    CX, AX
    LFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·addUint32Acquire(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    MOVL    AX, CX
    LOCK
    XADDL   AX, (BX)
    ADDL    CX, AX
    LFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·addInt64Acquire(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    LFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·addUint64Acquire(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    LFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·addUintptrAcquire(SB), NOSPLIT, $0-24
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
TEXT ·addInt32Release(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    MOVL    AX, CX
    SFENCE
    LOCK
    XADDL   AX, (BX)
    ADDL    CX, AX
    MOVL    AX, ret+16(FP)
    RET

TEXT ·addUint32Release(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    MOVL    AX, CX
    SFENCE
    LOCK
    XADDL   AX, (BX)
    ADDL    CX, AX
    MOVL    AX, ret+16(FP)
    RET

TEXT ·addInt64Release(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    SFENCE
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·addUint64Release(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    MOVQ    AX, CX
    SFENCE
    LOCK
    XADDQ   AX, (BX)
    ADDQ    CX, AX
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·addUintptrRelease(SB), NOSPLIT, $0-24
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
TEXT ·addInt32AcqRel(SB), NOSPLIT, $0-20
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

TEXT ·addUint32AcqRel(SB), NOSPLIT, $0-20
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

TEXT ·addInt64AcqRel(SB), NOSPLIT, $0-24
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

TEXT ·addUint64AcqRel(SB), NOSPLIT, $0-24
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

TEXT ·addUintptrAcqRel(SB), NOSPLIT, $0-24
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
TEXT ·addInt32SeqCst(SB), NOSPLIT, $0-20
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

TEXT ·addUint32SeqCst(SB), NOSPLIT, $0-20
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

TEXT ·addInt64SeqCst(SB), NOSPLIT, $0-24
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

TEXT ·addUint64SeqCst(SB), NOSPLIT, $0-24
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

TEXT ·addUintptrSeqCst(SB), NOSPLIT, $0-24
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
TEXT ·swapInt32Relaxed(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    XCHGL   AX, (BX)
    MOVL    AX, ret+16(FP)
    RET

TEXT ·swapUint32Relaxed(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    XCHGL   AX, (BX)
    MOVL    AX, ret+16(FP)
    RET

TEXT ·swapInt64Relaxed(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    XCHGQ   AX, (BX)
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·swapUint64Relaxed(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    XCHGQ   AX, (BX)
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·swapUintptrRelaxed(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    XCHGQ   AX, (BX)
    MOVQ    AX, ret+16(FP)
    RET

// XchgAcquire - XCHG + LFENCE
TEXT ·swapInt32Acquire(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    XCHGL   AX, (BX)
    LFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·swapUint32Acquire(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    XCHGL   AX, (BX)
    LFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·swapInt64Acquire(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    XCHGQ   AX, (BX)
    LFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·swapUint64Acquire(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    XCHGQ   AX, (BX)
    LFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·swapUintptrAcquire(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    XCHGQ   AX, (BX)
    LFENCE
    MOVQ    AX, ret+16(FP)
    RET

// XchgRelease - SFENCE + XCHG
TEXT ·swapInt32Release(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    SFENCE
    XCHGL   AX, (BX)
    MOVL    AX, ret+16(FP)
    RET

TEXT ·swapUint32Release(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    SFENCE
    XCHGL   AX, (BX)
    MOVL    AX, ret+16(FP)
    RET

TEXT ·swapInt64Release(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    SFENCE
    XCHGQ   AX, (BX)
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·swapUint64Release(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    SFENCE
    XCHGQ   AX, (BX)
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·swapUintptrRelease(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    SFENCE
    XCHGQ   AX, (BX)
    MOVQ    AX, ret+16(FP)
    RET

// XchgAcqRel - full barrier
TEXT ·swapInt32AcqRel(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    MFENCE
    XCHGL   AX, (BX)
    MFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·swapUint32AcqRel(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    MFENCE
    XCHGL   AX, (BX)
    MFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·swapInt64AcqRel(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    MFENCE
    XCHGQ   AX, (BX)
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·swapUint64AcqRel(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    MFENCE
    XCHGQ   AX, (BX)
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·swapUintptrAcqRel(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    MFENCE
    XCHGQ   AX, (BX)
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

// XchgSeqCst - fill barrier + global order
TEXT ·swapInt32SeqCst(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    MFENCE
    XCHGL   AX, (BX)
    MFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·swapUint32SeqCst(SB), NOSPLIT, $0-20
    MOVQ    ptr+0(FP), BX
    MOVL    new+8(FP), AX
    MFENCE
    XCHGL   AX, (BX)
    MFENCE
    MOVL    AX, ret+16(FP)
    RET

TEXT ·swapInt64SeqCst(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    MFENCE
    XCHGQ   AX, (BX)
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·swapUint64SeqCst(SB), NOSPLIT, $0-24
    MOVQ    ptr+0(FP), BX
    MOVQ    new+8(FP), AX
    MFENCE
    XCHGQ   AX, (BX)
    MFENCE
    MOVQ    AX, ret+16(FP)
    RET

TEXT ·swapUintptrSeqCst(SB), NOSPLIT, $0-24
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

// orRelaxed - no barrier
TEXT ·orInt32Relaxed(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    LOCK
    ORL     BX, (AX)
    RET

TEXT ·orUint32Relaxed(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    LOCK
    ORL     BX, (AX)
    RET

TEXT ·orInt64Relaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ORQ     BX, (AX)
    RET

TEXT ·orUint64Relaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ORQ     BX, (AX)
    RET

TEXT ·orUintptrRelaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ORQ     BX, (AX)
    RET

// orAcquire - OR + LFENCE
TEXT ·orInt32Acquire(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    LOCK
    ORL     BX, (AX)
    LFENCE
    RET

TEXT ·orUint32Acquire(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    LOCK
    ORL     BX, (AX)
    LFENCE
    RET

TEXT ·orInt64Acquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ORQ     BX, (AX)
    LFENCE
    RET

TEXT ·orUint64Acquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ORQ     BX, (AX)
    LFENCE
    RET

TEXT ·orUintptrAcquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ORQ     BX, (AX)
    LFENCE
    RET

// orRelease - SFENCE + OR
TEXT ·orInt32Release(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    SFENCE
    LOCK
    ORL     BX, (AX)
    RET

TEXT ·orUint32Release(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    SFENCE
    LOCK
    ORL     BX, (AX)
    RET

TEXT ·orInt64Release(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    SFENCE
    LOCK
    ORQ     BX, (AX)
    RET

TEXT ·orUint64Release(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    SFENCE
    LOCK
    ORQ     BX, (AX)
    RET

TEXT ·orUintptrRelease(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    SFENCE
    LOCK
    ORQ     BX, (AX)
    RET

// orSeqCst - full barrier
TEXT ·orInt32SeqCst(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    MFENCE
    LOCK
    ORL     BX, (AX)
    MFENCE
    RET

TEXT ·orUint32SeqCst(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    MFENCE
    LOCK
    ORL     BX, (AX)
    MFENCE
    RET

TEXT ·orInt64SeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    MFENCE
    LOCK
    ORQ     BX, (AX)
    MFENCE
    RET

TEXT ·orUint64SeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    MFENCE
    LOCK
    ORQ     BX, (AX)
    MFENCE
    RET

TEXT ·orUintptrSeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    MFENCE
    LOCK
    ORQ     BX, (AX)
    MFENCE
    RET

// or32AcqRel full barrier with old value return
TEXT ·or32AcqRel(SB), NOSPLIT, $0-20
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

// or64AcqRel full barrier with old value return
TEXT ·or64AcqRel(SB), NOSPLIT, $0-24
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

// andRelaxed - no barrier
TEXT ·andInt32Relaxed(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    LOCK
    ANDL    BX, (AX)
    RET

TEXT ·andUint32Relaxed(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    LOCK
    ANDL    BX, (AX)
    RET

TEXT ·andInt64Relaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ANDQ    BX, (AX)
    RET

TEXT ·andUint64Relaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ANDQ    BX, (AX)
    RET

TEXT ·andUintptrRelaxed(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ANDQ    BX, (AX)
    RET

// andAcquire - AND + LFENCE
TEXT ·andInt32Acquire(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    LOCK
    ANDL    BX, (AX)
    LFENCE
    RET

TEXT ·andUint32Acquire(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    LOCK
    ANDL    BX, (AX)
    LFENCE
    RET

TEXT ·andInt64Acquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ANDQ    BX, (AX)
    LFENCE
    RET

TEXT ·andUint64Acquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ANDQ    BX, (AX)
    LFENCE
    RET

TEXT ·andUintptrAcquire(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    LOCK
    ANDQ    BX, (AX)
    LFENCE
    RET

// andRelease - SFENCE + AND
TEXT ·andInt32Release(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    SFENCE
    LOCK
    ANDL    BX, (AX)
    RET

TEXT ·andUint32Release(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    SFENCE
    LOCK
    ANDL    BX, (AX)
    RET

TEXT ·andInt64Release(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    SFENCE
    LOCK
    ANDQ    BX, (AX)
    RET

TEXT ·andUint64Release(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    SFENCE
    LOCK
    ANDQ    BX, (AX)
    RET

TEXT ·andUintptrRelease(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    SFENCE
    LOCK
    ANDQ    BX, (AX)
    RET

// andSeqCst - full barrier
TEXT ·andInt32SeqCst(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    MFENCE
    LOCK
    ANDL    BX, (AX)
    MFENCE
    RET

TEXT ·andUint32SeqCst(SB), NOSPLIT, $0-12
    MOVQ    ptr+0(FP), AX
    MOVL    val+8(FP), BX
    MFENCE
    LOCK
    ANDL    BX, (AX)
    MFENCE
    RET

TEXT ·andInt64SeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    MFENCE
    LOCK
    ANDQ    BX, (AX)
    MFENCE
    RET

TEXT ·andUint64SeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    MFENCE
    LOCK
    ANDQ    BX, (AX)
    MFENCE
    RET

TEXT ·andUintptrSeqCst(SB), NOSPLIT, $0-16
    MOVQ    ptr+0(FP), AX
    MOVQ    val+8(FP), BX
    MFENCE
    LOCK
    ANDQ    BX, (AX)
    MFENCE
    RET

// and32AcqRel full brarrier with old value return
TEXT ·and32AcqRel(SB), NOSPLIT, $0-20
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

// and64AcqRel full brarrier with old value return
TEXT ·and64AcqRel(SB), NOSPLIT, $0-24
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
