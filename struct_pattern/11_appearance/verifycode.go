// @Author Bing 
// @Desc
package main

type VerifyCode struct {
	Code string
}

func NewVerifyCode(code string) *VerifyCode {
	return &VerifyCode{Code: code}
}

func (v *VerifyCode) CheckCode(code string) bool {
	return v.Code == code
}
