// @Author Bing 
// @Desc
package main

type Memo struct {
	State string
}

func (m *Memo) GetMemoState() string {
	return m.State
}
