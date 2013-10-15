package main

import (
	"flag"
	"fmt"
	"gutscript"
	"os"
)

var watch = flag.Bool("w", false, "watch for changes.")

func main() {
	flag.Parse()

	gutscript.GutsDebug = 0

	out, err := gutscript.CompileFile(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(out)
}
