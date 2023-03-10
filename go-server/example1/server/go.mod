module main

go 1.19

require (
	bitknife.se/game v0.0.0-00010101000000-000000000000
	bitknife.se/core v0.0.0-00010101000000-000000000000
	bitknife.se/socketserver v0.0.0-00010101000000-000000000000
	github.com/c-bata/go-prompt v0.2.6
)

require (
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mattn/go-tty v0.0.3 // indirect
	github.com/pkg/term v1.2.0-beta.2 // indirect
	golang.org/x/sys v0.0.0-20200918174421-af09f7315aff // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace bitknife.se/socketserver => ./socketserver
replace bitknife.se/core => ./core
replace bitknife.se/game => ./game

