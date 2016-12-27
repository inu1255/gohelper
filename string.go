package gohelper

import (
	"io"
	"io/ioutil"
	"regexp"
	"strings"
	"unsafe"
)

var (
	re_hump *regexp.Regexp
)

func humpStrFunc(from string) string {
	if len(from) == 2 {
		return strings.ToUpper(from[1:])
	}
	return strings.ToUpper(from)
}

func HumpStr(s string) string {
	if re_hump == nil {
		re_hump = regexp.MustCompile(`([-_]|^)[a-z]`)
	}
	return re_hump.ReplaceAllStringFunc(s, humpStrFunc)
}

func UpperFirstLetter(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(s[:1]) + s[1:]
	}
	return s
}

func ReadAll(r io.Reader) (string, error) {
	s, err := ioutil.ReadAll(r)
	return Bytes2str(s), err
}

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
