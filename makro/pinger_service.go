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
	Ping(name string) error
}

func NewPingerService(typingSrv key_input.KeyInputService, logger *log.Entry) PingerService {
    return &InternalPingerService {
    	typingSrv: typingSrv,
        log: logger,
    }
}

func (s *InternalPingerService) Ping(name string) error {
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
	time.Sleep(time.Millisecond * 350)

	err = s.typingSrv.Submit()
	if err != nil {
		return err
	}
	return nil
}