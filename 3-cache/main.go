//内存缓存系统
package main

import (
	"fmt"
	cache_server "memCache/cache-server"
	"time"
)

func main() {

	cache := cache_server.NewMemCache()
	cache.SetMaxMemory("200MB")
	cache.Set("int", 1, time.Second)
	cache.Set("bool", false, time.Second)
	cache.Set("data", map[string]interface{}{"a": 1}, time.Second)
	cache.Set("int", 1)
	cache.Set("bool", false)
	cache.Set("data", map[string]interface{}{"a": 1})
	cache.Get("int")
	cache.Del("int")
	cache.Flush()
	fmt.Println(cache.Keys())

	cache.SetMaxMemory("200MB")
	cache.Set("int", 1, time.Second)
	cache.Set("bool", false, time.Second)
	cache.Set("data", map[string]interface{}{"a": 1}, time.Second)
	cache.Set("int", 1)
	cache.Set("bool", false)
	cache.Set("data", map[string]interface{}{"a": 1})
	fmt.Println(cache.Get("int"))
	fmt.Println(cache.Get("bool"))
	fmt.Println(cache.Keys())

	/*cache.GetValueSize(1)
	cache.GetValueSize(false)
	cache.GetValueSize("zhangsan")
	cache.GetValueSize(map[string]string{
		"name": "张三",
		"addr": "1111111111",
	})*/

}
