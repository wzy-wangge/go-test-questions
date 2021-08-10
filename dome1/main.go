package main

import (
	"fmt"
	"sync"
)
//题目介绍
//有3个函数可以分别输出aa,bb,cc,请开3个goroutine调用函数进行有序输出,aabbcc为一行

func main()  {
	var wg sync.WaitGroup
	wg.Add(3)
	aChan := make(chan struct{},1)
	bChan := make(chan struct{},1)
	cChan := make(chan struct{},1)
	go aa(&wg,aChan,bChan)
	go bb(&wg,bChan,cChan)
	go cc(&wg,aChan,cChan)

	aChan <- struct{}{}
	wg.Wait()
}



func aa(wg *sync.WaitGroup,aChan,bChan chan struct{}){
	for  i:=0;i<100;i++{
		<- aChan
		fmt.Printf("%s","aaa:")
		bChan <- struct{}{}
	}
	wg.Done()
}


func bb(wg *sync.WaitGroup,bChan,cChan chan struct{}){
	for  i:=0;i<100;i++{
		<- bChan
		fmt.Printf("%s","bbb;")
		cChan <- struct{}{}
	}
	wg.Done()
}


func cc(wg *sync.WaitGroup,aChan,cChan chan struct{}){
	for  i:=0;i<100;i++{
		<- cChan
		fmt.Printf("%s\n","ccc:")
		aChan <- struct{}{}
	}
	wg.Done()
}