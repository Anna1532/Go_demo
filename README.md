# GO学习笔记

##### 结构

bin文件存放程序编译之后的可执行文件；src存放项目源文件；pkg存放编译之后的包文件

简短函数声明不能用在全局变量

本机ip地址 192.168.2.114

##### go module

`go module`是Go1.11版本之后官方推出的版本管理工具,使用 go module 管理依赖后会在项目根目录下生成两个文件`go.mod`和`go.sum`

##### go mod 命令

需要用到第三方包的时候需要有go.mod文件，以package的方式运行

```go
go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit        编辑go.mod文件
go mod graph       打印模块依赖图
go mod init        初始化当前文件夹, 创建go.mod文件
go mod tidy        增加缺少的module，删除无用的module
go mod vendor      将依赖复制到vendor下
go mod verify      校验依赖
```

go.mod文件记录了项目所有的依赖信息，其结构大致如下：

```go
module github.com/Q1mi/studygo/blogger

go 1.12

require (
	github.com/DeanThompson/ginpprof v0.0.0-20190408063150-3be636683586
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/satori/go.uuid v1.2.0
	google.golang.org/appengine v1.6.1 // indirect间接引用：我引用的包引用了别的包，在这里标注出来
)
```

##### iota

常量计算器，只能在常量表达式中使用。const中每新增一行会加1

##### 布尔值

默认是false，%v表示显示变量的值，值可以是任何类型

##### 字符串

字符串用双引号，字符用单引号。转义字符用\，用‘’反引号定义的字符串按照原格式输出

##### type

数组，切片，字典，通道等类型与具体元素类型或长度等属性相关，称为未命名类型，可以用type为其提供具体名称，将其改编为命名类型.

取别名：

```go
// 将NewInt定义为int类型
type NewInt int
```



##### 运算符

++和--是单独的语句，不能放在等号的右边（如a=a++为非法）  

##### 三个点(...)

1.用于函数有多个不定参数的情况，可以接受多个不确定数量的参数

2.slice可以被打散进行传递

```go
func test1(args ...string) { //可以接受任意个string参数
	for _, v:= range args{
		fmt.Println(v)
	}
}

func main(){
	var strss= []string{
		"qwr",
		"234",
		"yui",
		"cvbc",
	}
	//test1(strss[0],strss[1]) 可以传多个值
    //切片被打散传入,将切片中的4个值分别传到函数中
    test1(strss...)
}
```



##### 数组

多维数组值有第一层可以使用[...]来让编译器推导数组长度（例如a:=\[...][5]可以，但是不可以使用a:=\[5][...]或者a:=\[...][...]

##### 切片

数组的局限性：1.数组的容量固定，不能自动扩展。2.值传递，数组作为函数参数时，将整个数组值拷贝一份给形参。

数组和切片的区别：创建数组时[]指定数组长度`arr:=[6]int{1,2,3,4,5,6}`；创建切片时，[]为空，或者...`s:=arr[1:5]`。

使用：切片名称[low:high:max]\(low：起始下标位置；high：结束下标位置；容量：cap=max-low)。截取数组，初始化切片时，若未指定容量，那么切片容量跟随原数组。

make函数创建切片：切片是基于数组类型做的一层封装。若需要动态的创建切片，需要使用内置的make()函数。判断一个切片是不是空的，不能用s==nil，而应该用len()==0来判断

append()函数给切片添加元素：调用append函数必须要用原来的切片变量接受返回值，原来的底层数组容不下，Go底层就会换一个底层数组。append()函数可以动态的给切片添加元素：一次添加一个元素s=append(s,1)，一次添加多个元素s=append(s,1,2,3)，也可以添加另一个切片中的元素，后面加...

##### 指针

不能进行偏移和运算

栈帧：用来给函数运行提供内存空间，取内存放到stack上。当函数调用时产生栈帧；函数调用结束时释放栈帧。

栈帧存储的内容：1.局部变量 。2.形参（形参与局部变量存储地位等同）。3.内存字段描述值

