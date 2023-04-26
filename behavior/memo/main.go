// @Author Bing 
// @Desc
package main

import "fmt"

// 备忘录模式，可以记录快照，在某个时候将快照取出
// 例如某白领使用备忘录，记录了 1 2 3 条信息，然后想调出 2 快照来进行查看
func main() {

	u := &User{MemoList: make([]*Memo, 0)}

	o := &Originator{}

	o.SetState("1")
	u.AddMemo(o.NewMemo(o.GetState()))


	o.SetState("2")
	u.AddMemo(o.NewMemo(o.GetState()))


	o.SetState("3")
	u.AddMemo(o.NewMemo(o.GetState()))


	fmt.Println(u.GetMemo(1))
}
