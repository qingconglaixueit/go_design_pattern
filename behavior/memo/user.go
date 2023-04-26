// @Author Bing 
// @Desc
package main

import "fmt"

type User struct {
	MemoList []*Memo
}

func (u *User) AddMemo(m *Memo) {
	u.MemoList = append(u.MemoList, m)
	fmt.Println("AddMemo state : ",m.State)
}

func (u *User) GetMemo(index int) *Memo {
	if index < len(u.MemoList) {
		return u.MemoList[index]
	}
	return nil
}
