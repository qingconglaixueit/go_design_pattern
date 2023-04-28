// @Author Bing 
// @Desc
package main

import "fmt"

func main() {

	ng := NewNginxSvr()
	fmt.Println(ng.HandReq("/get_info", "GET"))
	fmt.Println(ng.HandReq("/edit_info", "POST"))
	fmt.Println(ng.HandReq("/edit_info", "POST"))
	fmt.Println(ng.HandReq("/edit_infoxxx", "POST"))
}
