package benchmark

import (
	"fmt"
	"testing"
	"time"

	"github.com/xusworld/ConcurrentMap/hashmap"
	"github.com/xusworld/infinity-go/benchmark"
)

func batchReadOnlyBenchmark(hashMaps []hashmap.HashMap, goroutinesConf []int, readTimes int, hint []string) {
	for _, goroutinesNum := range goroutinesConf {
		durations := make([]time.Duration, 0)

		for _, m := range hashMaps {
			durations = append(durations, readOnlyBenchmark(m, goroutinesNum, readTimes))
		}

		fmt.Printf("Concurrency %d, ReadTimes %d\n", goroutinesNum, readTimes)
		for idx, duration := range durations {
			fmt.Printf("TimeCost %s [%s]  \n", duration.String(), hint[idx])
		}
		fmt.Println()
	}
}

// readOnlyBenchmark 对一个HashMap在 使用channel编排任务
func readOnlyBenchmark(m hashmap.HashMap, goroutinesNum int, readTimes int) time.Duration {
	start := time.Now()

	finished := make(chan struct{}, goroutinesNum)

	for i := 0; i < goroutinesNum; i++ {
		go func() {
			for j := 0; j < readTimes; j++ {
				_ = m.Get("MagicNumber")
			}
			// 子goroutines执行结束之后向main goroutine写入信号
			finished <- struct{}{}
		}()
	}

	// main goroutines等待所有子goroutines执行结束
	for i := 0; i < goroutinesNum; i++ {
		<-finished
	}
	close(finished)

	return time.Now().Sub(start)
}


func TestBatchReadOnly(t *testing.T) {
	// 并发读取数据的goroutines数量
	concurrency := []int{8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192}
	// 单个goroutines中循环读取数据的次数
	readTimes := 10000

	bim := benchmark.NewBuildInMap()
	bim.Set("MagicNumber", 42)

	sm := &benchmark.SyncMap{}
	sm.Set("MagicNumber", 42)

	rwm := benchmark.NewRWMutexMap()
	rwm.Set("MagicNumber", 42)

	mm := benchmark.NewMutexMap()
	mm.Set("MagicNumber", 42)


	newCM := hashmap.NewConcurrentMap(32)
	newCM.Set("MagicNumber", 42)

	cm := benchmark.NewConcurrentMap()
	cm.Set("MagicNumber", 42)


	batchReadOnlyBenchmark([]hashmap.HashMap{sm, bim, rwm, mm, newCM, cm}, concurrency, readTimes,
		[]string{"sync.Map", "build-in map", "RWMutexMap", "MutexMap", "New concurrent-map", "concurrent-map"})
}
