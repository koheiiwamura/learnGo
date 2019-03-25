package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	PAA   = 1
	GUU   = 2
	CHOKI = 3
)

func main() {
	COUNT := 100000
	// ランダムなSeedを生成する(生成しないと、何度処理を走らせても同じ結果になるため)
	rand.Seed(time.Now().UnixNano())
	// [引き分け, 勝利, 敗北]
	result := [3]int{}

	for i := 0; i < COUNT; i++ {
		y := yukipiz()
		k := kent()

		switch {
		case y == k:
			result[0]++
		case y == PAA && k == GUU, y == GUU && k == CHOKI, y == CHOKI && k == PAA:
			result[1]++
		case y == PAA && k == CHOKI, y == GUU && k == PAA, y == CHOKI && k == GUU:
			result[2]++
		default:
			panic("Error")
		}
	}

	fmt.Printf("引き分け： %d\n", result[0])
	fmt.Printf("勝利： %d\n", result[1])
	fmt.Printf("敗北： %d\n", result[2])

}

func yukipiz() int {
	s := [3]int{PAA, GUU, CHOKI}
	return s[rand.Intn(len(s))]
}

func kent() int {
	s := []int{PAA, GUU, CHOKI}
	return s[rand.Intn(len(s))]
}
