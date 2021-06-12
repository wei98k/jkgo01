package main

// 模拟多个gorutine在其中某几个gorutine退出后 其他的gorutine还在运行

import (
    "fmt"
    "time"
)

func main() {
    
    ch := make(chan bool, 1)

    ch1 := make(chan bool, 1)

    go func() {
        timeTicker := time.NewTicker(time.Second * 2)

        for i:=0; i < 10; i++ {
        
            fmt.Println(i)

            <-timeTicker.C
        }
    
        timeTicker.Stop()

        ch <- true
       
    }()


    go func() {
        
        timeTicker := time.NewTicker(time.Second * 2)

        for i:=0; i < 5; i++ {
            
            fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

            <-timeTicker.C
        }

        timeTicker.Stop()
        
        ch1 <- true
    }()
    
    <-ch
    <-ch1

    fmt.Println("are you ok in main")
}
