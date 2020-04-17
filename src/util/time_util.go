package util

import "time"

var JST = time.FixedZone("Asia/Tokyo", 9*60*60)

func GetExpire(d time.Duration) time.Time {
	return time.Now().In(JST).Add(d)
}
