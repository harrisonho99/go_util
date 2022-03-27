package file_test

import (
	"testing"

	"github.com/hotsnow199/go_util/directory"
	"github.com/hotsnow199/go_util/util"
)

func TestWatchFile(t *testing.T) {
	currentWorkspace, err := directory.GetCurrentDir()
	util.CheckTestError(t, err)
	t.Log(currentWorkspace)
}