![image-20210113163129917](C:\Users\23013\Desktop\Go学习\笔记\内存分配图.png)

main函数开始运行时stack给他分配一段栈帧，main如果调用其他函数，那么给该函数也分配栈帧，函数调用结束释放内存。

空指针：未被初始化的指针 `var p *int`

野指针：被一片无效的地址空间初始化

`指针初始化`:在heap上申请一片内存地址空间 

```go
p=new(int)
fmt.println(p)
```

函数调用：值传递，实参将自己的值拷贝一份给形参，传变量传的是变量值，传指针传的是地址值。

##### 反射

包括两种方法，类型（`reflect.TypeOf()`）和种类（`reflect.ValueOf()`）

##### map

Go语言中的map(映射、字典)是一种内置的数据结构，它是一个无序的key—value对的集合，比如以身份证号作为唯一键来标识一个人的信息。

##### 函数

```go
func 函数名(参数)(返回值){
    函数体
}
func intSum(x int, y int) int {
	return x + y
}
```

##### 匿名函数

匿名函数一般用在函数内部，因为在函数内部无法定义其他函数

```
f1:=func(x,y int){
	fmt.Println(x+y)
}
//若只是使用一次的函数，还可以简写成立即执行的函数
func(x,y int){
	fmt.Println("hello")
}(100,200)//这里是参数
```

##### 闭包



##### 结构体

结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）

##### 方法

包括值接收者和指针接收者，什么时候使用指针类型接收者？

1. 需要修改接收者中的值
2. 接收者是拷贝代价比较大的大对象
3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

##### 接口

接口是一种类型，是一个需要实现的方法列表。为什么要使用接口，因为方法只能是特定类型的变量来调用

##### 结构体与json

json的作用:

1.Go语言中的结构体变量---->json格式的字符串(序列化)

2.json格式的字符串---->Go语言中的结构体变量（反序列化）

##### 包

包的路径从GOPATH/src开始写起，路径分隔符用/

想要被别的包调用的包标识符首字母需要大写

导入包的时候可以指定别名

##### 文件操作

打开文件：os.open()

关闭：文件名.Close()

##### flag

flag包支持的命令行参数类型有`bool`、`int`、`int64`、`uint`、`uint64`、`float` `float64`、`string`、`duration`

需要通过调用`flag.Parse()`来对命令行参数进行解析

##### runtime.Caller

用来获取调用runtime.Caller()方法的当前文件以及行数

```go
//pc表示函数名
//file表示调用的文件
//line表示调用的行数
//参数n表示经过几层调用，例如当前就是main函数中调用，就填0，如果是main调用了一个方法，该方法中有Caller这个方法，就填1
pc,file,line,ok:=runtime.Caller(n)
//获取函数名
funcName:=runtime.FuncForPC(pc).Name()
```



#### 并发

同步：goroutine、channel、waitgroup、context、mutex这几种，后面可以写协程池，当然也有开源的协程池，github上都有相应的

#### goroutine：

用户态的线程，由Go语言的运行时（runtime）调度完成，goroutine在开辟的时候很小，但是内存不是固定的，可以按需增大或减小，所以可以一次创建多个goroutine

#### WaitGroup

```go
var wg sync.WaitGroup
```

如果需要等待多个任务结束，可以使用`sync.WaitGroup`，通过设定计数器让每个goroutine在退出时递减，直到归0时解除阻塞

一共包括三个值：

`wg.Add(n)`:表示计数器+1

`wg.Wait`：等待

`wg.Done`：表示计数器-1

启动多个goroutine，使用`sync.WaitGroup`来实现goroutine的同步

```go
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}
```



#### channel：

用来在多个goroutine之间进行通信，协调通信

通道必须使用make初始化才能使用（make初始化：slice，map，channel）

###### 1、发送值

将一个值发送到通道中。

```go
ch := make(chan int)
ch <- 10 // 把10发送到ch中
```

###### 2、接收值

从一个通道中接收值。

