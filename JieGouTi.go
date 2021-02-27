package main

import "fmt"

type Person struct {
	name string
	age int
	gender string
}
type dove interface {
	gugu()
}
type lingmengjing interface {
	lingmeng()
}
type zhenxiangguai interface {
	zhenxiang()
}
type repeater interface {
	repeat(string)
}

func (p *Person) repeat(word string)  {
	fmt.Println(word)
}
func (p *Person) gugu()  {
	fmt.Println(p.name,"又放鸽子了！！")
}
func (p *Person) lingmeng() {
	fmt.Println(p.name,"又开始酸了！！！")
}
func (p *Person) zhenxiang() {
	fmt.Println(p.name,"又开始真相了！！")
}
func Receiver(v interface{})  {
	switch v.(type) {
	case string:
		fmt.Println("这是个string" )
	case int:
		fmt.Println("这是个int")
	case float32:
		fmt.Println("这是个float32")
	case bool:
		fmt.Println("这是个bool")
	}
}
func main()  {
	p:=&Person{}
	p.name="lujing"
	p.age=21
	p.gender="nan"
	p.gugu()
	p.repeat("sb lujing")
	p.lingmeng()
	p.zhenxiang()
	Receiver("卢靖sb！")
	Receiver(21)
	Receiver(21.5)
	Receiver(true)
}
