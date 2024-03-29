// Copyright (c) 2012 VMware, Inc.

package gosigar

import (
	"unsafe"
)

// TODO (2020-07-14): Fix the unsafe pointer conversion. https://github.com/elastic/gosigar/issues/138
//go:nocheckptr
func bytePtrToString(ptr *int8) string {
	bytes := (*[10000]byte)(unsafe.Pointer(ptr))

	n := 0
	for bytes[n] != 0 {
		n++
	}

	return string(bytes[0:n])
}

func chop(buf []byte) []byte {
	return buf[0 : len(buf)-1]
}
