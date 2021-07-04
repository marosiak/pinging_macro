package key_input

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

type KeyStruct struct {
	DisplayName string
	VirtualCode uint16
	IsSpecial   bool
	IsUpperCase bool // Empty in case of the KeysList var which holds
}

type KeysListType []KeyStruct

func (k KeysListType) FindKeyByDisplayName(displayName string) (KeyStruct, error) {
	displayNameUppercase := strings.ToUpper(displayName)

	for _, key := range k {
		if key.DisplayName == displayNameUppercase {
			key.IsUpperCase = displayName == displayNameUppercase
			return key, nil
		}
	}
	return KeyStruct{}, errors.New("cannot find key")
}


type KeyInputService interface {
	Type(text string) error
	Submit() error
	GetKeysList() KeysListType
}

func NewKeyInputService(logger *log.Entry) KeyInputService {
	if runtime.GOOS == "linux" {
		// Todo, I am counting on you!
	}
	return newWindowsKeyInputService(logger)
}