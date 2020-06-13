package lib

import (
	"fmt"

	"github.com/mbndr/figlet4go"
	"github.com/nicosrgh/straw-hat/config"
)

// Greet :
func Greet() {
	ascii := figlet4go.NewAsciiRender()

	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{figlet4go.ColorRed}

	renderStr, _ := ascii.RenderOpts(config.C.AppName, options)
	fmt.Print(renderStr)
	fmt.Print(config.C.AppQuote)
	fmt.Printf("\nv%s\n", config.C.Version)
}
