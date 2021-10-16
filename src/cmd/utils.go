package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Luzifer/go-openssl/v4"
)

// -------- -------- -------- -------- -------- -------- -------- -------- -------- --------
func fileRead(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		return nil, err
	}
	return b, nil
}

// -------- -------- -------- -------- -------- -------- -------- -------- -------- --------
func fileWrite(path string, data []byte) error {
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// -------- -------- -------- -------- -------- -------- -------- -------- -------- --------
func checkPathAreSame(pathFirst string, pathSecond string) bool {
	stat1, err := os.Stat(pathFirst)
	if err != nil {
		fmt.Printf("Unable to Read File: %s\n", pathFirst)
	}
	stat2, err := os.Stat(pathSecond)
	if err != nil {
		fmt.Printf("Unable to Read File: %s\n", pathSecond)
	}

	return os.SameFile(stat1, stat2)

}

// -------- -------- -------- -------- -------- -------- -------- -------- -------- --------
func encryptFile(path string, password string, pathout string) error {

	if checkPathAreSame(filein, fileout) {
		return errors.New("file in and file out cannot be same")
	}
	o := openssl.New()
	data, err := fileRead(path)
	fmt.Println(string(data))
	if err != nil {
		return err
	}
	enc, err := o.EncryptBytes(password, data, openssl.PBKDF2SHA256)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
	}

	fmt.Printf("Encrypted text: %s\n", string(enc))

	err = fileWrite(pathout, enc)
	if err != nil {
		return err
	}
	return nil
}

// -------- -------- -------- -------- -------- -------- -------- -------- -------- --------
func decryptFile(path string, password string, pathout string) error {

	if checkPathAreSame(filein, fileout) {
		return errors.New("file in and file out cannot be same")
	}
	o := openssl.New()
	data, err := fileRead(path)
	fmt.Println(string(data))
	if err != nil {
		return err
	}
	enc, err := o.DecryptBytes(password, data, openssl.PBKDF2SHA256)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
	}

	fmt.Printf("Decrypted text: %s\n", string(enc))

	err = fileWrite(pathout, enc)
	if err != nil {
		return err
	}
	return nil
}

// -------- -------- -------- -------- -------- -------- -------- -------- -------- --------
