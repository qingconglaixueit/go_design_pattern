// @Author Bing 
// @Desc
package main

type UserArray struct {
	users []*User
}

func (u *UserArray) createIterator() Interator {
	return &UserIterator{Users: u.users}
}
