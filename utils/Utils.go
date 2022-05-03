package utils

import "log"

func HandleError(err error, message string) {
	if err != nil {
		log.Println(message)
		log.Println(err)
	}
}
