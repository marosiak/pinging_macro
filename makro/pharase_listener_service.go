package makro

import (
	"github.com/orzel1244/pinging_macro/makro/key_input"
	log "github.com/sirupsen/logrus"
	"github.com/wangch/win"
	"gobot.io/x/gobot/platforms/keyboard"
	"time"
)



type InternalPharaseListenerService struct {
	keyboard *keyboard.Driver
	keysList key_input.KeysListType
	log *log.Entry
}
type PharaseListenerService interface {
	Listen(func(name string) error) error
}

func NewPharaseListenerService(keyboardDriver *keyboard.Driver, keysList key_input.KeysListType, logger *log.Entry) PharaseListenerService {
	return &InternalPharaseListenerService {
		keyboard: keyboardDriver,
		keysList: keysList,
		log: logger,
	}
}


type charStruct struct {
	Char string
	PressedAt time.Time
}

func (s *InternalPharaseListenerService) Listen(pingFunc func(name string) error) error {
	var pharase []charStruct
	// TODO: Worker który sprawdzi, kiedy był kliknięty ostatnio przycisk, i jeżeli ten czas jest dłuższy niż X wtedy usuwamy fraze
	for {
		// TODO: Jakiś sposób na cross platform, można np zrobić w interfejsie key_input'a funkcje zwracającą liste przycisków dla danej platformy
		for _, key := range s.keysList {
			if win.GetKeyState(int32(key.VirtualCode))>>15 != 0 {
				var lastChar charStruct

				pharaseLen := len(pharase)
				if pharaseLen != 0 {
					lastChar = pharase[pharaseLen-1]
					if lastChar.Char == key.DisplayName {
						if time.Now().Before(lastChar.PressedAt.Add(time.Millisecond * 800)) {
							break // To early to read the same key
						}
					}
				}

				if time.Now().Before(lastChar.PressedAt.Add(time.Millisecond * 50)) {
					// In case you'd accidently click 2 keys at the same time, it would multiply it few times
					break
				}

				pharase = append(pharase, charStruct{
					Char:      key.DisplayName,
					PressedAt: time.Now(),
				})

				// For dev/debug purpose only
				for _, v := range pharase {
					print(v.Char)
				}
				println("\n")
			}
		}
		//err := pingFunc("@stary")
		//if err != nil {
		//	return err
		//}
	}
}
