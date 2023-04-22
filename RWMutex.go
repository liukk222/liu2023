/* package main

import "sync"

var count int
var rwmutex sync.RWMutex

func read() {
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	// 读取 count 变量的值
}

func write() {
	rwmutex.Lock()
	defer rwmutex.Unlock()
	// 修改 count 变量的值
}

func main() {
	for i := 0; i < 10; i++ {
		go read()
	}

	for i := 0; i < 3; i++ {
		go write()
	}

	select {}
}
*/