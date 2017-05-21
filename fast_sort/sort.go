package main

import (
	"fmt"
	"sync"
	"time"
)

func dump(a []int) {
	for i, v := range a {
		fmt.Printf("%d:%d ", i, v)
	}
	fmt.Println()
}

/*func sort1(b []int, wg sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	k := len(b) - 1
	if k <= 0 {
		return
	}
	i := 0
	j := 1
	for {
		if j >= k {
			if b[k] > b[i] {
				k -= 1
			}
			if k != i {
				fmt.Println(k, "-", b[k], " ", i, "-", b[i])
				b[k], b[i] = b[i], b[k]
			}
			break
		}
		if b[j] <= b[i] {
			j += 1
			continue
		}
		if b[k] >= b[i] {
			k -= 1
			continue
		}
		fmt.Println(j, "-", b[j], " ", k, "-", b[k])
		b[j], b[k] = b[k], b[j]
	}
	go sort1(b[:k], wg)
	go sort1(b[k+1:], wg)
}*/

func main() {
	var a []int = []int{12, 4, 3, 5, 67, 32, 13, 11, 2, 1, 45, 88, 6, 15, 34, 22}
	dump(a)
	var wg sync.WaitGroup
	var sort func(b []int) //必须先声明，否则不能递归调用
	sort = func(b []int) {
		wg.Add(1)
		defer wg.Done()
		k := len(b) - 1
		if k <= 0 {
			return
		}
		i := 0
		j := 1
		for {
			if j >= k {
				if b[k] > b[i] {
					k -= 1
				}
				if k != i {
					b[k], b[i] = b[i], b[k]
				}
				break
			}
			if b[j] <= b[i] {
				j += 1
				continue
			}
			if b[k] >= b[i] {
				k -= 1
				continue
			}
			b[j], b[k] = b[k], b[j]
		}
		go sort(b[:k])
		go sort(b[k+1:])
	}
	//go sort1(a, wg)
	go sort(a)
	time.Sleep(time.Second)
	wg.Wait()
	dump(a)
}
