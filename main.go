package main

import "fmt"

func main() {
	// useWatchFile()
	ch := make(chan int)
	go func() {
		close(ch)
	}()
	for range ch {
	}
	var sl = []int{1, 2, 3, 4}
	fmt.Println(sl[1:])
}
