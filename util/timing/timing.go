package timing

import (
	"time"
)

//
// Timing
//  @Description: 统计 fn 执行时长
//  @param fn
//  @return sub
//
func Timing(fn func()) (sub time.Duration) {
	begin := time.Now()
	fn()
	end := time.Now()

	sub = end.Sub(begin)
	return
}
