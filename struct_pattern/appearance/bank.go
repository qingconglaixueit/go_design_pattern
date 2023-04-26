// @Author Bing 
// @Desc
package main

import "fmt"

type Bank struct {
	Ac *Account
	Co *VerifyCode
	Re *Recorder
	No *NoticeSys
}

func NewBank(account, code string) *Bank {
	return &Bank{
		Ac: NewAccount(account),
		Co: NewVerifyCode(code),
		Re: NewRecorder(),
		No: NewNotice("[Bank business]"),
	}
}

func (b *Bank) SaveMoney(name, code string, money int) {
	// 检查账户
	if !b.Ac.CheckAccount(name){
		fmt.Errorf("check account fail ... ")
		return
	}
	// 检查验证码
	if !b.Co.CheckCode(code){
		fmt.Errorf("check code fail ... ")
		return
	}

	// 计算，并 写入数据
	sum := b.Re.GetData() + money
	b.Re.SaveData(sum)
	// 通知客户
	b.No.SendNotice(fmt.Sprintf("save money ok,now money is [%d]", sum))
}

func (b *Bank) WithdrawMoney(name, code string, money int) {
	// 检查账户
	if !b.Ac.CheckAccount(name){
		fmt.Println("check account fail ... ")
		return
	}
	// 检查验证码
	if !b.Co.CheckCode(code){
		fmt.Println("check code fail ... ")
		return
	}
	// 写入数据
	nowMoney :=b.Re.GetData()
	if nowMoney > money{
		nowMoney -= money
	}else{
		fmt.Println("Insufficient account balance ... ")
		return
	}
	b.Re.SaveData(nowMoney)
	// 通知客户
	b.No.SendNotice(fmt.Sprintf("Withdraw money ok,now money is [%d]", nowMoney))
}
