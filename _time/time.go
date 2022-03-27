package _time

import (
	"github.com/chu108/tools/_conf"
	"github.com/chu108/tools/_num"
	"log"
	"time"
)

// GetStrDate 返回年月日
func GetStrDate() string {
	return time.Now().Format(_conf.LayoutDate)
}

// GetStrTime 返回年月日时分秒
func GetStrTime() string {
	return time.Now().Format(_conf.LayoutTime)
}

// GetStrDateTime 返回年月日时分秒
func GetStrDateTime() string {
	return time.Now().Format(_conf.LayoutDateTime)
}

// UnixToDateTime 时间戳转日期
func UnixToDateTime(t int64) string {
	return time.Unix(t, 0).Format(_conf.LayoutTime)
}

// DateTimeToUnix 日期字符串转时间戳
func DateTimeToUnix(t string) (int64, error) {
	tm, err := time.Parse(t, _conf.LayoutTime)
	if err != nil {
		return 0, err
	}
	return tm.Unix(), nil
}

// MilliSecondToDateTime 毫秒转日期
func MilliSecondToDateTime(t int64) string {
	return time.Unix(0, t*int64(time.Millisecond)).Format(_conf.LayoutTime)
}

// SleepRange 随机睡眠指定范围毫秒
func SleepRange(max, min int) {
	time.Sleep(time.Millisecond * time.Duration(_num.RandRange(min, max)))
}

// RunTime 计算程序运行时间
func RunTime(callback func()) time.Duration {
	st := time.Now()
	callback()
	return time.Since(st)
}

func GetLocal(local string) *time.Location {
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

func GetTimeByStr(str, local string) time.Time {
	t, err := time.ParseInLocation(_conf.LayoutTime, str, GetLocal(local))
	if err != nil {
		log.Fatal("GetTimeByStr.err: ", err)
		return time.Time{}
	}
	return t
}

func GetDateByStr(str, local string) time.Time {
	t, err := time.ParseInLocation(_conf.LayoutDate, str, GetLocal(local))
	if err != nil {
		log.Fatal("GetDateByStr.err: ", err)
		return time.Time{}
	}
	return t
}

func GetDateTimeByStr(str, local string) time.Time {
	t, err := time.ParseInLocation(_conf.LayoutDateTime, str, GetLocal(local))
	if err != nil {
		log.Fatal("GetDateTimeByStr.err: ", err)
		return time.Time{}
	}
	return t
}
