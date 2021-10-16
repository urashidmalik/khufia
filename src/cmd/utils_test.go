package cmd

import (
	"io/ioutil"
	"os"
	"testing"
)

func createTestFile(fileName string, data string) {
	d1 := []byte(data)
	err := os.WriteFile(fileName, d1, 0644)
	check(err)
}
func removeTestFile(fileName string) {
	os.Remove(fileName)
}

// -------- -------- -------- -------- -------- -------- -------- -------- -------- --------
func TestIfFileReadCanDetectFileDoesNotExist(t *testing.T) {
	createTestFile("dat1", "hello go")
	_, err := fileRead("dat2")
	if err == nil {
		t.Error("Unable to detect missing file")
	}
	removeTestFile("dat1")
}

// -------- -------- -------- -------- -------- -------- -------- -------- -------- --------
func TestSamefileDetect(t *testing.T) {
	createTestFile("dat1", "hello go")
	if !checkPathAreSame("dat1", "dat1") {
		t.Error("Failed Unable to detect same file")
	}

	defer os.Remove("dat1")
}

// -------- -------- -------- -------- -------- -------- -------- -------- -------- --------
func TestEncryptInputfileOutoutFileNotSame(t *testing.T) {
	createTestFile("dat1", "hello go")
	err := encryptFile("dat1", "sd", "dat1")
	if err != nil {
		if err.Error() != "file in and file out cannot be same" {
			t.Error("Failed: FileIn and FileOut are same but system was not able to detect it")
			t.Error(err.Error())
		}
	}
	removeTestFile("dat1")
}

// -------- -------- -------- -------- -------- -------- -------- -------- -------- --------
func TestEncryptFailTest(t *testing.T) {
	createTestFile("dat1", "hello go")
	err := encryptFile("dat11", "sd", "dat1.goenc")
	if err == nil {
		t.Error("Failed: File does not exist system unable to detect")
	}
	// removeTestFile("dat1")
}

// -------- -------- -------- -------- -------- -------- -------- -------- -------- --------
func TestDecryptInputfileOutoutFileNotSame(t *testing.T) {
	createTestFile("dat1", "hello go")
	err := decryptFile("dat1", "sd", "dat1")
	if err != nil {
		if err.Error() != "file in and file out cannot be same" {
			t.Error("Failed: FileIn and FileOut are same but system was not able to detect it")
			t.Error(err.Error())
		}
	}
	removeTestFile("dat1")
}

// -------- -------- -------- -------- -------- -------- -------- -------- -------- --------

func TestEncryptionDecryptionTest(t *testing.T) {
	initData := "hello go"
	createTestFile("dat1", initData)

	err := encryptFile("dat1", "sd", "dat1.goenc")
	if err != nil {
		t.Error(err.Error())
	}
	err = decryptFile("dat1.goenc", "sd", "dat1.plain")
	if err != nil {
		t.Error(err.Error())
	}
	dataB, err := ioutil.ReadFile("dat1.plain")
	if err != nil {
		t.Error(err.Error())
	}
	if string(dataB) != initData {
		t.Error("Data corrupted in Encryption/Decryption")
	}
	removeTestFile("dat1")
	removeTestFile("dat1.goenc")
	removeTestFile("dat1.plain")
}

// -------- -------- -------- -------- -------- -------- -------- -------- -------- --------
func check(e error) {
	if e != nil {
		panic(e)
	}
}
