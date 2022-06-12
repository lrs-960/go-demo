package internal

import (
	"sync"
	"time"
)

type Bucket struct {
	sync.RWMutex
	Total     int
	Failed    int
	Timestamp time.Time
}

func NewBucket() *Bucket {
	return &Bucket{
		Timestamp: time.Now(),
	}
}

// Record 记录请求结果
func (b *Bucket) Record(result bool) {
	b.Lock()
	defer b.Unlock()
	if !result {
		b.Failed++
	}
	b.Total++
}
