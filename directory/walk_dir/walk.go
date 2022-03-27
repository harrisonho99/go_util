package walk_dir

import (
	"io/fs"
	"path"

	"github.com/hotsnow199/go_util/directory"
	"github.com/hotsnow199/go_util/file"
	"github.com/hotsnow199/go_util/memory"
)

func WalkConcurent(dirPath string, fileChan chan *file.FileStat, getDirInfo bool) error {
	listEntries, err := directory.GetDirEntries(dirPath)
	if err != nil {
		return err
	}

	for _, entry := range listEntries {
		if entry.IsDir() {
			//some case client want dir status
			if getDirInfo {
				fileChan <- fileStatConstructor(entry)
			}
			//get sub-dir path and recursive walk
			subDirPath := path.Join(dirPath, entry.Name())
			WalkConcurent(subDirPath, fileChan, getDirInfo)
		} else {
			//base-case in recursive walk
			fileChan <- fileStatConstructor(entry)
		}
	}

	return nil
}

func fileStatConstructor(f fs.FileInfo) *file.FileStat {
	return &file.FileStat{
		Name:  f.Name(),
		Mode:  f.Mode(),
		Size:  memory.MemorySize(f.Size()),
		IsDir: f.IsDir(), ModTime: f.ModTime()}
}
