package directory

import (
	"io/fs"
	"io/ioutil"
)

func GetDirEntries(dirName string) (entries []fs.FileInfo, err error) {
	return ioutil.ReadDir(dirName)
}
