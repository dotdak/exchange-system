package infrastructure

import (
	"io/ioutil"
	"log"
)

func NewLogger() *log.Logger {
	return log.Default()
}

func NullLogger() *log.Logger {
	return log.New(ioutil.Discard, "", 0)
}
