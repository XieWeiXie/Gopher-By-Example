package main

import (
	"fmt"
	"time"
)

/*
1. 单个时间操作
2. 两个时间操作
3. 时间和字符串之间的操作，相互转换，或者格式化操作
4. 枚举类型

*/

var currentTime = func() time.Time {
	return time.Now()
}

var printTimeInfo = func() {
	now := time.Now()

	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Date())

	fmt.Println(now.UnixNano())
	fmt.Println(now.Unix())

	fmt.Println(time.Unix(1542119403, 1542119403000000000).UTC())

	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.AddDate(1, 0, 0))
	fmt.Println(now.AddDate(1, 1, 0))
	fmt.Println(now.AddDate(1, 0, 1))
	fmt.Println(now.AddDate(-1, 0, 0))
	fmt.Println(now, now.Add(-2*time.Hour))
	fmt.Println(now, now.Add(-2*time.Minute))
	fmt.Println(now, now.Add(-1000*time.Hour))

	fmt.Println(now, now.Truncate(time.Hour)) // 向下
	fmt.Println(now, now.Round(time.Hour))    // 向上

	timeString := "2019-01-02 15:03:00"

	fmt.Println(time.ParseDuration("1h"))
	fmt.Println(time.Parse("2006-01-02 15:04:05", timeString))
	fmt.Println(time.ParseInLocation("2006-01-02 15:04:05", timeString, time.Local))

	y, w := now.ISOWeek()
	fmt.Println(now.Weekday(), y, w)

	timeStringZero := "2019-01-01 00:00:00"
	fmt.Println(now.IsZero())

	dd, _ := time.Parse("2006-01-02 15:04:05", timeStringZero)
	fmt.Println(time.Parse("2006-01-02 15:04:05", timeStringZero))
	fmt.Println(dd.IsZero(), dd)

	fmt.Println(dd.After(now))
	fmt.Println(now.After(dd))
	fmt.Println(now.Before(dd))
	fmt.Println(dd.Before(now))
}

func main() {

	//fmt.Println(currentTime())
	//printTimeInfo()
	//timeToString()
	//timeOp()
	//timeOpDate()
	//timeInterval()

	fmt.Println(time.Now().ISOWeek())
	fmt.Println(time.Now().Weekday())

	fmt.Println(Red,
		Orange,
		Yellow,
		Green,
		Cyan,
		Blue,
		Purple)
}

var timeToString = func() {

	now := time.Now()

	fmt.Println(now.Format("2006-01-02 15:00:00"))
	fmt.Println(now.Format("2006-01-02 15:04:00"))
	fmt.Println(now.Format("2006-01-02 00:00:00"))

	fmt.Println(now.Round(time.Hour))
	fmt.Println(now.Round(time.Minute))
	fmt.Println(now.Round(time.Second))
	fmt.Println(now.Truncate(time.Hour))
	fmt.Println(now.Truncate(time.Minute))
	fmt.Println(now.Truncate(time.Second))

}

var timeOp = func() {
	now := time.Now()

	fmt.Println(now.Add(1 * time.Hour))
	fmt.Println(now.Add(24 * time.Hour))
	fmt.Println(now.Add(-10 * time.Hour))
}

var timeOpDate = func() {
	now := time.Now()

	fmt.Println(now.AddDate(-1, 0, 0))
	fmt.Println(now.AddDate(2, 0, 0))
	fmt.Println(now.AddDate(0, 1, 0))
}

var timeInterval = func() {
	now := time.Now()
	stringTime := "2018-11-14 20:00:00"

	newTime, _ := time.ParseInLocation("2006-01-02 15:04:05", stringTime, time.Local)

	if newTime.After(now) {
		subTime := newTime.Sub(now)
		fmt.Println("newTime after now")
		fmt.Println(subTime.Hours())
		fmt.Println(subTime.Minutes())
		fmt.Println(subTime.Seconds())
		fmt.Println(subTime.Nanoseconds())
	}
	if newTime.Before(now) {
		subTime := now.Sub(newTime)
		fmt.Println(subTime.String())
	}
}

const (
	Red = iota
	Orange
	Yellow = iota + 10
	Green  = iota
	Cyan
	Blue = 10
	Purple
)
