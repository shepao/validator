package example

import (
	"github.com/shepao/valid"
)

type TestEmail struct {
	Email string `valid:"Email"`
}

func Test_Email(v valid.IValidation) {
	Println(v.Valid(&TestEmail{Email: "11@qq.com"}))
	Println(v.Valid(&TestEmail{Email: "11@qqcom"}))
}
