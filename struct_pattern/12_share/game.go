// @Author Bing 
// @Desc
package main

//
//  Game
//  @Description: 足球比赛，分成 A 队，B 队
//
type Game struct {
	APlayers []*Player
	BPlayers []*Player
}

func NewGame() *Game {
	return &Game{
		APlayers: make([]*Player, 0, 2),
		BPlayers: make([]*Player, 0, 2),
	}
}

func (g *Game) AddAPlayer() {
	g.APlayers = append(g.APlayers, NewPlayer(APlayer))
}

func (g *Game) AddBPlayer() {
	g.BPlayers = append(g.BPlayers, NewPlayer(BPlayer))
}
