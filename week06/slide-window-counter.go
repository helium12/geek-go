package week06

import (
	"sync"
	"time"
)

type slot struct {
	start time.Time // 时间槽起点
	count int       // 时间槽内请求数
}

type SlideWindow struct {
	mu        sync.Mutex
	SlotLen   time.Duration // time slot长度
	WindowLen time.Duration // 滑动窗口长度
	slotNum   int           // 时间窗口内有多少个slot
	windows   []*slot
	maxReq    int // 滑动窗口时间内支持的最大请求数
}

func (l *SlideWindow) validate() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	// 移出过期time slot
	timeoutOffset := -1
	for i, ts := range l.windows {
		if ts.start.Add(l.WindowLen).After(now) {
			break
		}
		timeoutOffset = i
	}

	if timeoutOffset > -1 {
		l.windows = l.windows[timeoutOffset+1:]
	}

	var result bool
	// 请求超过最大请求限制
	if countReq(l.windows) < l.maxReq {
		result = true
	}

	// 统计请求数
	var lastSlot *slot
	if len(l.windows) > 0 {
		lastSlot = l.windows[len(l.windows)-1]

		// 新建time slot
		if lastSlot.start.Add(l.SlotLen).Before(now) {
			lastSlot = &slot{start: now, count: 1}
			l.windows = append(l.windows, lastSlot)
		} else {
			lastSlot.count++
		}
	} else {
		lastSlot = &slot{start: now, count: 1}
		l.windows = append(l.windows, lastSlot)
	}

	return result
}

//统计时间窗口内发生请求次数
func countReq(win []*slot) int {
	var count int
	for _, ts := range win {
		count += ts.count
	}
	return count
}

func NewSlide(slotLen time.Duration, winLen time.Duration, maxReq int) *SlideWindow {
	return &SlideWindow{
		SlotLen:   slotLen,
		WindowLen: winLen,
		slotNum:   int(winLen / slotLen),
		maxReq:    maxReq,
	}
}