```go
x := <- ch // 从ch中接收值并赋值给变量x
<-ch       // 从ch中接收值，数据就丢了
```

###### 3、关闭

我们通过调用内置的`close`函数来关闭通道。

```go
close(ch)
```

###### 4、单向通道

多用在函数参数中

#### select

需要同时从多个通道中接收数据，各个case随机选择操作，不是按顺序判断

#### Context

用来控制goroutine的退出，其余两种方式是管道和全局变量

`context`包的核心就是`Context`接口，其定义如下：

```
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

- `context.WithDeadline`会返回一个超时时间，Goroutine获得了超时时间后，例如可以对某些io操作设定超时时间。

- `Done`方法返回一个信道（channel），当`Context`被撤销或过期时，该信道是关闭的，即它是一个表示Context是否已关闭的信号。

- 当`Done`信道关闭后，`Err`方法表明`Contex`t被撤的原因。

- `Value`可以让Goroutine共享一些数据，当然获得数据是协程安全的。但使用这些数据的时候要注意同步，比如返回了一个map，而这个map的读写则要加锁。

  共有4个方法：`context.WithCancel()`,`context.WithTimeout`,`conttext.WithDeadline`,`context.WithValue`

###### 树结构

Goroutine的创建和调用关系总是像层层调用进行的，而更靠顶部的Goroutine应有办法主动关闭其下属的Goroutine的执行（不然程序可能就失控了）。为了实现这种关系，Context结构也应该像一棵树，叶子节点须总是由根节点衍生出来的。

要创建Context树，第一步就是要得到根节点，`context.Background`函数的返回值就是根节点：

```
func Background() Context
```

该函数返回空的Context，该Context一般由接收请求的第一个Goroutine创建，是与进入请求对应的Context根节点，它不能被取消、没有值、也没有过期时间。它常常作为处理Request的顶层context存在。

`context.TODO()`:未决定要传的值，可以用ＴＯＤＯ来占位

#### mutex(互斥锁)

`sync.Mutex`：底层是一个结构体，是值类型，给函数传参数的时候要传指针（同`sync.Waitgroup`）

互斥锁是并发程序中对共享资源进行访问控制的主要手段，对此Go语言提供了非常简单易用的Mutex，Mutex为一结构体类型，对外暴露两个方法Lock()和Unlock()分别用于加锁和解锁。

```go
var lock sync.Mutex
lock.Lock()//加锁
lock.Unlock()//解锁
```

防止多个goroutine同一时刻操作同一资源

##### 读写互斥锁

适用于读多写少的场景下，才能提高程序的执行效率

特点：

​		1.（读是并发的，写是串行的）读的goroutine来了获取的是读锁，后续的goroutine能读不能写

​		2.写的goroutine来了，获取的是写锁，后续的goroutine不管是读还是写都要等待获取锁

```go
var rwLock sync.RWMutex
rwLock.RLock()//获取读锁
rwLock.RUnLock()//释放读锁

