// @Author Bing
// @Desc
package main

import "errors"

// 具体的车厂

func getCar(band string) (ICar, error) {
	if band == "Bmw" {
		return newBmw(), nil
	}
	if band == "Benz" {
		return newBenz(), nil
	}
	return nil, errors.New("not support car band ... ")
}
