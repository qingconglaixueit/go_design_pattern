// @Author Bing 
// @Desc
package main

import (
	"errors"
	"fmt"
)

// 用于在别人已有的接口上加一层代理，自己可以在代理中做自己想做的事情（例如校验某些数据，访问控制等等），做完之后再去调用真实的接口

type Server interface {
	HandReq(string, string) (int, string)
}

type Nginx struct {
	AppLi       *App
	MaxAllowReq int
	RateLimit   map[string]int
}

func NewNginxSvr() *Nginx {
	return &Nginx{
		AppLi:       &App{},
		MaxAllowReq: 1,
		RateLimit:   map[string]int{},
	}
}

func (n *Nginx) HandReq(url string, method string) (int, string) {
	if err := n.checkParams(url); err != nil {
		return 500, err.Error()
	}

	return n.AppLi.HandReq(url, method)
}

func (n *Nginx) checkParams(url string) error {
	if n.RateLimit[url] == 0 {
		n.RateLimit[url] = 1
	} else {
		n.RateLimit[url]++
	}

	if n.RateLimit[url] > n.MaxAllowReq {
		fmt.Println("Beyond the limit number ... ")
		return errors.New("Beyond the limit number ... ")
	}

	return nil
}
