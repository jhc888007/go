package main

import "fmt"

func PrintCounter(i int, ch chan int) {
    fmt.Println(i);
    ch <- 1;
}

func main() {
    var chs []chan int

    for i := 0; i < 5; i++ {
        chs = append(chs, make(chan int))
        fmt.Println("chs len:", len(chs), ", cap:", cap(chs))
        go PrintCounter(i, chs[i])
    }

    /*for i := 0; i < 10; i++ {
        <-chs[i];
    }*/

    for _, ch := range(chs) {
        <-ch
    }

    chs[2] = nil;
    fmt.Println("chs len:", len(chs), ", cap:", cap(chs));
    for i, ch := range(chs) {
        if ch == nil {
            fmt.Println("nil")
        } else {
            fmt.Println(i)
        }
    }

    chs = append(chs[0:2],nil);
    fmt.Println("chs len:", len(chs), ", cap:", cap(chs));
    for i, ch := range(chs) {
        if ch == nil {
            fmt.Println("nil")
        } else {
            fmt.Println(i)
        }
    }

}
