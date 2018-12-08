package main

import (
	"log/syslog"

	log "github.com/Sirupsen/logrus"
	logrus_syslog "github.com/Sirupsen/logrus/hooks/syslog"
)

//Sources for older versions is available at
//https://my.balabit.com/downloads/syslog-ng/sources
//Binary packages from 3rd parties are available from
//https://www.balabit.com/network-security/syslog-ng/opensource-logging-system/downloads/3rd_party

func main() {

	//  hook, err := logrus_syslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, "")
	hook, err := logrus_syslog.NewSyslogHook("udp", "59.37.66.156:514", syslog.LOG_INFO, "")

	if err != nil {
		log.Error("Unable to connect to local syslog daemon")
	} else {
		log.AddHook(hook)
	}

	contextLogger := log.WithFields(log.Fields{
		"common": "XXX common content XXX",
		"other":  "YYY special context YYY",
	})

	contextLogger.Info("AAAAAAAAAAAA")
	contextLogger.Info("BBBBBBBBBBBB")

}
