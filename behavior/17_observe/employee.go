// @Author Bing 
// @Desc
package main

import "fmt"

type Emploee struct {
	Name string
	Id   string
}

func (e *Emploee) update() {
	fmt.Println("Send sms to ", e.Name)
}
func (e *Emploee) getEmId() string {
	fmt.Printf(" %s getEmId\n", e.Name)
	return e.Id
}
