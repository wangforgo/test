package main

import (
	"fmt"
	"runtime"
	"time"
)

var x [][]*int
func initData() {
	x = make([][]*int,100)
	for i:= range x {
		x[i] = make([]*int,10000000)
	}
}

func main() {
	fmt.Println(runtime.Version())
	runtime.GOMAXPROCS(4)

	// initData for long GC mark time
	initData()

	go ticker(200)
	for i:=0;i<4;i++ {
		go busyTask()
	}

	go func() {
		for {
			time.Sleep(time.Second)
			runtime.GC()
		}
	}()

	time.Sleep(time.Second*20)
}



func ticker(periodms int64) {
	tc := time.NewTicker(time.Duration(periodms) * time.Millisecond)
	ticks := 0
	var lastTimestamp time.Time
	for t := range tc.C {
		if ticks == 0 {
			lastTimestamp = t
		} else {
			elapsedTime := t.Sub(lastTimestamp).Milliseconds()
			if elapsedTime > periodms * 2 {
				fmt.Printf("%v: %v ======== TIMER WAKEUP TOO LATE!!!\n",ticks,elapsedTime)
			} else {
				fmt.Printf("%v: %v\n",ticks,elapsedTime)
			}
			lastTimestamp = t
		}
		ticks ++
	}
}

func busyTask() {
	for {

	}
}