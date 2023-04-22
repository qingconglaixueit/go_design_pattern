// @Author Bing
// @Desc
package main

import "fmt"

// 车

type Car struct {
	Name  string
	Color string
}

func (c *Car) Run() {
	fmt.Printf(" %s is running ...\n", c.Name)

}
func (c *Car) SetName(name string) {
	c.Name = name
}
func (c *Car) GetName() string {
	return c.Name
}
func (c *Car) SetColor(color string) {
	c.Color = color
}
func (c *Car) GetColor() string {
	return c.Color
}
