// @Author Bing 
// @Desc
package main

// 医院里面有很多不同职业的人来看病，如何保证用户进入医生办公室发生冲突呢？这个时候就有一个医院服务人员来进行管理，负责和不同的职业的病人沟通，
// 虽然病人之间没有沟通，但是他们依然能有有序的进入到医生办公室内看病，此时医院服务人员就是一个中介者
func main() {

	// 创建一个中介者对象
	hs := &HosService{Queue: make([]IPeople, 0), IsFree: true}

	t1 := &Teacher{Mgr: hs}

	t2 := &Teacher{Mgr: hs}

	w1 := &Worker{Mgr: hs}

	w2 := &Worker{Mgr: hs}

	t1.Enter()
	t2.Enter()
	w1.Enter()
	w2.Enter()

	t2.Leave()
	w1.Leave()
	w2.Leave()

}
