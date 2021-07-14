# dconcurrent-map

dconcurrent-map based on orcaman's cool job [concurrent-map](https://github.com/orcaman/concurrent-map)

## API

```
// create a brand new thread safe map
sm := New()
	
// set key/val pair to map
sm.Set("Hello", "World")
	
// get the specified element
sm.Get("Hello")
	
// remove the specified key/val pair from map 
sm.Remove("Hello")
	
// returns the number of elements within the map
sm.Count()
	
// checks if map is empty
sm.IsEmpty()
	
// removes all items from map
sm.Clear()
	
// returns all items as map[interface{}]interface{}
sm.Items()
	
// returns all keys as []interface{}
sm.Keys()
	
// looks up an item under specified key
sm.Has("Hello")
	
// removes an element from the map and returns it
sm.Pop("Hello")
	
	
items := map[interface{}]interface{}{
	"Hello":"world",
}
// set all items of parameter  to map
sm.MSet(items)
	
```

## hash function 

`hash()` can many golang build-in type.

```
func hash(key interface{}) uint32 {
	buff, err := toBytes(key)
	if err != nil {
		panic("toBytes() error")
		return 0
	}
	return crc32.ChecksumIEEE(buff)
}

func toBytes(key interface{}) ([]byte, error) {
	bs := make([]byte, 8)
	buff := bytes.NewBuffer(bs)
	switch v := key.(type) {
	case *string:
		return []byte(*v), nil
	case string:
		return []byte(v), nil
	case *bool:
		if *v {
			buff.WriteByte(byte(1))
		} else {
			buff.WriteByte(byte(0))
		}
		return buff.Bytes()[:1], nil
	case bool:
		if v {
			buff.WriteByte(byte(1))
		} else {
			buff.WriteByte(byte(0))
		}
		return buff.Bytes()[:1], nil
	case *int8:
		buff.WriteByte(byte(*v))
		return buff.Bytes()[:1], nil
	case int8:
		buff.WriteByte(byte(v))
		return buff.Bytes()[:1], nil
	case *uint8:
		buff.WriteByte(*v)
		return buff.Bytes()[:1], nil
	case uint8:
		buff.WriteByte(v)
		return buff.Bytes()[:1], nil
	case *int16:
		_ = binary.Write(buff, binary.BigEndian, v)
		return buff.Bytes()[:2], nil
	case int16:
		_ = binary.Write(buff, binary.BigEndian, v)
		return buff.Bytes()[:2], nil
	case *int32:
		_ = binary.Write(buff, binary.BigEndian, v)
		return buff.Bytes()[:4], nil
	case int32:
		_ = binary.Write(buff, binary.BigEndian, v)
		return buff.Bytes()[:4], nil
	case *int64:
		_ = binary.Write(buff, binary.BigEndian, v)
		return buff.Bytes()[:8], nil
	case int64:
		_ = binary.Write(buff, binary.BigEndian, v)
		return buff.Bytes()[:8], nil
	default:
		return nil, errInvalidKeyType
	}
}

```


## TODO
1. change key type to interface{} will cause performance loss  
