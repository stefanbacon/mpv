package mpvtime

import (
	"math"
	"fmt"
)

type Time struct {
	hh float64
	mm float64
	ss float64
}

func GetTime(seconds float64) Time {
	time := Time{0, 0, seconds}
	if time.ss >= 60 {
		time.ss = math.Mod(time.ss, 60)
		time.mm = (seconds - time.ss) / 60
		if time.mm >= 60 {
			minutes := time.mm
			time.mm = math.Mod(time.mm, 60)
			time.hh = (minutes - time.mm) / 60
		}
	}
	return time
}

func PrintTime (time Time) string {
	return fmt.Sprintf("%02.0f:%02.0f:%02.0f", time.hh, time.mm, time.ss)
}
