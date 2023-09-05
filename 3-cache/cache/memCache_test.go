package cache

import (
	"testing"
	"time"
)

func TestCachOP(t *testing.T) {
	testData := []struct {
		key    string
		val    interface{}
		expire time.Duration
	}{
		{"mndjenfoenf", 7977966, time.Second * 10},
		{"bgrfdhgrfg", false, time.Second * 10},
		{"gerhgaer", map[string]interface{}{"a": 3, "b": false}, time.Second * 12},
		{"mfnmoejrdngujirenbg", "ehnfuhenwuo", time.Second * 10},
		{"snfciujwedbnfiujweb", "这里是字符串", time.Second * 15},
	}
	c := NewMemCache()
	c.SetMaxMemory("10MB")
	for _, item := range testData {
		c.Set(item.key, item.val, item.expire)
		val, ok := c.Get(item.key)
		if !ok {
			t.Error("缓存取值失败")
		}
		if item.key != "gerhgaer" && val != item.val {
			t.Error("缓存取值数据与预期不一致")
		}
		_, ok1 := val.(map[string]interface{})
		if item.key == "gerhgaer" && !ok1 {
			t.Error("缓存数据与预期不一致")
		}
	}
	if int64(len(testData)) != c.Keys() {
		t.Error("缓存数量不一致")
	}

	c.Del(testData[0].key)
	c.Del(testData[1].key)
	if int64(len(testData)) != c.Keys()+2 {
		t.Error("缓存数量不一致")
	}
	time.Sleep(time.Second * 16)
	if c.Keys() != 0 {
		t.Error("过期缓存清空失败")
	}

}
