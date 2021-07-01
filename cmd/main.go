package main

import (
	"github.com/orzel1244/pinging_macro/makro"
	"github.com/orzel1244/pinging_macro/makro/key_input"
	log "github.com/sirupsen/logrus"
	"gobot.io/x/gobot/platforms/keyboard"
)

func main() {
	mainLogger := log.New()
	mainLogger.SetLevel(log.DebugLevel)
	logger := log.NewEntry(mainLogger)

	typingSrv := key_input.NewKeyInputService(logger)
	pingingSrv := makro.NewPingerService(typingSrv, logger)

	keyboardDriver := keyboard.NewDriver()
	pharaseSrv := makro.NewPharaseListenerService(keyboardDriver, typingSrv.GetKeysList(), logger)

	if err := pharaseSrv.Listen(pingingSrv.Ping); err != nil {
		log.Fatal(err)
	}
}
