package main

// 处理多个gorutine运行时必须同时运行同时结束

import (
    "fmt"
    "time"
)

func main() {
    
    ch := make(chan bool, 1)

    ch1 := make(chan bool, 1)

    go func() {
        timeTicker := time.NewTicker(time.Second * 2)

        for i:=0; i < 3; i++ {
        
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

    if a1, _ := <-ch; a1 == true {
        fmt.Println("a1 exit;")        
    }

    if a2, _ := <-ch1; a2 == true {
        fmt.Println("a2 exit;") 
    }

    fmt.Println("are you ok in main")
}
