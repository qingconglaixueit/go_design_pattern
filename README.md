# go_design_pattern
go_design_pattern

日常工作中免不了使用设计模式，那么你使用了哪些设计模式呢？

## 设计模式是什么？

设计模式是一种在软件设计中对常见问题的通用解决方案。

**它们是经过验证的、可重用的设计思想，可以帮助解决开发过程中遇到的各种问题。**

设计模式提供了一种共同的词汇表和方法论，让不同团队的开发人员能够更有效地沟通和协作，从而提高软件的稳定性、可靠性和可维护性。




整体来看，设计模式包含了如下 **22 种，主要分为三大类**

1.  创造型
1.  结构型
1.  行为型

![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/17ee132f37ed494f9f97648570076b88~tplv-k3u1fbpfcp-zoom-1.image)

看上去很多种，实际上咱们实际仔细去看每种模式的思想和细节的时候，能够发现其实我们日常工作中有用到，只是不知道原来是应用了设计模式




**有道无术术可求，有术无道止于术** ，那么先来学习学习每种设计模式的思想，以及 GO 语言实现的 demo




## 创建型模式

1.  工厂模式

工厂模式 **解决了在不指定具体类的情况下创建产品对象的问题**

案例：

写一个工厂，可以生产车，目前这个工厂可以生产 bmw 和 benz ，那么我们需要获取一辆车的时候，就直接去车厂拿车即可

