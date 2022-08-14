package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	StartTime time.Time
}

func (s Stopwatch) Start() {
	if s.StartTime.IsZero() {
		s.StartTime = time.Now()
	}
}

// Тип должен обладать следующими методами:
// Start() — запустить/сбросить секундомер;
// SaveSplit() — сохранить промежуточное время;
// GetResults() []time.Duration — вернуть текущие результаты.

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
