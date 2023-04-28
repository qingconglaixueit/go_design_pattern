// @Author Bing 
// @Desc
package main

//  对于我们生成 源 msg，期望放到缓存中，生成具体的回复信息，将消息发送出去
// 对于这种步骤固定，但是对于每一步骤期望有自己的实现方式的，期望去重写具体的实现方式的，可以使用模板方法
// 获取到信息之后，wechat ，wework，lark 有不同的处理方式，但是他们的整体处理结构是一样的
func main(){

	m := &Mgr{}

	// wechat
	m.B = &Wechat{AutoMsg: "你好呀啊，这是 gpt 生成的哦 ..."}
	m.GetAndSendMsg()

	// wework
	m.B = &Wework{AutoMsg: "我是企业微信小冰 ..."}
	m.GetAndSendMsg()


	// lark
	m.B = &Lark{AutoMsg: "我是 飞书小哥 ..."}
	m.GetAndSendMsg()
}
