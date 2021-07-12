# dconcurrent-map
Golang 高性能并发HashMap实现，弥补 sync.Map 性能上的缺陷。

map[interface{}]interface{}
## 接口介绍

```
cm := NewConcurrentMap(42)

cm.Set("Golang", "Cool")

val := cm.Get("Golang")

cm.Del("Golang")
```

## 性能优化分析 

1. 数据分段，减少锁粒度

## TODO