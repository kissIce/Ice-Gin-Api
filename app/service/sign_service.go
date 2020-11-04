package service
import (
	"errors"
	"ice/utils/binary"
	"time"
)

/**
 * 签到今天
 */
func SignToday(num int) (int, error) {
	dayOfMonth := time.Now().Day()
	return PatchSignDay(num, dayOfMonth)
}

/**
 * 签到指定日期
 */
func PatchSignDay(num int, day int) (int, error) {
	isSign := binary.GetBit(num, day - 1)
	if isSign {
		return 0, errors.New("您已签到")
	}
	num = binary.SetBitOne(num, day - 1)
	return num, nil
}

/**
 * 获取今天签到情况
 */
func GetTodaySignStatus(num int) bool {
	dayOfMonth := time.Now().Day()
	return GetSignStatusWithDay(num, dayOfMonth)
}

/**
 * 获取指定本月指定日期签到情况
 */
func GetSignStatusWithDay(num int, day int) bool {
	return binary.GetBit(num, day - 1)
}

/**
 * 统计本月所有签到天数
 */
func CountAllSign(num int) int {
	return binary.CountBitOne(num)
}

/**
 * 统计本月最大连续签到天数
 */
func CountSuccessiveSign(num int) int {
	return binary.CountSuccessiveOne(num)
}
