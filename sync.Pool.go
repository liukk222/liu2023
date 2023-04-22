/* package main

import (
	"fmt"
	"sync"
)

type Object struct {
	data string
}

func NewObject(data string) *Object {
	return &Object{data}
}

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return NewObject("default")
		},
	}

	obj := pool.Get().(*Object)
	fmt.Println(obj.data)

	obj.data = "foo"
	fmt.Println(obj.data)
	pool.Put(obj)

	obj = pool.Get().(*Object)
	fmt.Println(obj.data)
} */

/* package main

import (
	"fmt"
	"sync"
)

var stringPool = sync.Pool{
	New: func() interface{} {
		return ""
	},
}

func main() {
	// 从 sync.Pool 中获取字符串对象
	str1 := stringPool.Get().(string)
	str1 = "Hello, World!"
	fmt.Println(str1)

	// 再次从 sync.Pool 中获取字符串对象
	str2 := stringPool.Get().(string)
	str2 = "str2"
	fmt.Printf("ok:%s\n", str2) // 输出 ""

	// 将字符串对象归还到 sync.Pool 中
	stringPool.Put(str1)
	stringPool.Put(str2)
	str3 := stringPool.Get().(string)
	str4 := stringPool.Get().(string)
	fmt.Printf("str3: %v\n", str3)
	fmt.Printf("str4: %v\n", str4)
}
*/