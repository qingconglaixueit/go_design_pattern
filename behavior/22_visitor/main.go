// @Author Bing 
// @Desc
package main
// @Title main
// @Description  目前有 3 种结构， Dog，Cat，Monkey ，访问者有 AreaVs，VolumeVs，ColorVs ，
// 以后如果需要类似于增加计算体脂率的，那么就可以加一个 visitor 即可，且对其他的访问者和结构没有影响
// 如果需要增加结构，例如增加 pig ，那么直接去实现 Animal 即可，对其他的结构和访问者没有影响

func main() {

	dog := &Dog{}
	cat := &Cat{}
	monkey := &Monkey{}

	vs := &AreaVs{}

	dog.Accept(vs)
	cat.Accept(vs)
	monkey.Accept(vs)



}
