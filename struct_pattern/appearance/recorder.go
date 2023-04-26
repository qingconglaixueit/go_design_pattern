// @Author Bing 
// @Desc
package main

import "fmt"

type Recorder struct {
	money int
}

func NewRecorder() *Recorder {
	return &Recorder{money: 0}
}

func (r *Recorder) GetData() int {
	return r.money
}
func (r *Recorder) SaveData(money int) {
	r.money = money
	fmt.Println("has been save data ... ")
}
