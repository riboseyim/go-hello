package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/logmatic/logmatic-go"
)

func main() {
	// instantiate a new Logger with your Logmatic APIKey
	// 国内访问比较慢
	//log.AddHook(logmatic.NewLogmaticHook("p53uTkOhSEqI3-116DynkQ"))

	// use JSONFormatter
	log.SetFormatter(&logmatic.JSONFormatter{})
	// log an event as usual with logrus
	//log.WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1 }).Info("My first ssl event from golang")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "XXX common content XXX",
		"other":  "YYY special context YYY",
	})

	contextLogger.Info("AAAAAAAAAAAA")
	contextLogger.Info("BBBBBBBBBBBB")
}
