// @Author Bing 
// @Desc
package main

import "errors"

// BigShop
type BigShop interface {
	MakeShoe() IShoe
	MakeCloth() ICloth
}
// @Title GetShopFac
// @Description  获取商品的工厂
// @Param band
// @Return BigShop
// @Return error
func GetShopFac(band string) (BigShop, error) {
	if band == "NB" {
		return NB{}, nil
	}
	if band == "Nike" {
		return Nike{}, nil
	}
	return nil, errors.New("band is not suppport ...")
}
