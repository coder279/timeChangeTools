package internal

import "time"

// GetnowTime 获取本地当前时间
func GetnowTime() time.Time {
	location,_:= time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

// GetCalculateTime 计算间隔时间
func GetCalculateTime(currentTimer time.Time,d string)(time.Time,error){
	duration,err := time.ParseDuration(d)
	if err != nil {
		return time.Time{},err
	}
	return currentTimer.Add(duration),nil
}