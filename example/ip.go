package example

import (
	"github.com/shepao/valid"
)

type TestIp struct {
	Ip string `valid:"IP"`
}

func Test_Ip(v valid.IValidation) {
	Println(v.Valid(&TestIp{Ip: "218.88.124.41"}))
	Println(v.Valid(&TestIp{Ip: "218.256.124.41"}))
}
