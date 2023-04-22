// @Author Bing 
// @Desc
package main

type Manager struct {
	Worker IBuild
}

func NewManager(b IBuild) *Manager {
	return &Manager{Worker: b}
}
func (m *Manager) SetWorker(b IBuild) {
	m.Worker = b
}
// @Title GetHouse
// @Description  建一个房子的步骤
// @Return *House
func (m *Manager) GetHouse() *House {
	m.Worker.SetFrame()

	m.Worker.SetStyle()

	m.Worker.SetDoor()

	m.Worker.SetBed()

	return m.Worker.BuildHouse()
}
