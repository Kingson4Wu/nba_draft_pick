package nbadraft_test

import (
	"fmt"
	"nba_draft_pick/nbadraft"
	"testing"
)

/** 抽一次 */
func TestFirstPick(t *testing.T) {
	fmt.Println("模拟抽签---")

	//模拟100次，统计每个排名获得状元签的概率
	allHitResult := make([]int, 14)
	for xx := 0; xx < 100000; xx++ {

		allHitResult[nbadraft.PickTeamId()-1]++
	}
	for i, resultHit := range allHitResult {
		fmt.Printf("%d -> %d\n", i+1, resultHit)
	}

	/**
		模拟10w次！
	1 -> 25088
	2 -> 19949
	3 -> 15772
	4 -> 11926
	5 -> 8730
	6 -> 6201
	7 -> 4276
	8 -> 2830
	9 -> 1641
	10 -> 1032
	11 -> 802
	12 -> 673
	13 -> 580
	14 -> 500
	*/

	/**
	模拟前三顺位
	*/

}

/** 抽三次 */
func TestThreePick(t *testing.T) {

	/**
	模拟前三顺位
	*/
	allHitResult := make([]int, 14)
	allHitResultSecondRound := make([]int, 14)
	allHitResultThirdRound := make([]int, 14)
	for xx := 0; xx < 1000000; xx++ {

		hitMap := make(map[int]bool)
		for round := 0; round < 3; round++ {

			teamId := nbadraft.PickTeamId()

			/** 这个球队本轮已经抽中过，需要重抽 */
			if _, ok := hitMap[teamId]; ok {
				round--
				continue
			}
			hitMap[teamId] = true

			if round == 0 {
				allHitResult[teamId-1]++
			}
			if round == 1 {
				allHitResultSecondRound[teamId-1]++
			}
			if round == 2 {
				allHitResultThirdRound[teamId-1]++
			}

		}

	}
	fmt.Println("第一轮")
	for i, resultHit := range allHitResult {
		fmt.Printf("%d -> %d\n", i+1, resultHit)
	}
	fmt.Println("第二轮")
	for i, resultHit := range allHitResultSecondRound {
		fmt.Printf("%d -> %d\n", i+1, resultHit)
	}
	fmt.Println("第三轮")
	for i, resultHit := range allHitResultThirdRound {
		fmt.Printf("%d -> %d\n", i+1, resultHit)
	}
}
