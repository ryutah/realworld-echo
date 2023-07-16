package xtime

import "time"

var nowFunc = time.Now

func SetNowFunc(f func() time.Time) (resetFunc func()) {
	nowFunc = f
	return func() {
		nowFunc = time.Now
	}
}

func Now() time.Time {
	return nowFunc()
}

var asiaTokyo = time.FixedZone("Asia/Tokyo", 9*60*60)

func JST(t time.Time) time.Time {
	return t.In(asiaTokyo)
}
