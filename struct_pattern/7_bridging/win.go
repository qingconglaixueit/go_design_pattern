// @Author Bing 
// @Desc
package main
type Win struct {
	Model string
	Dis IDisplay
}

func (m *Win) SetDisplay(dis IDisplay) {
	m.Dis = dis
}

func (m *Win) Show() {
	m.Dis.Show("win")
}