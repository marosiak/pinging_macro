package main

import (
	"errors"
	"fmt"
	"github.com/orzel1244/pinging_macro/makro/key_input"
	log "github.com/sirupsen/logrus"
	"time"
)
func ping(typingSrv key_input.KeyInputService, name string) error {
	err := typingSrv.Type(" ")
	if err != nil {
		return err
	}

	if len(name) == 0 {
		return errors.New("empty name")
	}

	err = typingSrv.Type(fmt.Sprintf("@%s", name))
	if err != nil {
		return err
	}
 	time.Sleep(time.Millisecond * 350)

	err = typingSrv.Submit()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	time.Sleep(1 * time.Second)

	mainLogger := log.New()
	mainLogger.SetLevel(log.DebugLevel)
	logger := log.NewEntry(mainLogger)

	typingSrv := key_input.NewKeyInputService(logger)


	err := ping(typingSrv, "sasin nie wiedzial o wYbOrAcH ktore sIe NiE oDbYly")
	if err != nil {
		println(err)
	}
}