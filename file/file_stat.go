package file

import (
	"io/fs"
	"time"

	"github.com/hotsnow199/go_util/memory"
)

type FileStat struct {
	Name    string
	Mode    fs.FileMode
	Size    memory.MemorySize
	IsDir   bool
	ModTime time.Time
}
