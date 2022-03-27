package util

import "log"

func CheckErrorAndPanic(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
