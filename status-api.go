package main

import (
	"math"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/loadavg"
	"github.com/mackerelio/go-osstat/memory"
)

type ServerStatus struct {
	ramPercent float64
	cpuPercent float64
	cpuAvg1    float64
	cpuAvg5    float64
	cpuAvg15   float64
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/status", func(c *gin.Context) {
		var serverStatus ServerStatus

		// get ram percent
		memory, err := memory.Get()
		if err == nil {
			serverStatus.ramPercent = math.Round(float64(memory.Used)/float64(memory.Total)*10000) / 100
		}

		// get cpu percent
		res, ok := getCpuUsage()
		if ok {
			serverStatus.cpuPercent = res
		}

		// get cpu average
		avg1, avg5, avg15, ok := getCpuAverage()
		if ok {
			serverStatus.cpuAvg1 = avg1
			serverStatus.cpuAvg5 = avg5
			serverStatus.cpuAvg15 = avg15
		}

		c.JSON(200, gin.H{
			"ram":   serverStatus.ramPercent,
			"cpu":   serverStatus.cpuPercent,
			"cpu1":  serverStatus.cpuAvg1,
			"cpu5":  serverStatus.cpuAvg5,
			"cpu15": serverStatus.cpuAvg15,
		})
	})

	r.Run(":8067")
}

func getCpuUsage() (float64, bool) {
	before, err := cpu.Get()
	if err != nil {
		return 0.00, false
	}

	time.Sleep(time.Duration(2) * time.Second)

	after, err := cpu.Get()
	if err != nil {
		return 0.00, false
	}

	total := float64(after.Total - before.Total)
	cpuPercent := float64(after.User-before.User) / total * 100

	return math.Round(cpuPercent*100) / 100, true
}

func getCpuAverage() (float64, float64, float64, bool) {
	load, err := loadavg.Get()
	if err != nil {
		return 0, 0, 0, false
	}

	cores := runtime.NumCPU()
	avg1 := math.Round(load.Loadavg1/float64(cores)*10000) / 100
	avg5 := math.Round(load.Loadavg5/float64(cores)*10000) / 100
	avg15 := math.Round(load.Loadavg15/float64(cores)*10000) / 100

	return avg1, avg5, avg15, true
}
