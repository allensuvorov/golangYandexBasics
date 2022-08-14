package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	StartTime time.Time
	Split     []time.Duration
}

// Start() — запустить/сбросить секундомер;
func (sw *Stopwatch) Start() {
	if sw.StartTime.IsZero() {
		sw.StartTime = time.Now() // save current time
	} else {
		sw.StartTime = time.Time{} // reset
	}
}

// SaveSplit() — сохранить промежуточное время;
func (sw *Stopwatch) SaveSplit() {
	if !sw.StartTime.IsZero() {
		sw.Split = append(sw.Split, time.Now().Sub(sw.StartTime))
	}
}

// Тип должен обладать следующими методами:
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
