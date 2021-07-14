# dconcurrent-map

fork from [concurrent-map](https://github.com/orcaman/concurrent-map)

## API

```
cm := NewConcurrentMap(42)

cm.Set("Golang", "Cool")

val := cm.Get("Golang")

cm.Del("Golang")
```

## TODO
1. change key type to interface{} will cause performance loss  
