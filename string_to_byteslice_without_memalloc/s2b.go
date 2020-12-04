package string_to_byteslice_without_memalloc

import (
	"reflect"
	"unsafe"
)

//May be useless when the struct of string/slice changed in the future version of Go

func s2b(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Cap = sh.Len
	bh.Len = sh.Len
	bh.Data = sh.Data
	return b
}
