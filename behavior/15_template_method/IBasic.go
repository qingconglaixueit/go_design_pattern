// @Author Bing 
// @Desc
package main

type IBasic interface{
	GetMsg()string
	SaveMsg(string)
	GenerateReply(string)string
	SendMsg(string)
}


type Mgr struct{
	B IBasic
}

func (m * Mgr)GetAndSendMsg(){
	msg := m.B.GetMsg()
	m.B.SaveMsg(msg)
	reply := m.B.GenerateReply(msg)
	m.B.SendMsg(reply)
}