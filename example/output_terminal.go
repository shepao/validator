package example

import (
	"github.com/shepao/valid/rule"
	"log"
)

func Println(v ...interface{}) {
	if rule.Model {
		log.Println(v...)
	}
}
