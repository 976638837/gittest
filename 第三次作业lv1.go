package main

import (
	"fmt"
	"math"
	"time"
)

var prim[50000] bool
func main()  {

	intChan:=make(chan int,1000)
	primeChan:=make(chan int,7000)
	exitChan:=make(chan bool,8)

	start := time.Now().UnixNano()
	go putNum(intChan)
	Sushu()
	for i:=0;i<8 ;i++  {
		go primeNum(intChan,primeChan,exitChan,&prim)
	}
	go func() {
		for i:=0;i<8 ;i++  {
			<-exitChan
		}

		close(primeChan)
	}()

	for{
		res,ok:=<-primeChan
		if !ok {
			break
		}
		fmt.Println("素数为：",res)
	}
	end := time.Now().UnixNano()
	fmt.Println("使用协程耗时=", end - start)
}

func putNum(intChan chan int)  {
	for i:=1;i<50000;i++{
		intChan<-i
	}
	close(intChan)
}

func primeNum(intChan chan int,primeChan chan int,exitChan chan bool,prim *[50000] bool)  {
	for{
		num,ok:=<-intChan
		if ! ok {
			break
		}
		if prim[num] {
			primeChan<-num
		}
	}
	exitChan<-true
}

func Sushu()  {
	var (
		i int
		j int
	)
	for i:=2;i<50000 ; i++ {
		if i%2!=0 {
			prim[i]=true
		}
	}
	for i=3;i<int(math.Sqrt(50000));i++  {
		if prim[i] {
			for j=i+i;j<50000 ;j=j+i  {
				prim[j]=false
			}
		}
	}
}