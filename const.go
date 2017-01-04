package gohelper

import "time"

func Today() time.Time {
	t := time.Now().Unix()
	return time.Unix(t-(t+28800)%86400, 0)
}
