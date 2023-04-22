// @Author Bing 
// @Desc
package main
// @Title main
// @Description
// 原始对象被称为上下文， 它包含指向策略对象的引用并将执行行为的任务分派给策略对象。
// 为了改变上下文完成其工作的方式， 其他对象可以使用另一个对象来替换当前链接的策略对象。
func main(){
	c := initCache(&Fifo{})

	c.Add("stu1","xiaozhu1")
	c.Add("stu2","xiaozhu2")
	c.Add("stu3","xiaozhu3")
	c.Add("stu4","xiaozhu4")

	c.SetEvi(&Lru{})
	c.Add("stu5","xiaozhu5")

	c.SetEvi(&Lfu{})
	c.Add("stu6","xiaozhu6")


}
