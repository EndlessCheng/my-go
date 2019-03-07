package main

import (
	"os"
	"strconv"
	"time"
	"fmt"
)

var (
	lenUnixTimestamp = len(strconv.Itoa(int(time.Now().Unix())))
	exp10            = [...]int64{
		1,
		1e1,
		1e2,
		1e3,
		1e4,
		1e5,
		1e6,
		1e7,
		1e8,
		1e9,
	}
	local, _ = time.LoadLocation("Asia/Shanghai")
)

func parseTimestamp(rawTs string) (time.Time, error) {
	ts, err := strconv.ParseInt(rawTs, 10, 64)
	if err != nil {
		// 寻找其他解析方式
		if t, err := time.ParseInLocation("2006-01-02 15:04:05", rawTs, local); err == nil {
			return t, nil
		}
		return time.ParseInLocation(time.RFC3339, rawTs, local)
	}

	switch {
	case len(rawTs) >= lenUnixTimestamp:
		_exp10 := exp10[len(rawTs)-lenUnixTimestamp]
		return time.Unix(ts/_exp10, ts%_exp10), nil
	default:
		return time.Unix(ts, 0), nil
	}
}

func main() {
	var t time.Time
	if len(os.Args) <= 1 {
		t = time.Now()
	} else {
		var err error
		t, err = parseTimestamp(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	t.In(local)
	fmt.Println(t.Unix(), t.Format(time.RFC3339))
}
