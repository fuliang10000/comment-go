package build

import (
	"flag"
	"fmt"
	"os"
)

var Version string
var BuildAt string

func PrintVersion() {
	fmt.Println("Version:", Version, "BuildAt:", BuildAt)
}

var (
	v = flag.Bool("v", false, "show version")
)

func CheckFlagPrintVersion() {
	flag.Parse()
	if *v {
		PrintVersion()
		os.Exit(0)
	}
}
