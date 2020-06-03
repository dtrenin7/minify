// +build gofuzz
package fuzz

import (
	"github.com/dtrenin7/minify/v2/svg"
	"github.com/dtrenin7/parse"
)

func Fuzz(data []byte) int {
	pathDataBuffer := svg.NewPathData(&svg.Minifier{Decimals: -1})
	data = parse.Copy(data) // ignore const-input error for OSS-Fuzz
	data = pathDataBuffer.ShortenPathData(data)
	return 1
}
