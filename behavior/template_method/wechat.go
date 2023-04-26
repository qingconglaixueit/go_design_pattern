// @Author Bing 
// @Desc
package main

import "fmt"

type Wechat struct {
	AutoMsg string
}

func (w *Wechat) GetMsg() string {
	return w.AutoMsg
}
func (w *Wechat) SaveMsg(string) {
	fmt.Printf("[%s] save msg to cache ...  \n", "wechat")
}
func (w *Wechat) GenerateReply(msg string) string {
	return fmt.Sprintf("[%s] %s", "wechat", msg)
}
func (w *Wechat) SendMsg(msg string) {
	fmt.Printf("send msg:[%s] to %s ... \n", msg, "wechat")
}
