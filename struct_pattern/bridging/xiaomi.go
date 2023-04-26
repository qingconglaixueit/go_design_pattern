// @Author Bing 
// @Desc
package main

import "fmt"

type Xiaomi struct {
}

func (p *Xiaomi) Show(content string) {
	fmt.Printf("[%s]Xiaomi show ... \n",content)
}