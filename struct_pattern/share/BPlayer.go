// @Author Bing 
// @Desc
package main

type SBPlayer struct {
	color string
}

func (a *SBPlayer) getColor() string {

	return a.color
}

func NewBPlayer() *SBPlayer {
	return &SBPlayer{color: "blue"}
}
