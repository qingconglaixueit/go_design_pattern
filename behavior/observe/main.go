// @Author Bing 
// @Desc
package main


// 模拟 员工观察食堂，食堂做好菜之后 订阅的员工都能够收到通知

func main(){

	o1 := &Emploee{Id: "10001",Name: "xiaozhu"}
	o2 := &Emploee{Id: "10002",Name: "xiaozhu2"}
	o3 := &Emploee{Id: "10002",Name: "xiaozhu3"}

	c := &Cantee{Name: " big cantee "}
	c.Register(o1)
	c.Register(o2)
	c.Register(o3)

	// 开饭了
	c.NoticeAll()

}
