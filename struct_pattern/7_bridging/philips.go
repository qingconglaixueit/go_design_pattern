// @Author Bing 
// @Desc
package main

import "fmt"

type Philips struct {
}

func (p *Philips) Show(content string) {
	fmt.Printf("[%s]Philips show ... \n",content)
}
