package main

import (
	"github.com/orzel1244/pinging_macro/makro"
	"github.com/orzel1244/pinging_macro/makro/key_input"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)

	mainLogger := log.New()
	mainLogger.SetLevel(log.DebugLevel)
	logger := log.NewEntry(mainLogger)

	typingSrv := key_input.NewKeyInputService(logger)
	pingingSrv := makro.NewPingerService(typingSrv, logger)

	err := pingingSrv.Ping("sasin nie wiedzial o wYbOrAcH ktore sIe NiE oDbYly")
	if err != nil {
		logger.WithError(err).Error("cannot ping")
	}
}