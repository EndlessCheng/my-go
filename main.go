package main

import (
	"os"
	"strconv"
	"time"
	"fmt"
)

var (
	lenUnixTimestamp     = len(strconv.Itoa(int(time.Now().Unix())))
	lenUnixNanoTimestamp = len(strconv.Itoa(int(time.Now().UnixNano())))
)

func main() {
	if len(os.Args) <= 1 {
		os.Exit(1)
	}

	ts, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	var t time.Time
	switch len(os.Args[1]) {
	case lenUnixTimestamp:
		t = time.Unix(int64(ts), 0)
	case lenUnixTimestamp + 3:
		t = time.Unix(int64(ts)/1000, int64(ts)%1000)
	case lenUnixNanoTimestamp:
		t = time.Unix(int64(ts)/int64(time.Second), int64(ts)%int64(time.Second))
	}

	fmt.Println(t.Format("2006-01-02 15:04:05"))
}
