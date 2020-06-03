// +build gofuzz
package fuzz

import (
	"github.com/dtrenin7/minify/v2"
	"github.com/dtrenin7/parse"
)

func Fuzz(data []byte) int {
	prec := 0
	if len(data) > 0 {
		x := data[0]
		data = data[1:]
		prec = int(x) % 32
	}
	data = parse.Copy(data) // ignore const-input error for OSS-Fuzz
	data = minify.Number(data, prec)
	return 1
}
