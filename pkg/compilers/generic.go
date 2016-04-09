package compilers

import unikos "github.com/emc-advanced-dev/unik/pkg/os"
import "io/ioutil"

func BuildBootableImage(kernel, cmdline string) (string, error) {
	rootFile, err := ioutil.TempFile("", "")
	if err != nil {
		return "", err
	}
	rootFileName := rootFile.Name()
	rootFile.Close()

	size := unikos.MegaBytes(100)

	if err := unikos.CreateBootImageWithSize(rootFileName, kernel, cmdline, size); err != nil {
		return "", err
	}

	return rootFileName, nil
}
