// @Author Bing 
// @Desc
package main

import "fmt"

type App struct {
}

func (a *App) HandReq(url string, method string) (int, string) {

	fmt.Println("url ==", url, "  method ==", method)

	if url == "/get_info" && method == "GET" {
		return 200, "success"
	}

	if url == "/edit_info" && method == "POST" {
		return 200, "success"
	}

	return 400, "failed  ..."

}
