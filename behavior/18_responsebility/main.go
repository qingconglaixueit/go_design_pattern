// @Author Bing 
// @Desc
package main

// @Description  员工入职可以看做是一个责任链， 前台报道，人事制度宣讲，合同部合同签订，送入具体部门  ,
// 这种部门职责明确，链路清晰，可以使用 责任链模式，允许你将请求沿着处理者链进行发送
// 对于新员工来说，只需要知道去去找前台报道即可
func main() {
	e := &Employee{Name: "xiaozhu"}

	// 前台报道，
	//人事制度宣讲
	//合同部合同签订
	//送入具体部门

	o := &OwnPart{}

	c := &Contract{}
	c.SetNext(o)

	h := &HrPart{}
	h.SetNext(c)

	f := &Front{}
	f.SetNext(h)
	f.Execute(e)

}
