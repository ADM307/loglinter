package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"github.com/ADM307/loglinter/pkg/loglinter"
)

func main() {
	singlechecker.Main(loglinter.Analyzer)
}
