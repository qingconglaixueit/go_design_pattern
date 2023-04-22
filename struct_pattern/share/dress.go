// @Author Bing 
// @Desc
package main

import (
	"errors"
	"fmt"
)

const (
	APlayer = "APlayer"
	BPlayer = "BPlayer"
)

type Dress interface {
	getColor() string
}

type DressFactory struct {
	Dm map[string]Dress
}

var DressInstance = &DressFactory{Dm: map[string]Dress{}}

func (d *DressFactory) getDressType(pType string) (Dress, error) {
	if d.Dm[pType] != nil {
		return d.Dm[pType], nil
	}

	// d.Dm[pType] 为 nil 的情况

	if pType == APlayer {
		d.Dm[pType] = NewAPlayer()
		fmt.Println("NewAPlayer dress ... ")
		return d.Dm[pType], nil
	}

	if pType == BPlayer {
		d.Dm[pType] = NewBPlayer()
		fmt.Println("NewBPlayer dress ... ")
		return d.Dm[pType], nil
	}

	return nil, errors.New("not support pType ... ")
}

func getDressInstance() *DressFactory {
	return DressInstance
}
