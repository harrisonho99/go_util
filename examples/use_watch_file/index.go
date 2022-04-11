package use_watch_file

import (
	"fmt"
	"os"
	"path"

	"github.com/hotsnow199/go_util/directory"
	"github.com/hotsnow199/go_util/file"
	"github.com/hotsnow199/go_util/util"
)

const filename1 = "watch_file.txt"
const filename2 = "event.txt"

func UseWatchFile() {
	dir, err := directory.GetCurrentDir()
	util.CheckErrorAndPanic(err)
	filepath1 := path.Join(dir, "test", filename1)
	filepath2 := path.Join(dir, "test", filename2)

	done := InputCancelationCommand()
	EventConsumer := new(file.FileEvent)

	//attach event callback
	EventConsumer.OnWrite = OnWriteFile
	util.CheckErrorAndPanic(file.WatchFile(EventConsumer, done, filepath1, filepath2))

}

func InputCancelationCommand() util.Done {
	abort := util.NewDoneChan()
	go func() {
		fmt.Println("Press any key to cancel")
		buffer := make([]byte, 1)
		_, err := os.Stdin.Read(buffer)
		util.CheckErrorAndPanic(err)
		abort <- util.WorkDone
	}()

	return abort
}

func OnWriteFile(e file.Event) {
	fmt.Printf("write file: %s \n", e.Name)
}
