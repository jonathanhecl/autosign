package autosign

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	signature = "autosignature"
)

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(0)
	}
}

func Init(salt ...[]byte) string {

	var err error = nil
	var needSign bool = true
	var sign_found []byte
	var byte_sign []byte = []byte(signature)
	var size_sign int = 40 + len(byte_sign)
	var executable string

	executable, err = os.Executable()
	checkError(err)

	file, err := ioutil.ReadFile(executable)
	checkError(err)

	if bytes.Equal(byte_sign, file[len(file)-size_sign:][:len(byte_sign)]) {

		needSign = false
		sign_found = file[len(file)-(size_sign-len(byte_sign)):]

	}

	hash := sha1.New()
	if needSign == true {

		_, err = hash.Write(file)

	} else {

		_, err = hash.Write(file[:len(file)-(size_sign)])

	}
	checkError(err)

	if len(salt) > 0 {
		_, err = hash.Write(salt[0])
		checkError(err)
	}

	hashInBytes := hash.Sum(nil)[:]
	shaFile := make([]byte, 40)
	hex.Encode(shaFile, hashInBytes)

	if needSign == true {

		var directory string
		var filename string
		var executable_signed string

		directory, filename = filepath.Split(executable)
		filename = filename[0:len(filename)-len(filepath.Ext(filename))] + "_" + string(shaFile[:]) + filepath.Ext(filename)
		executable_signed = directory + filename

		var _, err = os.Stat(executable_signed)

		if os.IsNotExist(err) {

			data, err := ioutil.ReadFile(executable)
			checkError(err)

			err = ioutil.WriteFile(executable_signed, append(data, append(byte_sign, shaFile...)...), 0644)
			checkError(err)
		}

	} else {

		if !bytes.Equal(shaFile, sign_found) {
			checkError(errors.New("File corrupted!"))
		}

	}

	return string(shaFile[:])

}
