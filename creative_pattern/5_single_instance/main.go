// @Author Bing 
// @Date 2023/3/27 17:37:00 
// @Desc
package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

var one sync.Once

type Single struct {
}

// 此时创建的  SingleInstance 默认为 零值， 也就是 nil
var SingleInstance *Single

// getOneInstance 显示使用 锁的方式，保证线程安全
func getOneInstance() *Single {
	if SingleInstance == nil {
		// 此处可能会有多个协程进来
		lock.Lock()
		defer lock.Unlock()
		// 此处再继续 判断 SingleInstance 是否为 nil，是因为可能其他协程已经拿到锁，且初始化好了 SingleInstance 后，当前协程才拿到锁
		if SingleInstance == nil {
			fmt.Println("getOneInstance: Create SingleInstance successfully ...")
			SingleInstance = &Single{}
		} else {
			fmt.Println("1 getOneInstance: SingleInstance has been created ...")
		}
	}else {
		fmt.Println("2 getOneInstance: SingleInstance has been created ...")
	}
	return SingleInstance
}

// getTwoInstance 使用 sync.Once 达到同样的效果
func getTwoInstance() *Single {
	if SingleInstance == nil {
		fmt.Println("getTwoInstance: in SingleInstance == nil ...")
		one.Do(func() {
			// 因此 one.Do 只会执行一次，因此不会出现  getOneInstance 的情况，因此此处无需判断 SingleInstance 是否为 nil
			fmt.Println("getTwoInstance: Create SingleInstance successfully ...")
			SingleInstance = &Single{}
		})
	} else {
		fmt.Println("getTwoInstance: SingleInstance has been created ...")
	}
	return SingleInstance
}

func main() {

	for i := 0; i < 20; i++ {
		getTwoInstance()
	}

	fmt.Scanln()// 输出回车就会结束

}
