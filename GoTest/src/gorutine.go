package main

import (
	"fmt"
	"time"
	"sync"
)

func check(id int){	
	fmt.Println("checked", id)
}

func testGOrutine(){
	<-time.After(time.Duration(50)*time.Millisecond)
	fmt.Println("haha ~")
}


func check2(id int, wg *sync.WaitGroup) {
    fmt.Println("Checked", id)
    <-time.After(time.Duration(id)*time.Millisecond)
//	time.After(time.Duration(id)*time.Millisecond)
    fmt.Println("Woke up", id)
    wg.Done()
}

//调用者等待所有goruntine结束
//++++++++++++++++++++++++++++++++++++++++++++++++++++++
////让主进程停住，不然主进程退了，goroutine也就退了
//++++++++++++++++++++++++++++++++++++++++++++++++++++++
func waitingUtillAllOver (){
	var wg sync.WaitGroup
    for i := 0; i <= 10; i++ {
        wg.Add(1)
        fmt.Println("Called with", i)
        go check2(i, &wg)
    }
    wg.Wait()
    fmt.Println("Done for")
}


func main(){
	fmt.Println("haha 22~")
	go testGOrutine()
	
//	waitingUtillAllOver()
}

func testGorutine() {
	for i:= 0; i <= 10 ; i ++{
		fmt.Println("Called with", i)
		go check(i)
	}
	fmt.Println("Done for")
}


