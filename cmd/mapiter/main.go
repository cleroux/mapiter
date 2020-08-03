package main

import (
	"github.com/cleroux/mapiter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(mapiter.Analyzer)
}
