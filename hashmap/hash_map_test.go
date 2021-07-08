package hashmap

import "testing"

func TestConcurrentMap(t *testing.T) {
	cm := NewConcurrentMap(42)
	cm.Set("Golang", "Cool")
	t.Log(cm.Get("Golang"))
}
