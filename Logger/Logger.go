package Logger

import "log"

func CheckError(err error, msg string) {

	if err != nil {
		log.Fatal(msg, err)
	}
}
