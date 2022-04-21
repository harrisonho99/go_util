package alloc

import (
	"testing"
)

func TestFree(t *testing.T) {
	ptr := Malloc(1, 1)
	var byte_ptr = (*byte)(ptr)
	*byte_ptr = 20
	Free(&ptr)
}
