package cntime

import (
	"fmt"
	"strings"
	"time"
)

var BEIJING = time.FixedZone("CST", 8*3600)

func GetYearMonthDay(t time.Time) (string, string, string) {
	str := t.Format("2006-01-02")
	s := strings.Split(str, "-")

	return s[0], s[1], s[2]
}

func NowCN() time.Time {
	// fmt.Println("")
	// fmt.Println("国际时间为", time.Now().UTC().Format("2006-01-02 15:04:05"))
	// fmt.Println("中国时间为", time.Now().In(BEIJING).Format("2006-01-02 15:04:05"))
	return time.Now().In(BEIJING)
}
func PrintNow() {
	fmt.Println("国际时间", time.Now().UTC().Format("2006-01-02 15:04:05"))
	fmt.Println("中国时间", time.Now().In(BEIJING).Format("2006-01-02 15:04:05"))
	fmt.Println("")
}
