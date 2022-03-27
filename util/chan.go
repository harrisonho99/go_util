package util

var WorkDone = struct{}{}

type Done chan struct{}

func NewDoneChan() Done {
	return make(Done)
}
