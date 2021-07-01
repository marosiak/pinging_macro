package key_input

import (
	log "github.com/sirupsen/logrus"
	"syscall"
	"unsafe"
)

var (
	user32        = syscall.NewLazyDLL("user32.dll")
	sendInputProc = user32.NewProc("SendInput")
)

type InternalKeyInputService struct {
	log *log.Entry
}

type keyboardInput struct {
	wVk         uint16
	wScan       uint16
	dwFlags     uint32
	time        uint32
	dwExtraInfo uint64
}

type input struct {
	inputType uint32
	ki        keyboardInput
	padding   uint64
}

func newWindowsKeyInputService(logger *log.Entry) KeyInputService {
	return &InternalKeyInputService{
		log: logger,
	}
}

func (s *InternalKeyInputService) findKeys(text string) (KeysListType, error) {
	var keys KeysListType

	for _, char := range text {
		key, err := keysList.FindKeyByDisplayName(string(char))

		if err != nil {
			s.log.WithField("key", char).WithError(err).Error("cannot translate key to code")
			return nil, err
		}
		keys = append(keys, key)
	}
	return keys, nil
}


func (s *InternalKeyInputService) keyStroke(keycode uint16, state keyState) {
	var i input
	i.inputType = 1 // INPUT_KEYBOARD

	i.ki.wScan = 0 // hardware scan code for key
	i.ki.time = 0
	i.ki.dwExtraInfo = 0

	// Press the "A" key
	i.ki.wVk = keycode
	i.ki.dwFlags = uint32(state) // 0 for key press

	_, _, _ = sendInputProc.Call(
		uintptr(1),
		uintptr(unsafe.Pointer(&i)),
		unsafe.Sizeof(i),
	)
}

type keyState uint32
const (
	Press keyState = iota
	Release keyState = 0x0002
)

func (s *InternalKeyInputService) release(keycode uint16) {
	s.keyStroke(keycode, Release)
}

func (s *InternalKeyInputService) press(keycode uint16) {
	s.keyStroke(keycode, Press)
}

func (s *InternalKeyInputService) Submit() error {
	key, err := keysList.FindKeyByDisplayName("ENTER")
	if err != nil {
		s.log.WithError(err).Error("cannot find ENTER keycode")
		return err
	}
	s.press(key.VirtualCode)
	s.release(key.VirtualCode)
	return nil
}

func (s *InternalKeyInputService) GetKeysList() KeysListType {
	return keysList
}

func (s *InternalKeyInputService) Type(text string) error {
	keys, err := s.findKeys(text)
	if err != nil {
		return err
	}

	for _, key := range keys {
		if key.IsUpperCase || key.IsSpecial {
			s.press(shift)
		}
		s.press(key.VirtualCode)
		s.release(key.VirtualCode)
		s.release(shift)
	}

	return nil
}

var shift uint16 = 0x10
var keysList = KeysListType{
	{DisplayName: " ", VirtualCode: 0x20},
	{DisplayName: "0", VirtualCode: 0x30},
	{DisplayName: "1", VirtualCode: 0x31},
	{DisplayName: "2", VirtualCode: 0x32},
	{DisplayName: "3", VirtualCode: 0x33},
	{DisplayName: "4", VirtualCode: 0x34},
	{DisplayName: "5", VirtualCode: 0x35},
	{DisplayName: "6", VirtualCode: 0x36},
	{DisplayName: "7", VirtualCode: 0x37},
	{DisplayName: "8", VirtualCode: 0x38},
	{DisplayName: "9", VirtualCode: 0x39},


	{DisplayName: "Q", VirtualCode: 0x51},
	{DisplayName: "W", VirtualCode: 0x57},
	{DisplayName: "E", VirtualCode: 0x45},
	{DisplayName: "R", VirtualCode: 0x52},
	{DisplayName: "T", VirtualCode: 0x54},
	{DisplayName: "Y", VirtualCode: 0x59},
	{DisplayName: "U", VirtualCode: 0x55},
	{DisplayName: "I", VirtualCode: 0x49},
	{DisplayName: "O", VirtualCode: 0x4F},
	{DisplayName: "P", VirtualCode: 0x50},
	{DisplayName: "A", VirtualCode: 0x41},
	{DisplayName: "S", VirtualCode: 0x53},
	{DisplayName: "D", VirtualCode: 0x44},
	{DisplayName: "F", VirtualCode: 0x46},
	{DisplayName: "G", VirtualCode: 0x47},
	{DisplayName: "H", VirtualCode: 0x48},
	{DisplayName: "J", VirtualCode: 0x4A},
	{DisplayName: "K", VirtualCode: 0x4B},
	{DisplayName: "L", VirtualCode: 0x4C},
	{DisplayName: "Z", VirtualCode: 0x5A},
	{DisplayName: "X", VirtualCode: 0x58},
	{DisplayName: "C", VirtualCode: 0x43},
	{DisplayName: "V", VirtualCode: 0x56},
	{DisplayName: "B", VirtualCode: 0x42},
	{DisplayName: "N", VirtualCode: 0x4E},
	{DisplayName: "M", VirtualCode: 0x4D},

	{DisplayName: "ENTER", VirtualCode: 0x0D},
	{DisplayName: "\n", VirtualCode: 0x0D},

	{DisplayName: "!", VirtualCode: 0x31, IsSpecial: true},
	{DisplayName: "@", VirtualCode: 0x32, IsSpecial: true},
	{DisplayName: "#", VirtualCode: 0x33, IsSpecial: true},
	{DisplayName: "$", VirtualCode: 0x34, IsSpecial: true},
	{DisplayName: "%", VirtualCode: 0x35, IsSpecial: true},
	{DisplayName: "^", VirtualCode: 0x36, IsSpecial: true},
	{DisplayName: "&", VirtualCode: 0x37, IsSpecial: true},
}