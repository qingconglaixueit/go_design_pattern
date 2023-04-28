// @Author Bing 
// @Desc
package main

type User struct {
	Name string
	Age  int
}

type UserIterator struct {
	Users []*User
	Index int
}

func (u *UserIterator) HasNext() bool {
	if u.Index >= len(u.Users) {
		return false
	}
	return true
}

func (u *UserIterator) GetNext() *User {
	if u.Index >= len(u.Users) {
		return nil
	}
	res := u.Users[u.Index]
	u.Index++
	return res
}
