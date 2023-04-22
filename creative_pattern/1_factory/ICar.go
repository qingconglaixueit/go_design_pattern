// @Author Bing
// @Desc
package main

// ICar 接口
type ICar interface{
	Run()
	SetName(string)
	GetName()string
	SetColor(string)
	GetColor()string
}
