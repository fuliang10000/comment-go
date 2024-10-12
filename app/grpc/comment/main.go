package main

import (
	"commentService/internal/comment"
	"commentService/pkg/build"
)

func main() {
	build.CheckFlagPrintVersion()
	comment.Run()
}
