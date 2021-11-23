package utils

import (
	"errors"
	"sync"
	"time"
)

const (
	workerBits  uint8 = 10
	numberBits  uint8 = 12
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	startTime   int64 = 1136185445000
)

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerID  int64
	number    int64
}

func (w *Worker) GetID() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6

	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	return (now-startTime)<<timeShift | (w.workerID << workerShift) | (w.number)
}

func NewSnowflake(workerID int64) (*Worker, error) {
	if workerID < 0 || workerID > workerMax {
		return nil, errors.New("Worker ID Excess Of Quantity")
	}
	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		workerID:  workerID,
		number:    0,
	}, nil
}
