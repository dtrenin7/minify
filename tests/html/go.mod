module github.com/dtrenin7/fuzz/minify/html

go 1.13

replace github.com/dtrenin7/minify/v2 => ../../../minify

require (
	github.com/dvyukov/go-fuzz v0.0.0-20191022152526-8cb203812681 // indirect
	github.com/dtrenin7/minify/v2 v2.5.2
)
