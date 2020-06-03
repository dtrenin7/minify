// +build gofuzz
package fuzz

import (
	"github.com/dtrenin7/minify/v2"
	"github.com/dtrenin7/parse"
)

func Fuzz(data []byte) int {
	data = parse.Copy(data) // ignore const-input error for OSS-Fuzz
	data = minify.Mediatype(data)
	return 1
}
