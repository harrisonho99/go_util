package tools

import (
	"bytes"
	"encoding/gob"
	"errors"
	"reflect"
)

type P struct {
	A, B, C int
	D       string
	E       []int
}

type Q struct {
	A, C int
	D    string
	E    []int
}

// Deep copy will copy data from src to dst.
// src and dst must be same type.
// dst must be a pointer.
// Works for struct, map and slice
func DeepCopy(src, dst interface{}) (err error) {
	if T := reflect.TypeOf(dst); T.Kind() != reflect.Ptr {
		return errors.New("destination must be a pointer type")
	}

	buff := new(bytes.Buffer)

	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)

	enc.Encode(src)
	dec.Decode(dst)

	return
}
