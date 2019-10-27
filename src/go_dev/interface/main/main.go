package main

import "log"

/*
1、创建接口
	type 接口名称 interface {
		Writer(b []byte) (n int,err error)
	}
	type Stringer interface {
		String() string
	}
2、实现接口方式、实现接口的条件
	1、结构体实现接口 方法名称和参数类型必须一样
	2、返回变量是接口，接口中所有方法必须全部实现，如果实例返回的是类，不需要全部实现
	type FileIo struct {
	}
	func (f FileIo) Write(b []byte) (n int,err error){
	}
3、一个类型实现多个接口
	type Writer interface {
		Write(data interface{})
	}
	type Closer interface {
		Close()
	}

	type File struct {
	}

	func (f File) Close(){
	}
	// 这里的interface是任意类型，循环遍历的使用时需要做类型转换
	func (f File) Write(data interface{}) {

	}
4、多个类型实现相同接口
	type Service interface {
		Start()
		Log(string)
	}
	type Logger struct {

	}
	func (l *Logger) Start(){

	}

	type GameService struct {
		Logger
	}

	func (g GameService) Log(strs string) {

	}
	var g =new(GameService)
	g.Start()
	g.Log("xxx")
5、组合接口
	type File struct {
	}
	func (f File) Close() {
		log.Println("close")
	}
	func (f File) Write(data interface{}) {
		log.Println(data)
	}
	type Writer interface {
		Write(data interface{})
	}
	type Closer interface {
		Close()
	}
	// 聚合接口
	type WriterCloser interface {
		Write(data interface{})
		Close()
	}
	// 组合接口
	var wc WriterCloser = new(File)
	wc.Write("component interface")
	wc.Close()
6、接口和类型之间的转换
	类型断言的方式
	t := i.(T) 这种方式如果会触发宕机
	t,ok := i.(T) 不会出现触发宕机，如果转换T类型没有实现该接口ok为false，t为0
	var a inteface = new(bird)
	a,ok = a.(Flyer)
*/

type bird struct {
}
type Flyer interface {
	Fly()
}
type Walker interface {
	Walk()
}

// 实现接口
func (b *bird) Fly() {
	log.Println("fly")
}
func (b *bird) Walk() {
	log.Println("Walk")
}

type pig struct {
}

func main() {
	// 声明接受的任意类型，相当于java的Objcet
	var obj interface{} = new(bird)
	b, ok := obj.(Flyer)
	if ok {
		b.Fly()
	}
	// 为实现Flyer 接口
	var obj1 interface{} = new(pig)
	// f := obj1.(Flyer) 报错 *main.pig is not main.Flyer: missing method Fly
	f, ok := obj1.(Flyer)
	if ok {
		f.Fly()
	} else {
		log.Println(f, ok)
	}
	// 接口和结构体互相转换
	var fx Flyer = new(bird)
	f3, ok := fx.(*bird)
	if ok {
		f3.Walk()
	}

	f1 := new(FileIo)
	var w Writer = f1
	w.Write(nil)
	w.WriteData([]int{1, 2, 3, 4, 5})

	var f2 = new(File)
	var w1 Writer1 = f2
	var c1 Closer = f2
	w1.Write("data")
	c1.Close()

	// 组合接口
	var wc WriterCloser = new(File)
	wc.Write("component interface")
	wc.Close()

	// 此处如果变量名为接口，需要实现所有接口，如果不是，可以选择实现
	// var g Service = new(GameService)
	var g = new(GameService)
	g.Start()
	g.Log("xxx")
	var l = new(Logger)
	l.Start()

	log := createLogger()
	log.Log("w22222")
}

func createLogger() *logger {
	l := NewLogger()
	cw := newConsoleWriter()
	l.RegisterWriter(cw)
	fw := newFileWriter()
	if err := fw.SetFile("log.log"); err != nil {
		log.Println(err)
	}
	l.RegisterWriter(fw)
	return l
}

type Service interface {
	Start()
	Log(str string) bool
	WriteData(data interface{}) error
	// 能否写入
	CanWrite() bool
}

type Logger struct {
}

func (l *Logger) Start() {
	log.Println("===start===")
}

type GameService struct {
	Logger
}

func (g *GameService) Log(strs string) {
	log.Println(strs)
}

type Writer1 interface {
	Write(data interface{})
}
type Closer interface {
	Close()
}

// 聚合接口
type WriterCloser interface {
	Write(data interface{})
	Close()
}

type File struct {
}

func (f File) Close() {
	log.Println("close")
}

func (f File) Write(data interface{}) {
	log.Println(data)
}

type FileIo struct {
}

func (f FileIo) Write(b []byte) (a int, err error) {
	return 0, nil
}

func (f FileIo) WriteData(data interface{}) {
	b := data.([]int)
	for _, v := range b {
		log.Printf("%v", v)
	}
}

type Writer interface {
	Write(b []byte) (n int, err error)
	WriteData(data interface{})
}
