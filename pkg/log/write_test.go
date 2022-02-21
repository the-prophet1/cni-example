package log

import "testing"

func TestWrite(t *testing.T) {
	DefaultPath = "./hello.log"
	for i := 0; i < 10; i++ {
		Write("hello")
	}
}
