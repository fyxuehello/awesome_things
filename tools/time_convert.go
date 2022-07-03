package tools

import (
	"fmt"
	"time"
)

// 时间转字符串

func Time2String(t time.Time) string {
	timeStr := t.Format("2006-01-02 15:04:05")
	return timeStr
}

// 字符串转时间
func String2Time(str string) time.Time {
	// 注意此处有坑，一般都是用time.Parse方法，time.Parse方法是转为格林威治时间的，也就是0时区，再转为时间戳就就加上时区了
	var tt time.Time
	var err error
	tt, err = time.ParseInLocation("2006-01-02 15:04:05", str, time.Local) // 按照当前时区转换
	if err != nil {
		fmt.Printf("Time convert failed. %s\n", err)
		return tt
	}
	return tt
}

// 获取时间戳
func Time2TimesStamp(tt time.Time) {
	sec := tt.Unix() // 秒
	fmt.Printf("秒：%v\n", sec)
	millSec := tt.UnixMilli() // 纳秒
	fmt.Printf("纳秒：%v\n", millSec)
}
