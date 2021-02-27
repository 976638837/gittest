package main

import "fmt"

type Author struct {
	Name string
	VIP bool
	Icon string
	Signature string
	Focus int
}
type BilBli struct {
	rname string
	sname string
	dianzanshu int
	shoucangliang int
	toubiliang int
}
type dianzan interface {
	dianzan()
}
type shoucang interface {
	shoucang()
}
type toubi interface {
	toubi()
}
type sanlian interface {
	sanlian()
}
func (p *BilBli) dianzan()  {
	p.dianzanshu++
}
func (p *BilBli) shoucang()  {
	p.shoucangliang++
}
func (p *BilBli) toubi(n int)  {
	p.toubiliang=p.toubiliang+n
}
func (p *BilBli) sanlian()  {
	p.shoucangliang++
	p.toubiliang++
	p.dianzanshu++
}

func (p *BilBli)FaBu(sname string,rname string)  {

		p.rname= rname
		p.sname= sname
		p.dianzanshu= 0
		p.shoucangliang=0
		p.toubiliang= 0


}
func main()  {
	p:=&BilBli{}
	sname:="gaoxiao"
	rname:="lujing"
	p.FaBu(sname,rname)
	fmt.Println(p)
	p.dianzan()
	p.shoucang()
	p.toubi(50)
	fmt.Println(p)
}
