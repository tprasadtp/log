package main

import (
	"github.com/tprasadtp/log"
)

func main() {

	l := log.NewDefault()
	l.Level = log.TraceLevel
	log.Current = l

	log.Trace("Hello,", "World")
	log.Debug("Hello,", "World")
	log.Info("Hello,", "World")
	log.Warn("Hello,", "World")
	log.Fatal("Hello,", "World")

}
