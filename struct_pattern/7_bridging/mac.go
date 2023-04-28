// @Author Bing 
// @Desc
package main

type Mac struct {
	Model string
	Dis IDisplay
}

func (m *Mac) SetDisplay(dis IDisplay) {
	m.Dis = dis
}

func (m *Mac) Show() {
	m.Dis.Show("mac")
}
