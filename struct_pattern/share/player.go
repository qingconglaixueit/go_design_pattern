// @Author Bing 
// @Desc
package main

type Player struct {
	pType    string
	dress    Dress
}

func NewPlayer(pType string) *Player {
	d, _ := getDressInstance().getDressType(pType)
	//if d == nil{
	//
	//}

	return &Player{
		pType: pType,
		dress: d,
	}
}
