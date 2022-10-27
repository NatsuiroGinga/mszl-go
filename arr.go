package main

//
//import (
//	"fmt"
//	"math/rand"
//)
//
//func RandInt64(min, max int64) int64 {
//	if min >= max || min == 0 || max == 0 {
//		return max
//	}
//	return rand.Int63n(max-min) + min
//}
//
//func SplitScores(score int64) []int {
//	var splits []int
//
//	for i := 1; i <= 25; i++ {
//		for j := 1; j <= 25; j++ {
//			for k := 1; k <= 25; k++ {
//				for m := 1; m <= 25; m++ {
//					if int64(i+j+k+m) == score {
//						splits = append(splits, i, j, k, m)
//						return splits
//					}
//				}
//			}
//		}
//	}
//	return nil
//}
//
//func main() {
//	rand.Seed(20011005)
//	var scores []int64
//
//	for i := 0; i < 10; i++ {
//		score := RandInt64(70, 99)
//		scores = append(scores, score)
//		splits := SplitScores(score)
//
//		fmt.Printf("第%d组的总成绩是%d, 四个方面分为%v\n", i+1, score, splits)
//	}
//}
