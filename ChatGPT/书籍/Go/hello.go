package main

import (
	"fmt"
	"sync"
)

var (
	inventory = 100
	wg        sync.WaitGroup
	mutex     sync.Mutex
)

func buy() {
	defer wg.Done()
	mutex.Lock()
	if inventory > 0 {
		// 模拟购买过程
		fmt.Println("购买成功,库存剩余：", inventory)

		inventory--
	} else {
		fmt.Println("库存不足")
	}
	mutex.Unlock()
}

func main() {
	wg.Add(200)
	for i := 0; i < 200; i++ {
		go buy()
	}
	wg.Wait()
	fmt.Println("库存剩余：", inventory)
}
