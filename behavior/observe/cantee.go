// @Author Bing 
// @Desc
package main

import "fmt"

type Cantee struct {
	Obs  []Observe
	Name string
}

func (c *Cantee) Register(o Observe) {
	c.Obs = append(c.Obs, o)
	fmt.Println("add new Observe  ",o.getEmId())
}
func (c *Cantee) Deregister(o Observe) {
	for index, v := range c.Obs {
		if v.getEmId() == o.getEmId() {
			c.Obs = append(c.Obs[:index], c.Obs[index+1:]...)
			fmt.Println("Deregister Observe  ", o.getEmId())
			return
		}
	}
}
func (c *Cantee) NoticeAll() {
	for _, v := range c.Obs {
		v.update()
	}
}
