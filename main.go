package main

import (
	"github.com/alecthomas/kong"
)

type Context struct{}

type _App struct {
	Clean _CleanCmd `cmd help:"Clean a yara file by removing all comments and metadata."`
}

func main() {
	ctx := kong.Parse(&_App{})
	err := ctx.Run(&Context{})
	ctx.FatalIfErrorf(err)
}
