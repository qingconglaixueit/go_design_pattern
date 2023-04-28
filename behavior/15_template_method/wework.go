// @Author Bing 
// @Desc
package main

import "fmt"

type Wework struct {
	AutoMsg string
}

func (w *Wework) GetMsg() string {
	return w.AutoMsg
}
func (w *Wework) SaveMsg(string) {
	fmt.Printf("[%s] save msg to cache ...  \n", "Wework")
}
func (w *Wework) GenerateReply(msg string) string {
	return fmt.Sprintf("[%s] %s", "Wework", msg)
}
func (w *Wework) SendMsg(msg string) {
	fmt.Printf("send msg:[%s] to %s ... \n", msg, "Wework")
}
