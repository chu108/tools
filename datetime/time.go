package datetime

import (
	"github.com/chu108/tools/config"
	"github.com/chu108/tools/number"
	"time"
)

type Time struct {
}

func NewTime() *Time {
	return &Time{}
}

//返回年月日
func (*Time) GetDate() string {
	return time.Now().Format(config.LayoutDate)
}

//返回年月日时分秒
func (*Time) GetTime() string {
	return time.Now().Format(config.LayoutTime)
}

//时间戳转日期
func (*Time) UnixToDateTime(t int64) string {
	return time.Unix(t, 0).Format(config.LayoutTime)
}

//日期字符串转时间戳
func (*Time) DateTimeToUnix(t string) (int64, error) {
	tm, err := time.Parse(t, config.LayoutTime)
	if err != nil {
		return 0, err
	}
	return tm.Unix(), nil
}

//毫秒转日期
func (*Time) MilliSecondToDateTime(t int64) string {
	return time.Unix(0, t*int64(time.Millisecond)).Format(config.LayoutTime)
}

//随机睡眠指定范围毫秒
func (*Time) SleepRange(max, min int) {
	time.Sleep(time.Millisecond * time.Duration(number.NewRand().RandRange(min, max)))
}

//计算程序运行时间
func (*Time) RunTime(callback func()) time.Duration {
	st := time.Now()
	callback()
	return time.Since(st)
}
