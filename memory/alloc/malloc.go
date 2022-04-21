package alloc

/*
#include <stdlib.h>
#include <stdio.h>
void *wrap_malloc(size_t count, size_t size);
*/
import "C"
import (
	"reflect"
	"unsafe"
)

func Malloc(count int, size int) unsafe.Pointer {
	p := C.wrap_malloc(C.size_t(count), C.size_t(size))
	return p
}

func MakeSlice(count int, size int) []interface{} {
	p := Malloc(count, size)
	data := reflect.ValueOf(p).Pointer()
	var sl_head = reflect.SliceHeader{Data: data, Len: count, Cap: count}
	var sl = *(*[]interface{})(unsafe.Pointer(&sl_head))
	return sl
}
