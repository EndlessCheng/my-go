package main

import (
	"strings"
	"fmt"
	"os"
)

var mahjong = [...]string{
	"1m", "2m", "3m", "4m", "5m", "6m", "7m", "8m", "9m",
	"1p", "2p", "3p", "4p", "5p", "6p", "7p", "8p", "9p",
	"1s", "2s", "3s", "4s", "5s", "6s", "7s", "8s", "9s",
	"dong", "nan", "xi", "bei",
	"zhong", "bai", "fa",
}

func _convert(single string) int {
	for i, m := range mahjong {
		if m == single {
			return i
		}
	}
	return -1
}

func convert(raw string) []int {
	var result []int

	raw = strings.TrimSpace(raw)
	splits := strings.Split(raw, " ")

	for _, split := range splits {
		if split[0] >= '1' && split[0] <= '9' {
			for i := range split[:len(split)-1] {
				single := split[i:i+1] + split[len(split)-1:]
				r := _convert(single)
				if r == -1 {
					return nil
				}
				result = append(result, r)
			}
		} else {
			r := _convert(split)
			if r == -1 {
				return nil
			}
			result = append(result, r)
		}
	}

	return result
}

var cnt []int

func search(dep int) bool {
	for i := range mahjong {
		if cnt[i] >= 3 { // 刻子
			if dep == 3 { // 4组面子
				return true
			}
			cnt[i] -= 3
			if search(dep+1) {
				return true
			}
			cnt[i] += 3
		}
		for i := 0; i <= 24; i++ { // 一直到 7s
			if i%9 <= 6 && cnt[i] >= 1 && cnt[i+1] >= 1 && cnt[i+2] >= 1 { // 顺子
				if dep == 3 { // 4组面子
					return true
				}
				cnt[i]--
				cnt[i+1]--
				cnt[i+2]--
				if search(dep+1) {
					return true
				}
				cnt[i]++
				cnt[i+1]++
				cnt[i+2]++
			}
		}
	}
	return false
}

// 检查是否和牌
func checkWin() bool {
	for i := range mahjong {
		if cnt[i] >= 2 { // 雀头
			cnt[i] -= 2
			if search(0) {
				return true
			}
			cnt[i] += 2
		}
	}
	return false
}

func main() {
	//if len(os.Args) <= 1 {
	//	fmt.Fprintln(os.Stderr, "参数错误")
	//	os.Exit(1)
	//}
	//raw := os.Args[1]
	mj := convert("11222333789s fa fa")

	var ans []interface{}
	for i := range mahjong {
		cnt = make([]int, 34)
		for _, m := range mj {
			cnt[m]++
		}
		if cnt[i] > 4 {
			fmt.Fprintln(os.Stderr, "超过4张一样的牌！")
			os.Exit(1)
		}
		if cnt[i] == 4 {
			continue
		}

		cnt[i]++
		if checkWin() {
			ans = append(ans, mahjong[i])
		}
		cnt[i]--
	}

	if len(ans) > 0 {
		fmt.Println(ans...)
	} else {
		fmt.Println("尚未听牌")
	}
}
