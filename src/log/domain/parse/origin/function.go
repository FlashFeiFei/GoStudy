package origin

import (
	"strconv"
	"sync/atomic"
	"time"
)

/**
获取年月日
 */
func GetYMD() string {

	return time.Now().Format("2006-01-02")
}

func GetFileName(counter *int64) string {
	//获取毫秒级的
	var fileName string
	//获取毫秒级时间戳
	t := time.Now()
	//将毫秒级时间戳转成字符串
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)

	//原子新的对counter增加1
	new_counter := atomic.AddInt64(counter, 1)

	new_counter_change_string := strconv.FormatInt(new_counter, 10)


	fileName = timestamp + "_" + new_counter_change_string

	return fileName
}
