// +build gofuzz
package fuzz

import (
	"github.com/dtrenin7/parse/v2/css"
)

func Fuzz(data []byte) int {
	_ = css.IsIdent(data)
	return 1
}
