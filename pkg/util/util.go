package util

import (
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func SafeDeref[T any](p *T) T {
	if p == nil {
		var v T
		return v
	}
	return *p
}

func GenerateFromPassword(secret string) (string, error) {
	s, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

func ComparePassword(hash, secret string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
}

func ParseTime(at string) (time.Time, error) {
	return time.Parse("200601021504", at)
}

func Now() string {
	now := time.Now()
	return now.Format("20060102150405")
}

func Unix() int64 {
	now := time.Now()
	return now.Unix()
}

func NowInAsia() int64 {
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)
	return now.UnixNano()
}

func BeforeOneDayInAsia() int64 {
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)
	return now.Add(-24 * time.Hour).UnixNano()
}

func ConvertToMillisecond(millisecond string) (string, error) {
	millisInt, err := strconv.ParseInt(millisecond, 10, 64)
	if err != nil {
		return "", err
	}
	t := time.Unix(0, millisInt*int64(time.Millisecond))
	return t.Format("2006-01-02 15:04:05"), nil
}
