package makro

import (
	"errors"
	"fmt"
	"github.com/orzel1244/pinging_macro/makro/key_input"
	log "github.com/sirupsen/logrus"
	"time"
)

type InternalPingerService struct {
	typingSrv key_input.KeyInputService
	log *log.Entry
}
type PingerService interface {
	Ping(names ...string) error
}

func NewPingerService(typingSrv key_input.KeyInputService, logger *log.Entry) PingerService {
    return &InternalPingerService {
    	typingSrv: typingSrv,
        log: logger,
    }
}

func (s *InternalPingerService) ping(name string, submit bool) error {
	err := s.typingSrv.Type(" ")
	if err != nil {
		return err
	}

	if len(name) == 0 {
		return errors.New("empty name")
	}

	err = s.typingSrv.Type(fmt.Sprintf("@%s", name))
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 650)

	if !submit {
		return nil
	}

	err = s.typingSrv.Submit()
	if err != nil {
		return err
	}
	return nil
}

func (s *InternalPingerService) Ping(names ...string) error {
	fmt.Printf("Received order for: %s\n", names)
	for i, name := range names {
		submit := true

		if i == len(names)-1{
			// IF it is last ping, we don't want to press enter,
			// because for instance in messenger any many other communicators, it would send that whole message
			// it is better to have control over decision if you want to send message, or write something else.
			submit = false
		}

		err := s.ping(name, submit)
		if err != nil {
			return err
		}
	}
	return nil
}