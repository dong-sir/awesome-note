package time

import (
	"time"
)

// 当前时间格式化
time.Now().Format("2006-01-02-15-04-05")

// 7 天前
const DaysToKeep = 7
updateTime := time.Now().AddDate(0, 0, -DaysToKeep).Unix()
fmt.Println(updateTime)