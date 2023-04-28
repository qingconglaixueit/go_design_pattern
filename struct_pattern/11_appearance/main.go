// @Author Bing 
// @Desc
package main

// 咱们去银行存钱，取钱，表面上看着很简单，一进一出即可，殊不知银行后面的内部系统之间复杂的交互，我们简化一下
// 例如银行后面涉及到 账户系统，检测系统，数据库系统，通知系统
// 这个时候 就可以使用 外观模式，对于客户端来说使用非常简单，无需关注其内部复杂系统
func main(){
	// 新建账号
	b := NewBank("xiaozhu","123456")

	// 存钱
	b.SaveMoney("xiaozhu","123456",1000)

	// 取钱
	b.WithdrawMoney("xiaozhu","123456",500)

	// 取钱
	b.WithdrawMoney("xiaozhu","123456",600)
}
