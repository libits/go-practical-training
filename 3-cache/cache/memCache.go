package cache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type memCache struct {
	//最大内存
	maxMemorySize int64
	//最大内存字符串显示
	maxMemorySizeStr string
	//当前已使用内存大小
	currMemorySize int64

	//缓存键值对
	values map[string]*memCacheValue
	//锁
	locker sync.RWMutex

	//清除过期缓存时间间隔
	clearExpiredItemTimeInerval time.Duration
}
type memCacheValue struct {
	//value值
	val interface{}
	//过期时间
	expireTime time.Time
	//有效时长
	expire time.Duration
	//value大小
	size int64
}

//实例化对象
func NewMemCache() Cache {
	mc := &memCache{
		values:                      make(map[string]*memCacheValue),
		clearExpiredItemTimeInerval: time.Second * 10,
	}
	go mc.clearExpiredItem()
	return mc
}

//size :1KB 1MB .....

func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)
	fmt.Println(mc.maxMemorySize, mc.maxMemorySizeStr)
	return false
}

//将value写入缓存
func (mc *memCache) Set(key string, val interface{}, expire time.Duration) bool {

	mc.locker.Lock()
	defer mc.locker.Unlock()
	fmt.Println("called set")
	v := &memCacheValue{
		val:        val,
		expireTime: time.Now().Add(expire),
		expire:     expire,
		size:       GetValueSize(val),
	}
	mc.del(key)
	mc.add(key, v)
	if mc.currMemorySize > mc.maxMemorySize {
		mc.del(key)
		log.Println(fmt.Sprintln("max memory size %d", mc.maxMemorySize))
		panic(fmt.Sprintln("max memory size %d", mc.maxMemorySize))
	}

	return false
}
func (mc *memCache) get(key string) (*memCacheValue, bool) {
	val, ok := mc.values[key]
	return val, ok
}
func (mc *memCache) del(key string) {
	tmp, ok := mc.get(key)
	if ok && tmp != nil {
		mc.currMemorySize -= tmp.size
		delete(mc.values, key)
	}
}
func (mc *memCache) add(key string, val *memCacheValue) {
	mc.values[key] = val
	mc.currMemorySize += val.size
}

//根据key值获得value
func (mc *memCache) Get(key string) (interface{}, bool) {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	mcv, ok := mc.get(key)
	if ok {
		//判断缓存是否过期
		if mcv.expire != 0 && mcv.expireTime.Before(time.Now()) {
			mc.del(key)
			return nil, false
		}
		return mcv.val, ok
	}
	return nil, false
}

//删除key值
func (mc *memCache) Del(key string) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.del(key)
	return true
}

//判断key是否存在
func (mc *memCache) Exists(key string) bool {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	_, ok := mc.values[key]
	return ok
}

//清空所有key
func (mc *memCache) Flush() bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	//golang的回收机制
	mc.values = make(map[string]*memCacheValue, 0)
	mc.currMemorySize = 0
	return true
}

//获取缓存中所有key的数量
func (mc *memCache) Keys() int64 {
	mc.locker.Lock()
	defer mc.locker.Unlock()

	return int64(len(mc.values))
}
func (mc *memCache) clearExpiredItem() {
	timeTicker := time.NewTicker(mc.clearExpiredItemTimeInerval)
	defer timeTicker.Stop()
	for {
		select {
		case <-timeTicker.C:
			for key, item := range mc.values {
				if item.expire != 0 && time.Now().After(item.expireTime) {
					mc.locker.Lock()
					mc.del(key)
					mc.locker.Unlock()
				}
			}
		}
	}
}
