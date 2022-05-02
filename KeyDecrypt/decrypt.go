// just load this to cyberchef
/*AES_Decrypt({'option':'UTF8','string':'thisis32bitlongpassphraseimusing'},{'option':'UTF8','string':''},'ECB/NoPadding','Hex','Raw',{'option':'Hex','string':''},{'option':'Hex','string':''})
Remove_whitespace(true,true,true,true,true,false)*/
//or
//https://gchq.github.io/CyberChef/#recipe=AES_Decrypt(%7B'option':'UTF8','string':'thisis32bitlongpassphraseimusing'%7D,%7B'option':'UTF8','string':''%7D,'ECB/NoPadding','Hex','Raw',%7B'option':'Hex','string':''%7D,%7B'option':'Hex','string':''%7D)Remove_whitespace(true,true,true,true,true,false)

package main

import (
	"bufio"
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// cipher key
	key := "thisis32bitlongpassphraseimusing"

	file, err := os.Open("/bin/keys")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var strip string = scanner.Text()

		strip = strings.Replace(strip, "\n", "", -1)
		// fmt.Println(strip)
		DecryptAES([]byte(key), strip)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func DecryptAES(key []byte, ct string) {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	fmt.Println("DECRYPTED:", s)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
