// @Author Bing 
// @Desc
package main

type IShoe interface {
	SetColor(string)
	SetSize(int)
	GetColor() string
	GetSize() int
	GetBrand() string
}

type Shoe struct {
	Color string
	Size  int
	Brand string
}

func (s Shoe) SetColor( col string)  {
	s.Color = col
}
func (s Shoe) SetSize(size int)   {
	s.Size = size
}
func (s Shoe) GetColor() string {
	return s.Color
}
func (s Shoe) GetSize() int  {
	return s.Size
}
func (c Shoe) GetBrand() string {
	return c.Brand
}