package benchmarks

import (
	"io/ioutil"

	"github.com/dtrenin7/minify/v2"
	"github.com/dtrenin7/parse/v2/buffer"
)

var m = minify.New()
var r = map[string]*buffer.Reader{}
var w = map[string]*buffer.Writer{}

func load(filename string) {
	sample, _ := ioutil.ReadFile(filename)
	r[filename] = buffer.NewReader(sample)
	w[filename] = buffer.NewWriter(make([]byte, 0, len(sample)))
}
