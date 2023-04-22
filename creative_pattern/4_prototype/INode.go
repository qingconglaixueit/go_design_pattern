// @Author Bing 
// @Desc
package main

type INode interface{
	Show(string)
	Clone()INode
}
