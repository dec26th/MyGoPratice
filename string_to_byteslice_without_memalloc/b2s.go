package string_to_byteslice_without_memalloc

import "unsafe"

//May be useless when the struct of string/slice changed in the future version of Go

func b2s(b []byte) (s string) {
	return *(*string)(unsafe.Pointer(&b))
}