-   [demo 1](https://github.com/qingconglaixueit/go_design_pattern/tree/main/creative_pattern/1_factory)




2.  抽象工厂模式

**它能创建一系列相关的对象， 而无需指定其具体类**

案例：

一个大卖场，有衣服，有 鞋子，有 NB 品牌，有 Nike 品牌，此时我们就可以使用抽象工厂，抽象衣服，鞋子，每一个品牌自行去实现具体的衣服和鞋子

-   [demo 2](https://github.com/qingconglaixueit/go_design_pattern/tree/main/creative_pattern/2_abstract_factory)




3.  生成器模式

**能够分步骤创建复杂对象**

案例：

建一座建筑，无论你是冰室，还是住房，是非常复杂的，我们演示简化为 4 个步骤，每个步骤都会创建新的对象，大体会经历如下步骤

-   搭建框架
-   装修
-   安装门
-   安装床




-   [demo 3](https://github.com/qingconglaixueit/go_design_pattern/tree/main/creative_pattern/3_generator)




4.  原型模式

**能够复制已有对象， 而又无需使代码依赖它们所属的类**

案例：

此时可以想到我们的程序目录，和目录下的文件，是否可以通过这种方式实现呢？

程序的目录和文件，能够满足递归的结构，目录和文件，都是同样的结构




-   [demo 4](https://github.com/qingconglaixueit/go_design_pattern/tree/main/creative_pattern/4_prototype)







5.  单例模式

**能够保证一个类只有一个实例， 并提供一个访问该实例的全局节点**

单例实例会在结构体首次初始化时创建，若多个协程创建，如何保证协程安全呢？

```
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
      //getOneInstance()
      getTwoInstance()
   }

   fmt.Scanln()// 输出回车就会结束

}
```

```
5_single_instance>go run main.go
getTwoInstance: in SingleInstance == nil ...
getTwoInstance: Create SingleInstance successfully ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
getTwoInstance: SingleInstance has been created ...
```




-   [demo 5](https://github.com/qingconglaixueit/go_design_pattern/tree/main/creative_pattern/5_single_instance)

## 结构型模式

6.  适配器模式

**让接口不兼容的对象能够相互合作**

案例：

可能我们都用过生活中的适配器吧，我们手机是 Type C 的接口，电脑是 USB 的接口，那么这个时候，如何让 TypeC 的接口和 USB 的接口连接上呢？这个时候就可以使用适配器来处理




-   [demo 6](https://github.com/qingconglaixueit/go_design_pattern/tree/main/struct_pattern/6_adapter)




7.  桥接模式

**可将一个大类或一系列紧密相关的类拆分为抽象和实现两个独立的层次结构， 从而能在开发时分别使用** 案例：

不同的显示器要接入计算机，那么计算机可以独立开发，显示器也可以独立开发，无论计算机里面有 windows 的，还是 mac 的 仍然支持设置 显示器




-   `demo 7`




8.  组合模式

使用它将对象组合成树状结构， 并且能像使用独立对象一样使用它们

案例：

咱们的目录树中有 目录，有 文件， 目录可以进行查询，文件也是可以进行查询，将他们组合起来，他们各自仍然是独立的对象




-   [demo 8](https://github.com/qingconglaixueit/go_design_pattern/tree/main/struct_pattern/8_component)




9.  代理模式

能够提供对象的替代品或其占位符

案例：

对于 nginx 反向代理很熟悉吧，咱们手撸一个简单的反向代理，运用代理模式




-   [demo 9](https://github.com/qingconglaixueit/go_design_pattern/tree/main/struct_pattern/9_proxy)




10. 装饰器模式

允许你通过将对象放入包含行为的特殊封装对象中来为原对象绑定新的行为

案例：

年轻人喜欢喝奶茶，不同的口味不同的配置有不同的价格，喝奶茶，有原味的，有加珍珠的，有加椰果的，有加奶的 ，各有各的价格

那么 奶茶 和 珍珠，椰果，奶 是相互独立的，各自是一个独立的对象 我们可以在珍珠对象中，放入奶茶对象，进而达到 奶茶中加珍珠，加椰果，加奶 达到不同的价格




-   [demo 10](https://github.com/qingconglaixueit/go_design_pattern/tree/main/struct_pattern/10_decorator)




11. 外观模式

**能为程序库、 框架或其他复杂类提供一个简单的接口**

案例：

咱们去银行存钱，取钱，表面上看着很简单，一进一出即可，殊不知银行后面的内部系统之间复杂的交互，我们简化一下 例如 银行后面涉及到 账户系统，检测系统，数据库系统，通知系统 这个时候 就可以使用 外观模式，对于客户端来说使用非常简单，无需关注其内部复杂系统




-   [demo 11](https://github.com/qingconglaixueit/go_design_pattern/tree/main/struct_pattern/11_decorator)




12. 享元模式

**通过共享多个对象所共有的相同状态**

案例：

设计一个足球比赛游戏，A 队队服是 红色，B 队队服是 蓝色，且整个队伍就只有一个套衣服 ，此处就应用了享元模式

如果不应用享元模式，那么我们每一个人都要 new 一套衣服，咱确实没有必要这样做




-   [demo 12](https://github.com/qingconglaixueit/go_design_pattern/tree/main/struct_pattern/21_share)




## 行为模式

13. 命令模式

**将请求转换为一个包含与请求相关的所有信息的独立对象**

案例：

咱们以插件的方式来注入到程序中，如 我们平日看电视，需要按按钮，按钮上有各种按键，例如开机键，关机键

另外，这个按键又是对于哪些设备生效呢？此处咱们是对 TV 生效，可以使用命令模式，很好的插入这些相关对象




-   [demo 13](https://github.com/qingconglaixueit/go_design_pattern/tree/main/behavior/13_command)




14. 策略模式

**让你定义一系列算法， 并将每种算法分别放入独立的类中， 以使算法的对象能够相互替换**

案例：

咱们知道缓存的处理方式，有 FIFO，LFU，LRU， 那么当我们将数据载入缓存的时候，使用的算法进行切换了，我们仍然需要这个缓存是可以正常处理的，这个时候，咱们就可以使用策略模式




又如我们需要取到一个目的地，我们可以坐飞机，可以乘高铁，也可以走路去，在过程中仍然可以随意切换交通方式，最终，我们都是抵达目的地，没有问题，同样也可以使用策略模式




-   [demo 14](https://github.com/qingconglaixueit/go_design_pattern/tree/main/behavior/14_strategy)




15. 模板模式

**在超类中定义一个算法的框架， 允许子类在不修改结构的情况下重写算法的特定步骤**

案例：

模板的方式，相信我们平日用的也非常多，举个例子我们就能够很好的感受到

我们要获取消息，保存消息到缓存，生成消息回复，发送消息，这几个步骤不变的情况下，我们可以使用 微信，可以使用企业微信，自然也是可以使用飞书来处理




**我们生成 源 msg，期望放到缓存中，生成具体的回复信息，将消息发送出去** 对于这种步骤固定，但是对于每一步骤期望有自己的实现方式的，期望去重写具体的实现方式的，可以使用模板方法

获取到信息之后，wechat ，wework，lark 有不同的处理方式，但是他们的整体处理结构是一样的




-   [demo 15](https://github.com/qingconglaixueit/go_design_pattern/tree/main/behavior/15_template_method)




16. 迭代器模式

让你能在不暴露集合底层表现形式 （列表、 栈和树等） 的情况下遍历集合中所有的元素

案例：

我们提供出去的接口，不期望别人能看到我们自身的数据结构，我们就可以使用迭代器的方式来进行处理，可以隐藏实际的数据结构




-   [demo 16](https://github.com/qingconglaixueit/go_design_pattern/tree/main/behavior/16_iterator)







17. 观察者模式

允许你定义一种订阅机制， 可在对象事件发生时通知多个 “观察” 该对象的其他对象

案例：

这种设计模式应用很多，我们举例一个食堂吃饭的例子

多个员工观察食堂，食堂做好菜之后 订阅的员工都能够收到通知，员工便倾巢而出




-   [demo 17](https://github.com/qingconglaixueit/go_design_pattern/tree/main/behavior/17_observe)




18. 责任链模式

允许你将请求沿着处理者链进行发送。 收到请求后， 每个处理者均可对请求进行处理， 或将其传递给链上的下个处理者

案例：

例如处理一件事情，有一个固定的流程，正如一个新员工入职

员工入职可以看做是一个责任链， **前台报道，人事制度宣讲，合同部合同签订，送入具体部门** , 这种部门职责明确，链路清晰，可以使用 责任链模式，允许你将请求沿着处理者链进行发送 对于新员工来说，只需要知道去去找前台报道即可




-   [demo 18](https://github.com/qingconglaixueit/go_design_pattern/tree/main/behavior/18_responsebility)




19. 备忘录模式

**允许在不暴露对象实现细节的情况下保存和恢复对象之前的状态。**

案例： 例如某白领使用备忘录，记录了 **1 ，2 ，3 条**信息，然后想调出 **2 快照**来进行查看




-   [demo 19](https://github.com/qingconglaixueit/go_design_pattern/tree/main/behavior/19_memo)




20. 中介者模式

**能让你减少对象之间混乱无序的依赖关系。 该模式会限制对象之间的直接交互， 迫使它们通过一个中介者对象进行合作。**

案例：

医院里面有很多不同职业的人来看病，如何保证用户进入医生办公室发生冲突呢？

这个时候就有一个医院服务人员来进行管理，负责和不同的职业的病人沟通， **虽然病人之间没有沟通，但是他们依然能有有序的进入到医生办公室内看病，此时医院服务人员就是一个中介者**




-   [demo 20](https://github.com/qingconglaixueit/go_design_pattern/tree/main/behavior/20_mediator)




21. 状态模式

让你能在一个对象的内部状态变化时改变其行为， 使其看上去就像改变了自身所属的类一样

案例：

售卖机暂时只售卖一种从产品，有 4 种状态，没货，有货，正在请求商品，已投币

对于不同的状态，应对不同的动作




-   [demo 21](https://github.com/qingconglaixueit/go_design_pattern/tree/main/behavior/21_status)




22. 访问者模式

**将算法与其所作用的对象隔离开来**

案例：

目前有 3 种结构， Dog，Cat，Monkey 都实现了 IAnimal interface，访问者有 AreaVs（计算面积），VolumeVs（计算体积），ColorVs （计算颜色） 都实现了 IVisitor interface

以后如果需要类似于增加计算体脂率的，那么就可以加一个 visitor 即可，且对其他的访问者和结构没有影响

如果需要增加结构，例如增加 pig ，那么直接去实现 Animal 即可，**对其他的结构和访问者没有影响**

有没有觉得很巧妙呢？




-   [demo 22](https://github.com/qingconglaixueit/go_design_pattern/tree/main/behavior/22_visitor)




## 欢迎点赞，关注，收藏

朋友们，你的支持和鼓励，是我坚持分享，提高质量的动力

![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/2e5a348410684831bf3cfd735a74fb0b~tplv-k3u1fbpfcp-zoom-1.image)

好了，本次就到这里

技术是开放的，我们的心态，更应是开放的。拥抱变化，向阳而生，努力向前行。

我是阿兵云原生，欢迎点赞关注收藏，下次见~