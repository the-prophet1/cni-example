package log

import (
	"fmt"
	"os"
)

var DefaultPath = "/var/log/cni/cni.log"

func Write(args ...interface{}) {

	file, err := os.OpenFile(DefaultPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	if _, err := file.WriteString(fmt.Sprintln(args...)); err != nil {
		fmt.Println(err)
	}
}
