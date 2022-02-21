package log

import (
	"fmt"
	"io/ioutil"
)

var DefaultPath = "/var/log/cni/cni.log"

func Write(args ...interface{}) {
	if err := ioutil.WriteFile(DefaultPath, []byte(fmt.Sprintln(args)), 0666); err != nil {
		fmt.Println(err)
	}
}
