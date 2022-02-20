package log

import (
	"fmt"
	"os"
)

var DefaultPath = "/var/log/cni/data.log"

func Write(args ...interface{}) {
	file, _ := os.Open(DefaultPath)
	defer file.Close()
	_, _ = file.Write([]byte(fmt.Sprintln(args)))
}
