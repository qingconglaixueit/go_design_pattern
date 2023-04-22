// @Author Bing 
// @Desc
package main

import "fmt"

type Tv struct {
	Switch string
}

func (t *Tv) on() {

	t.Switch = "enable"
	fmt.Println("tv on ... ")

}
func (t *Tv) off() {
	t.Switch = "disable"
	fmt.Println("tv off ... ")
}
