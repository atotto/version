package text

import (
	"bufio"
	"fmt"
	"os"
)

func Version(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	data, flagment, err := r.ReadLine()
	if flagment {
		return "", fmt.Errorf("too long: %s", filename)
	}
	if err != nil {
		return "", err
	}
	return string(data), nil
}
