// @Author Bing 
// @Desc
package main

type Originator struct {
	State string
}

func (o *Originator) NewMemo(s string) *Memo {
	return &Memo{State: s}
}

func (o *Originator) SaveMemo(m *Memo) {
	m.State = o.State
	return
}

func (o *Originator) GetState() string {
	return o.State
}

func (o *Originator) SetState(s string) {
	o.State = s
}
