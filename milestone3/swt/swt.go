package swt

import (
	"fmt"
	"time"
)

func CurrentTime() string {
	secondsInABeat := 86.4
	loc, _ := time.LoadLocation("UTC")
	t := time.Now().In(loc)
	h := t.Hour()
	m := t.Minute()
	s := t.Second()

	if h == 23 {
		h = 0
	} else {
		h++
	}

	timeInSeconds := (((h * 60) + m) * 60) + s
	beats := float64(timeInSeconds) / secondsInABeat

	return fmt.Sprintf("%s@%.2f", t.Format("2006-01-02"), beats)
}
