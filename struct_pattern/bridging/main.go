// @Author Bing 
// @Desc
package main

// 桥接模式
// 可将业务逻辑或一个大类拆分为不同的层次结构， 从而能独立地进行开发
// 例如不同的显示器要接入计算机，那么计算机可以独立开发，显示器也可以独立开发
// 无论计算机里面有 windows 的，还是 mac 的 仍然支持设置 显示器

func main() {

	p := &Philips{}
	x := &Xiaomi{}

	m := &Mac{Model: "mac-pro"}
	w := &Win{Model: "thinkpad"}

	m.SetDisplay(p)
	m.Show()

	m.SetDisplay(x)
	m.Show()


	w.SetDisplay(p)
	w.Show()

	w.SetDisplay(x)
	w.Show()




}
