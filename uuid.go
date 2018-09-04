package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/satori/go.uuid"
)

const GregorianOffset = 122192928000000000

// try   $ uuid `uuid`
func main() {
	if len(os.Args) <= 1 {
		uuid_, _ := uuid.NewV1()
		fmt.Println(uuid_)
		return
	}

	uuid_ := os.Args[1]
	splits := strings.Split(uuid_, "-")
	if len(uuid_) != 36 || len(splits) != 5 {
		fmt.Println("无法识别的字符串")
		return
	}
	if splits[2][0] != '1' {
		fmt.Println("非 UUID v1")
		return
	}
	rawTimestamp := strings.Join([]string{splits[2][1:], splits[1], splits[0]}, "")
	now, _ := strconv.ParseInt(rawTimestamp, 16, 64)
	now = (now - GregorianOffset) * 100
	fmt.Println(time.Unix(now/int64(time.Second), now%int64(time.Second)).Format("2006-01-02 15:04:05"))
}