rwLock.Lock()//获取写锁
rwLock.Unlock()//释放写锁
```



##### 常见权限表示形式

-rw------- (600)      只有拥有者有读写权限。
-rw-r--r-- (644)      只有拥有者有读写权限；而属组用户和其他用户只有读权限。
-rwx------ (700)     只有拥有者有读、写、执行权限。
-rwxr-xr-x (755)    拥有者有读、写、执行权限；而属组用户和其他用户只有读、执行权限。
-rwx--x--x (711)    拥有者有读、写、执行权限；而属组用户和其他用户只有执行权限。
-rw-rw-rw- (666)   所有用户都有文件读、写权限。
-rwxrwxrwx (777)  所有用户都有读、写、执行权限。

#### 多态

同一操作作用于不同的对象，可以产生不同的效果。就是多态，继承可以扩展已存在的代码模块，而目的就是为了代码重用。

##### 实现一个日志系统

1、实现往不同的地方输出日志

终端输出/文件输出

2、日志分级别（info，error，warning等）

3、支持开关控制（如上线之前什么级别都可以输出，但是上线止之后只有info级别以下才能输出）

4、完整的日志记录要包含时间、行号、文件名、日志级别、日志信息

5、日志文件要切割

按文件大小切割：每次记录日志之前都判断一下当前写的这个文件大小，如果太大就切割一下

按日期切割



#### sirupsen/logrus

###### 安装

```gobash
go get github.com/sirupsen/logrus
```

###### 特点

- 完全兼容标准日志库，拥有七种日志级别：`Trace`, `Debug`, `Info`, `Warning`, `Error`, `Fatal`and `Panic`。
- 可扩展的Hook机制，允许使用者通过Hook的方式将日志分发到任意地方，如本地文件系统，logstash，elasticsearch或者mq等，或者通过Hook定义日志内容和格式等
- 可选的日志输出格式，内置了两种日志格式JSONFormater和TextFormatter，还可以自定义日志格式

###### 字段

除了使用WithField或WithFields添加的字段外，某些字段还会自动添加到所有日志记录事件中：

`time`:创建条目的时间戳。

`msg`:在AddFields调用之后，日志消息传递给{Info,Warn,Error,Fatal,Panic}。 例如,Fail to send event。

`level`:日志记录级别,例如Info

###### 格式

内置的日志格式是：

`logrus.TextFormatter`生成文本文件

`logrous.JSONFormatter`生成JSON格式的文件

#### spf13/viper

用来配置管理Go应用程序，处理所有类型的配置需求和格式，可以将Viper视为满足您所有应用程序配置需求的注册表。支持： 

- 设置默认值
- 从JSON，TOML，YAML，HCL，envfile和Java属性配置文件中读取
- 实时观察和重新读取配置文件（可选）
- 从环境变量中读取
- 从远程配置系统（etcd或Consul）中读取，并观察更改
- 从命令行标志读取
- 从缓冲区读取
- 设置显式值

###### 优先级

- explicit call to `Set`
- flag
- env
- config
- key/value store
- default

Viper所需的配置最少，因此知道在哪里可以找到配置文件。 Viper支持JSON，TOML，YAML，HCL，INI，envfile和Java属性文件。 Viper可以搜索多个路径，但是当前单个Viper实例仅支持单个配置文件。 Viper不会默认使用任何配置搜索路径，而会将默认决定留给应用程序。

###### 寻找并读取配置文件

```go
viper.SetConfigName("config") // 设置配置文件的文件名
viper.SetConfigType("yaml") // 如果配置文件名没有后缀，在这里设置后缀
viper.AddConfigPath("/etc/appname/")   // 在该位置查找配置文件的路径
viper.AddConfigPath("$HOME/.appname")  // 可以多次调用来添加多条搜索路径
viper.AddConfigPath(".")               // （可选）在工作目录中查找配置
err := viper.ReadInConfig() //寻找并且读取配置文件
if err != nil { 
	panic(fmt.Errorf("Fatal error config file: %s \n", err))
}
```

错误处理

```go
if err := viper.ReadInConfig(); err != nil {
    if _, ok := err.(viper.ConfigFileNotFoundError); ok {
        // Config file not found; ignore error if desired
    } else {
        // Config file was found but another error was produced
    }
}

// Config file found and successfully parsed
```

###### 写配置文件

```go
 // 将当前的Viper配置写入预定义的路径（如果存在),如果没有预定义的路径，则错误。 如果存在，将覆盖当前的配置文件。将当前配置写入由“viper.AddConfigPath（）”和“ viper.SetConfigName”设置的预定义路径
viper.WriteConfig()
//将当前的viper配置写入预定义的路径。如果没有预定义的路径，则出错。不会覆盖当前配置文件(如果存在)
viper.SafeWriteConfig()
//将当前的viper配置写入给定的文件路径。将覆盖给定文件(如果存在)
viper.WriteConfigAs("/path/to/my/.config")
 // 将当前的viper配置写入给定的文件路径。不会覆盖给定文件(如果存在)。如果已经写过再执行会报错
