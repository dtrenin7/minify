// +build gofuzz
package fuzz

import (
	"bytes"
	"io/ioutil"

	"github.com/dtrenin7/minify/v2"
	"github.com/dtrenin7/minify/v2/html"
)

func Fuzz(data []byte) int {
	r := bytes.NewBuffer(data)
	_ = html.Minify(minify.New(), ioutil.Discard, r, nil)
	return 1
}
