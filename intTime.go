package mutt

import "time"

const (
	TIME_FORMAT = "2006/01/02 15:04:05"
)

type IntTime int64

func (t IntTime) Time() time.Time {
	return time.Unix(int64(t), 0)
}
func (t IntTime) String() string {
	return time.Unix(int64(t), 0).Format(TIME_FORMAT)
}
func (t *IntTime) Set() {
	*t = IntTime(time.Now().Unix())
}
