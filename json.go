package main

import (
	"fmt"
	"time"
)

type JSONTime time.Time

func (t JSONTime) String() string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Time(t).In(loc).Format("2006-01-02 15:04:05")
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf(`"%s"`, t.String())
	return []byte(stamp), nil
}
