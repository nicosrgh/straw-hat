package main

import (
	"github.com/nicosrgh/straw-hat/app/server"
	"github.com/nicosrgh/straw-hat/lib"
)

func main() {
	lib.Greet()
	server.Init()
}
