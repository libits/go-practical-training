/*
实现一个内存缓存系统
支持设定过期时间 精确到秒
支持设定最大内存 当内存超出时做出合适的处理
支持并发安全
*/
package cache

import (
	"time"
)

type Cache interface {
	//size: 1IKB 100KB.....
	SetMaxMemory(size string) bool
	//将value写入缓存
	Set(key string, val interface{}, expire time.Duration) bool
	//根据key获得value值
	Get(key string) (interface{}, bool)
	//删除key值
	Del(key string) bool
	//判断key是否存在
	Exists(key string) bool
	//清空所有key
	Flush() bool
	//获取缓存中所有key的数量
	Keys() int64
}
