package file

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/hotsnow199/go_util/util"
)

type FileOp = fsnotify.Op
type Event = fsnotify.Event

const (
	Create FileOp = 1 << iota
	Write
	Remove
	Rename
	Chmod
)

type FileEventEmitter interface {
	On(Event)
}

type FileEvent struct {
	OnCreate func(Event)
	OnWrite  func(Event)
	OnRemove func(Event)
	OnRename func(Event)
	OnChmod  func(Event)
}

func (fe *FileEvent) On(ev Event) {
	switch ev.Op {
	case Create:
		if fe.OnCreate != nil {
			fe.OnCreate(ev)
		}
	case Write:
		if fe.OnWrite != nil {
			fe.OnWrite(ev)
		}
	case Remove:
		if fe.OnRemove != nil {
			fe.OnRemove(ev)
		}
	case Rename:
		if fe.OnRename != nil {
			fe.OnRename(ev)
		}
	case Chmod:
		if fe.OnChmod != nil {
			fe.OnChmod(ev)
		}
	}
}

func WatchFile(eventConsumer FileEventEmitter, done util.Done, filePaths ...string) (err error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	// close watcher
	defer watcher.Close()
	abort := util.NewDoneChan()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				eventConsumer.On(event)

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Fprintf(os.Stderr, "Error:: %s", err)
			case <-abort:
				return
			}
		}
	}()

	for _, fp := range filePaths {
		err = watcher.Add(fp)
		if err != nil {
			return err
		}
	}
	<-done
	abort <- util.WorkDone
	return nil
}
