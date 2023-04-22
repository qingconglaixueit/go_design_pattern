// @Author Bing 
// @Desc
package main

import "fmt"
// @Title main
// @Description  设计一个足球比赛游戏，A 队队服是 红色，B 队队服是 蓝色，且整个队伍就只有一个套衣服 ，此处就应用了享元模式
func main() {

	g := NewGame()

	g.AddAPlayer()
	g.AddAPlayer()
	g.AddAPlayer()

	g.AddBPlayer()
	g.AddBPlayer()
	g.AddBPlayer()
	g.AddBPlayer()

	fmt.Printf("APlayer:[%d]\n", len(g.APlayers))
	fmt.Printf("BPlayer:[%d]\n", len(g.BPlayers))

	dressIn := getDressInstance()
	for k, v := range dressIn.Dm {
		fmt.Printf("%s : [%s]\n", k, v.getColor())
	}

	return
}
