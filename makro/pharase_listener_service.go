package makro

import (
	"github.com/orzel1244/pinging_macro/makro/key_input"
	log "github.com/sirupsen/logrus"
	"github.com/wangch/win"
	"gobot.io/x/gobot/platforms/keyboard"
	"strings"
	"sync"
	"time"
)

type InternalPharaseListenerService struct {
	keyboard  *keyboard.Driver
	keysList  key_input.KeysListType
	pingerSrv PingerService
	mutex     *sync.Mutex
	log       *log.Entry
}
type PharaseListenerService interface {
	Listen() error
}

func NewPharaseListenerService(keyboardDriver *keyboard.Driver, keysList key_input.KeysListType, pingerSrv PingerService, logger *log.Entry) PharaseListenerService {
	return &InternalPharaseListenerService{
		keyboard:  keyboardDriver,
		keysList:  keysList,
		pingerSrv: pingerSrv,
		mutex:     &sync.Mutex{},
		log:       logger,
	}
}

type charStruct struct {
	Char      string
	PressedAt time.Time
}

type pharaseType []charStruct

func (p pharaseType) String() string {
	var output string

	for _, v := range p {
		output = output + v.Char
	}
	return output
}

func (p pharaseType) Contains(text string, ignoreCase bool) bool {
	phrase := p.String()

	if ignoreCase {
		phrase = strings.ToUpper(phrase)
		text = strings.ToUpper(text)
	}
	return strings.Contains(p.String(), text)
}

func (s *InternalPharaseListenerService) autoClear(phrase *pharaseType) {
	clear := func(phrase *pharaseType) {
		s.mutex.Lock()
		*phrase = []charStruct{}
		s.mutex.Unlock()
	}

	for {
		phraseLen := len(*phrase)

		if phraseLen == 0 {
			continue
		}

		// It will clear too long phrase
		if phraseLen > 24 {
			clear(&*phrase)
			continue
		}

		//p := *phrase
		// If the last char was typed 8 secs ago - it will clear that phrase
		//if time.Now().After(p[phraseLen-1].PressedAt.Add(time.Second * 8)) {
		//	clear(&*phrase)
		//}
	}
}

func (s *InternalPharaseListenerService) Listen() error {
	var pharase pharaseType
	go s.autoClear(&pharase)

	for {
		for _, key := range s.keysList {
			if win.GetKeyState(int32(key.VirtualCode))>>15 != 0 {
				var lastChar charStruct

				phraseLen := len(pharase)
				if phraseLen != 0 {

					lastChar = pharase[phraseLen-1]

					if lastChar.Char == "@" && key.DisplayName == "2" {
						break // It is temporary PoC quick solution for bug, it is going to be changed
					}

					if lastChar.Char == key.DisplayName {
						if time.Now().Before(lastChar.PressedAt.Add(time.Millisecond * 600)) {
							break // To early to read the same key
						}
					}
				}

				if time.Now().Before(lastChar.PressedAt.Add(time.Millisecond * 50)) {
					// In case you'd accidentally click 2 keys at the same time, it would multiply it few times
					break
				}

				pharase = append(pharase, charStruct{
					Char:      key.DisplayName,
					PressedAt: time.Now(),
				})
			}
		}

		if pharase.Contains("@rap", true) {
			theGuys := []string{"marek", "fornal", "jan pawel drugi", "lenart", "synow"}

			err := s.pingerSrv.Ping(theGuys...)
			if err != nil {
				return nil
			}

			pharase = pharaseType{} // This way we'll avoid inf loop
		}
	}
}

@rap @marek
 @fornal
 @lenart
 @synow
 @jan pawel drugi
