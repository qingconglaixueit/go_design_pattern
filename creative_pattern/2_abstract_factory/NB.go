// @Author Bing 
// @Desc
package main

type NB struct {
}

func (NB) MakeShoe() IShoe {
	return  Shoe{
		Brand: "NB",
		Color: "black",
		Size: 40,
	}
}
func (NB) MakeCloth() ICloth {
	return Cloth{
		Brand: "NB",
		Color: "white",
		Size: "M",
	}
}
