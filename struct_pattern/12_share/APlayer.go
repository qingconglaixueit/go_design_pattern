// @Author Bing 
// @Desc
package main

type SAPlayer struct {
	color string
}

func (a *SAPlayer) getColor() string {

	return a.color
}

func NewAPlayer() *SAPlayer{
	return &SAPlayer{color: "red"}
}