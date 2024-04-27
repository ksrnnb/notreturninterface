package main

import (
	"github.com/ksrnnb/notreturninterface"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(notreturninterface.Analyzer) }
