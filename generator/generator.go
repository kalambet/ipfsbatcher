package generator

import (
	"crypto/rand"
	"fmt"
	"os"
)

const (
	dataSize = 5242880
)

func NewTestData(paths []string) error {
	for _, p := range paths {
		f, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("error creating file %s: %s", p, err.Error())
		}
		buf := make([]byte, dataSize)
		_, err = rand.Read(buf)
		if err != nil {
			return fmt.Errorf("error reading from /dev/urandom for file %s: %s", p, err.Error())
		}

		_, err = f.Write(buf)
		if err != nil {
			return fmt.Errorf("error writing to file %s: %s", p, err.Error())
		}
		err = f.Close()
		if err != nil {
			return fmt.Errorf("error closing file %s descriptor: %s", p, err.Error())
		}
	}

	return nil
}
