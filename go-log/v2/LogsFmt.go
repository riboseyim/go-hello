package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/logmatic/logmatic-go"
)

func main() {
	// use JSONFormatter
	log.SetFormatter(&logmatic.JSONFormatter{})
	// log an event as usual with logrus
	log.WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1}).Info("My first ssl event from golang")
}
