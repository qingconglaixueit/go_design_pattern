// @Author Bing 
// @Desc
package main
// @Title main
// @Description  以插件的方式来注入到程序中
func main() {

	onB := &Button{}
	offB := &Button{}

	tv := &Tv{}

	onCmd := &OnCmd{Dev: tv}
	offCmd := &OffCmd{Dev: tv}

	onB.Cmd = onCmd
	onB.Press()

	offB.Cmd = offCmd
	offB.Press()

}
