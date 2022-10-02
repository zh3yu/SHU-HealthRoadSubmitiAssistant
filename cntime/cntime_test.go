package cntime_test

import (
	"fmt"
	"selfreport/cntime"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	time0 := cntime.NowCN()
	fmt.Println("今天的时间是", time0)
	fmt.Println("昨天的时间是", time0.Add(-time.Hour*24))
	time1, _ := time.ParseInLocation("2006-1-2 15:04:05", "2019-7-20 07:14:44", cntime.BEIJING)
	fmt.Println("北京时间 2019-7-20 的时间是", time1)
	fmt.Println("北京时间 2019-7-20 的国际时间是", time1.UTC())
	fmt.Println(cntime.GetYearMonthDay(time1))
}
