package sync_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/hotsnow199/go_util/sync"
)

var ListURL = []string{"https://google.com", "https://twitter.com", "https://go.dev"}

func MockHTTPGet(url string, t *testing.T) int {
	resp, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}

	nBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	return len(nBytes)
}

func TestWaiter(t *testing.T) {

	var w sync.Waiter

	for _, url := range ListURL {
		w.Add(1)
		go func(url string) {
			defer w.Done()
			MockHTTPGet(url, t)
			fmt.Println("received: ", MockHTTPGet(url, t), " bytes")
		}(url)
	}

	w.Wait()
	t.Log("Wait Done.")
}

func TestWaiter1(t *testing.T) {

	var w sync.Waiter

	w.Add(3)
	for _, url := range ListURL {
		go func(url string) {
			defer w.Done()
			MockHTTPGet(url, t)
			fmt.Println("received: ", MockHTTPGet(url, t), " bytes")
		}(url)
	}
	w.Wait()
	t.Log("Wait Done.")
}

func TestWaiterShouldDeadlock(t *testing.T) {

	var w sync.Waiter

	w.Add(4)
	for _, url := range ListURL {
		go func(url string) {
			defer w.Done()
			MockHTTPGet(url, t)
			fmt.Println("received: ", MockHTTPGet(url, t), " bytes")
		}(url)
	}
	w.Wait()
	t.Log("Wait Done.")
}

func TestWaiterShouldErrror(t *testing.T) {

	var w sync.Waiter

	w.Add(2)
	for _, url := range ListURL {
		go func(url string) {
			defer w.Done()
			fmt.Println("received: ", MockHTTPGet(url, t), " bytes")
		}(url)
	}
	w.Wait()
	time.Sleep(time.Second * 2)
	t.Log("Wait Done.")
}
