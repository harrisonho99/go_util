package directory

import (
	"os"
)

func GetCurrentDir() (currentWorkspace string, err error) {
	return os.Getwd()
}
