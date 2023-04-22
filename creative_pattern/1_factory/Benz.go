// @Author Bing
// @Desc
package main

// Benz

type Benz struct {
	Car
}

func newBenz() ICar {
	return &Benz{
		Car: Car{
			Name:  "Benz",
			Color: "white",
		},
	}
}
