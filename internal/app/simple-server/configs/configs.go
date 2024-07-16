package configs

import (
	"sync"
	"time"
)

var (
	config sync.Map
)

func Init() {
	config.Store(DelayKey, int64(0))
	config.Store(MaintenanceKey, false)
}

func SetDelay(delay int64) {
	config.Store(DelayKey, delay)
}

func SetMaintenance(value bool) {
	config.Store(MaintenanceKey, value)
}

func Delay() {
	delay, _ := config.Load(DelayKey)
	if d, ok := delay.(int64); ok {
		time.Sleep(time.Duration(d) * time.Millisecond)
	}
}

func IsMaintenance() bool {
	isMaintenance, _ := config.Load(MaintenanceKey)
	if b, ok := isMaintenance.(bool); ok {
		return b
	}
	return false
}