viper.SafeWriteConfigAs("/path/to/my/.config")
viper.SafeWriteConfigAs("/path/to/my/.other_config")
```

viper支持在程序运行时实时读取配置文件，但是需要确保在调用WatchConfig()之前添加了所有配置路径

```go
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
	fmt.Println("Config file changed:", e.Name)
})
```

yaml是一种用来写配置文件的序列化语言

配置文件格式一般为TOML:https://github.com/LongTengDao/TOML/

###### 代码顺序

1、设置配置文件搜索路径

2、设置配置文件名称

3、设置默认值

4、读取文件

5、打印一部分读取到的值

6、重新写入文件

```go
func main(){

	viper.AddConfigPath("./") // 设置读取路径：就是在此路径下搜索配置文件。
	viper.SetConfigFile("server.yaml") // 设置被读取文件的全名，包括扩展名。
	viper.SetDefault("name", "dogger")//设置默认值
	viper.SetDefault("age", "18")
	viper.SetDefault("class", map[string]string{"class01": "01", "class02": "02"})
	viper.ReadInConfig()  // 读取配置文件： 这一步将配置文件变成了 Go语言的配置文件对象包含了 map，string 等对象。
	fmt.Println(
		viper.Get("name"), // 获取配置文件的信息也很容易，用 Get方法。
		viper.Get("age"),
		viper.Get("name.first"),
		viper.Get("hobbies"),
	)
	// 控制台输出： map[first:panda last:8z] 99 panda [Coding Movie Swimming]
	viper.WriteConfigAs("server-01.yaml")
}
```



#### urfave/cli

###### 介绍

cli是一个用于构建命令行程序的库，使用时创建一个`cli.App`结构的对象，然后调用其`Run()`方法，传入命令行的参数即可。一个空白的`cli`应用程序如下：

```go
func main() {
  (&cli.App{}).Run(os.Args)
}
```

通过`cli.Context`的相关方法我们可以获取传给命令行的参数信息：

- `NArg()`：返回参数个数；
- `Args()`：返回`cli.Args`对象，调用其`Get(i)`获取位置`i`上的参数。

`cli`设置和获取选项非常简单。在`cli.App{}`结构初始化时，设置字段`Flags`即可添加选项。

###### 子命令

`cli`通过设置`cli.App`的`Commands`字段添加命令，设置各个命令的`SubCommands`字段，即可添加子命令



#### jinzhu/gorm

连接数据库：mysql -u root -p





#### jsoniter





#### net/http

提供了HTTP客户端和服务端的实现

TCP粘包问题：由于neigo算法在发送数据包时，会等待有没有其他数据包一起发送，所以收到数据包时会出现好几条数据粘在一起的情况。

##### websocket

http是请求与响应，websocket是自动推送，不需要请求

#### sync.Once

sync.Once能确保实例化对象Do方法在多线程环境只运行一次,内部通过互斥锁实现

##### Sync.Map

使用场景：并发操作一个map的时候，内置的map不是并发安全的，sync.Map是一个开箱即用的并发安全的map

```go
var syncMap sync.Map
syncMap.Store(key,value)
syncMap.Load(key)
syncMap.LoadOrStore()
syncMap.Delete()
syncMap.Range()
```



#### grpc（程序之间通信）



#### elastic（做搜索的）



#### Kafka（传输以及存储数据）



#### GitLab相关

该项目已经提交到gitlab上了，但我又对项目中的某个文件进行了修改，所以要再次提交到gitab上保存

1.定位到曾经提交到gitlab上的文件夹，右键，找到Git Bash Here，进入命令行窗口

2.输入：git status，敲回车查看代码是否有更新，有更新的话会出现文件改变的文件名。（红色的）

3.输入：git add .

4.输入：git commit -a -m "tag"，敲回车，然后你就会看到有几个文件修改了。（如果写了-m，那么“”中一定要写内容，不然会报错，tag表示的内容为本次提交的备注信息，随便写）

5.输入：git push， 敲回车，这样所有的更新代码都上传到git上了。
