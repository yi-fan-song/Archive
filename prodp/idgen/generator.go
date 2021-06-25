package idgen

import (
	"fmt"
	"sync"
	"time"
)

type Generator struct {
	sync.Mutex
	epoch   uint64
	counter uint16
	id      uint16
}

const defaultEpoch uint64 = 1609459200000 // 2021-01-01 00:00:00 +0000 UTC in milliseconds

var defaultGen *Generator

func init() {
	defaultGen, _ = New(0)
}

func New(id uint16) (*Generator, error) {
	if id >= 512 {
		return nil, fmt.Errorf("")
	}

	return &Generator{
		epoch:   defaultEpoch,
		counter: 0,
		id:      id,
	}, nil
}

// Generate generates an time sortable id
// use 45 bits for timestamp in millisecond (over 1000 years after epoch)
// 9 bits for a process id for a total of 512 processes that can generate an Id
// 10 bits for within a millisecond increment
func (g *Generator) Generate() uint64 {
	g.Lock()
	defer g.Unlock()

	ts := uint64(time.Now().UnixMilli()) - g.epoch

	var newId uint64 = 0

	newId += ts << 19
	newId += uint64(g.id) << 10
	newId += uint64(g.counter % 1024)
	g.counter += 1

	return newId
}

// Generate an id using the default generator (id: 0)
func Generate() uint64 {
	return defaultGen.Generate()
}
