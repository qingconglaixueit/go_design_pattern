// @Author Bing 
// @Desc
package main

type Wood struct {
	Frame string
	Style string
	Door  string
	Bed   string
}

func NewWood()*Wood{
	return &Wood{}
}

func (w *Wood) SetFrame() {
	w.Frame = "Wood"
}
func (w *Wood) SetStyle() {
	w.Style = "black"
}
func (w *Wood) SetDoor() {
	w.Door = "Wood door"
}
func (w *Wood) SetBed() {
	w.Bed = "Wood bed"
}
func (w *Wood) BuildHouse() *House {
	return &House{
		Frame: w.Frame,
		Style: w.Style,
		Door:  w.Door,
		Bed:   w.Bed,
	}
}
