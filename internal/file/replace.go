package file

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

func Replace(filename string, old, new string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	data := bytes.Replace(buf, []byte(old), []byte(new), 1)

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err

}
