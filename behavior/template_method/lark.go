// @Author Bing 
// @Desc
package main

import "fmt"

type Lark struct {
	AutoMsg string
}

func (w *Lark) GetMsg() string {
	return w.AutoMsg
}
func (w *Lark) SaveMsg(string) {
	fmt.Printf("[%s] save msg to cache ...  \n", "Lark")
}
func (w *Lark) GenerateReply(msg string) string {
	return fmt.Sprintf("[%s] %s", "Lark", msg)
}
func (w *Lark) SendMsg(msg string) {
	fmt.Printf("send msg:[%s] to %s ... \n", msg, "Lark")
}
