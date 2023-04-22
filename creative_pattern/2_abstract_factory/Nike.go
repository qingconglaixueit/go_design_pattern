// @Author Bing 
// @Desc
package main

type Nike struct {
}

func (Nike) MakeShoe() IShoe {
	return Shoe{
		Brand: "Nike",
		Color: "brown",
		Size: 36,
	}
}
func (Nike) MakeCloth() ICloth {
	return Cloth{
		Brand: "Nike",
		Color: "purple",
		Size: "L",
	}
}
