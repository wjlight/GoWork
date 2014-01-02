package main

import (
	"fmt"
	"runtime"
	"time"
	"math/rand"
	"sync"
)


/**
如果一个goroutine没有被阻塞，那么别的goroutine就不会得到执行。
这并不是真正的并发，如果你要真正的并发，
你需要在你的main函数的第一行加上下面的这段代码：
runtime.GOMAXPROCS(4)
**/
var total_tickets int32 = 10;
//上锁之一：使用互斥量
var mutex = &sync.Mutex{} //可简写成：var mutex sync.Mutex

func sell_ticketsNotSafe(i int){
    for{
        if total_tickets > 0 { //如果有票就卖
            time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
            total_tickets-- //卖一张票
            fmt.Println("id:", i, "  ticket:", total_tickets)
        }else{
            break
        }
    }
}

func oper (){
	runtime.GOMAXPROCS(4) //我的电脑是4核处理器，所以我设置了4
    rand.Seed(time.Now().Unix()) //生成随机种子
 
    for i := 0; i < 5; i++ { //并发5个goroutine来卖票
//         go sell_ticketsNotSafe(i)
		go sell_ticketsSafe(i)
    }
    //等待线程执行完
    var input string
    fmt.Scanln(&input)
    fmt.Println(total_tickets, "done") //退出时打印还有多少票
}

func sell_ticketsSafe(i int){
	for total_tickets > 0 {
		mutex.Lock()
        if total_tickets > 0 { //如果有票就卖
            time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
            total_tickets-- //卖一张票
            fmt.Println("id:", i, "  ticket:", total_tickets)
        }
        mutex.Unlock()
    }
}

func main(){
	oper()
}
