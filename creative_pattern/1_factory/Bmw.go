// @Author Bing
// @Desc
package main

// bmw

type Bmw struct{
	Car
}

func newBmw() ICar{
	return &Bmw{
		Car:Car{
			Name: "Bmw",
			Color: "black",
		},
	}
}