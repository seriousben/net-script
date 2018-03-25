package binary

import (
	"bytes"
	"io/ioutil"
)

// IsBinaryBuffer checks if the buffer refers to binary data or not
func IsBinaryBuffer(buf []byte) (bool, error) {
	return bytes.Contains(buf, []byte("\x00")), nil
}

// IsBinaryFile checks if the file contains binary data or not
func IsBinaryFile(filepath string) (bool, error) {
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		return false, err
	}
	return IsBinaryBuffer(buf)
}
