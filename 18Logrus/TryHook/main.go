package main

//Hook
import (
	"github.com/sirupsen/logrus"
	lSyslog "github.com/sirupsen/logrus/hooks/syslog"
	"log/syslog"
)

func main() {
	log       := logrus.New()
	hook, err := lSyslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, "")

	if err == nil {
		log.Hooks.Add(hook)
	}
}
