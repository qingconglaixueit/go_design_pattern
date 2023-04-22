// @Author Bing 
// @Desc
package main

type ICloth interface {
	SetColor(string)
	SetSize(string)
	GetColor() string
	GetSize() string
	GetBrand()string
}

type Cloth struct {
	Color string
	Size  string
	Brand string
}

func (c Cloth) SetColor(col string) {
	c.Color = col
}
func (c Cloth) SetSize(size string) {
	c.Size = size
}
func (c Cloth) GetColor() string {
	return c.Color
}
func (c Cloth) GetSize() string {
	return c.Size
}
func (c Cloth) GetBrand() string {
	return c.Brand
}