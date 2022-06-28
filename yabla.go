package main

import (
	"os"
	"github.com/rnicolas/yabla/file"
)

func main() {
	fileName := os.Args[1]
	file.Open(fileName)
}