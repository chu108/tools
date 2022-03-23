package _time

import (
	"github.com/chu108/tools/_conf"
	"github.com/chu108/tools/_num"
	"log"
	"time"
)

const (
	LayoutDate     = "2006-01-02"
	LayoutTime     = "15:04:05"
	LayoutDateTime = "2006-01-02 15:04:05"
)

type Time struct {
}

func NewTime() *Time {
	return &Time{}
}

// GetStrDate 返回年月日
func (*Time) GetStrDate() string {
	return time.Now().Format(LayoutDate)
}

// GetStrTime 返回年月日时分秒
func (*Time) GetStrTime() string {
	return time.Now().Format(LayoutTime)
}

// GetStrDateTime 返回年月日时分秒
func (*Time) GetStrDateTime() string {
	return time.Now().Format(LayoutDateTime)
}

// UnixToDateTime 时间戳转日期
func (*Time) UnixToDateTime(t int64) string {
	return time.Unix(t, 0).Format(_conf.LayoutTime)
}

// DateTimeToUnix 日期字符串转时间戳
func (*Time) DateTimeToUnix(t string) (int64, error) {
	tm, err := time.Parse(t, _conf.LayoutTime)
	if err != nil {
		return 0, err
	}
	return tm.Unix(), nil
}

// MilliSecondToDateTime 毫秒转日期
func (*Time) MilliSecondToDateTime(t int64) string {
	return time.Unix(0, t*int64(time.Millisecond)).Format(_conf.LayoutTime)
}

// SleepRange 随机睡眠指定范围毫秒
func (*Time) SleepRange(max, min int) {
	time.Sleep(time.Millisecond * time.Duration(_num.NewRand().RandRange(min, max)))
}

// RunTime 计算程序运行时间
func (*Time) RunTime(callback func()) time.Duration {
	st := time.Now()
	callback()
	return time.Since(st)
}

func (*Time) GetLocal(local string) *time.Location {
	if local == "" {
		return time.UTC
	}
	l, err := time.LoadLocation(local)
	if err != nil {
		log.Fatal("GetLocal.err: ", err)
		return &time.Location{}
	}
	return l
}

func (obj *Time) GetTimeByStr(str, local string) time.Time {
	t, err := time.ParseInLocation(LayoutTime, str, obj.GetLocal(local))
	if err != nil {
		log.Fatal("GetTimeByStr.err: ", err)
		return time.Time{}
	}
	return t
}

func (obj *Time) GetDateByStr(str, local string) time.Time {
	t, err := time.ParseInLocation(LayoutDate, str, obj.GetLocal(local))
	if err != nil {
		log.Fatal("GetDateByStr.err: ", err)
		return time.Time{}
	}
	return t
}

func (obj *Time) GetDateTimeByStr(str, local string) time.Time {
	t, err := time.ParseInLocation(LayoutDateTime, str, obj.GetLocal(local))
	if err != nil {
		log.Fatal("GetDateTimeByStr.err: ", err)
		return time.Time{}
	}
	return t
}
