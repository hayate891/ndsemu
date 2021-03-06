#include "textflag.h"

// func mul128(x, y int64) (z1 int64, z0 unt64)
TEXT ·mul128(SB),NOSPLIT,$0
	MOVQ x+0(FP), AX
	IMULQ y+8(FP)
	MOVQ DX, z1+16(FP)
	MOVQ AX, z0+24(FP)
	RET

// func div128(hinum, lonum, den int64) (quo int64, rem unt64)
TEXT ·div128(SB),NOSPLIT,$0
	MOVQ hinum+0(FP), DX
	MOVQ lonum+8(FP), AX
	IDIVQ den+16(FP)
	MOVQ AX, quo+24(FP)
	MOVQ DX, rem+32(FP)
	RET

