module github.com/dtrenin7/fuzz/minify/number

go 1.13

replace github.com/dtrenin7/minify/v2 => ../../../minify

require (
	github.com/dvyukov/go-fuzz v0.0.0-20191022152526-8cb203812681 // indirect
	github.com/dtrenin7/minify/v2 v2.0.0-00010101000000-000000000000
)
