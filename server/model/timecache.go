package model

import (
	"sync"
	"time"
)

const (
	DateFormat = "20060102"
)

// timeFormatCacheType is a time formated cache
type timeFormatCacheType struct {
	// current time
	now time.Time
	// current date
	date string
	// yesterdate
	dateYesterday string

	// lock for read && write
	lock *sync.RWMutex
}

// global time cache instance used for every log writer
var timeCache = timeFormatCacheType{}

func init() {
	timeCache.lock = new(sync.RWMutex)
	timeCache.now = time.Now()
	timeCache.date = timeCache.now.Format(DateFormat)
	timeCache.dateYesterday = timeCache.now.Add(-24 * time.Hour).Format(DateFormat)

	// update timeCache every seconds
	go func() {
		// tick every seconds
		t := time.Tick(1 * time.Second)

		//UpdateTimeCacheLoop:
		for {
			select {
			case <-t:
				timeCache.fresh()
			}
		}
	}()
}

// Now now
func (timeCache *timeFormatCacheType) Now() time.Time {
	timeCache.lock.RLock()
	defer timeCache.lock.RUnlock()
	return timeCache.now
}

// Date date
func (timeCache *timeFormatCacheType) Date() string {
	timeCache.lock.RLock()
	defer timeCache.lock.RUnlock()
	return timeCache.date
}

// DateYesterday date
func (timeCache *timeFormatCacheType) DateYesterday() string {
	timeCache.lock.RLock()
	defer timeCache.lock.RUnlock()
	return timeCache.dateYesterday
}

// fresh data in timeCache
func (timeCache *timeFormatCacheType) fresh() {
	timeCache.lock.Lock()
	defer timeCache.lock.Unlock()

	// get current time and update timeCache
	now := time.Now()
	timeCache.now = now
	date := now.Format(DateFormat)
	if date != timeCache.date {
		timeCache.dateYesterday = timeCache.date
		timeCache.date = now.Format(DateFormat)
	}
}
