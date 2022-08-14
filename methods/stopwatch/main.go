package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	startTime time.Time
	splitTime []time.Duration
}

// Start() — запустить/сбросить секундомер;
func (sw *Stopwatch) Start() {
	sw.startTime = time.Now() // save current time
	sw.splitTime = nil
}

// SaveSplit() — сохранить промежуточное время;
func (sw *Stopwatch) SaveSplit() {
	if !sw.startTime.IsZero() {
		sw.splitTime = append(sw.splitTime, time.Now().Sub(sw.startTime))
	}
}

// Тип должен обладать следующими методами:
// GetResults() []time.Duration — вернуть текущие результаты.
func (sw *Stopwatch) GetResults() []time.Duration {
	return sw.splitTime
}

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}
