package alloc

/*
#include <stdlib.h>
#include <stdio.h>
void _safe_free(void **pointer);
*/
import "C"
import "unsafe"

func Free(p *unsafe.Pointer) {
	C._safe_free(p)
	// C.free(p)
}
