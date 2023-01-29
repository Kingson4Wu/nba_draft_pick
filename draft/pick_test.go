package draft_test

import (
	"fmt"
	"nba_draft_pick/draft"
	"testing"
)

/** 抽一次 */
func TestFirstPick(t *testing.T) {
	fmt.Println("模拟抽签---")

	d := draft.NewDraft([14]int{250, 199, 156, 119, 88, 63, 43, 28, 17, 11, 8, 7, 6, 5})

	firstPick(func() int {
		return d.PickTeamId()
	})

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

}

func firstPick(pickTeamId func() int) {
	//模拟100次，统计每个排名获得状元签的概率
	allHitResult := make([]int, 14)
	for xx := 0; xx < 100000; xx++ {

		allHitResult[pickTeamId()-1]++
	}
	for i, resultHit := range allHitResult {
		fmt.Printf("%d -> %d\n", i+1, resultHit)
	}
}

/** 抽三次 */
func TestThreePick(t *testing.T) {

	d := draft.NewDraft([14]int{250, 199, 156, 119, 88, 63, 43, 28, 17, 11, 8, 7, 6, 5})
	/**
	模拟前三顺位
	*/
	threePick(func() int {
		return d.PickTeamId()
	})
}

func threePick(pickTeamId func() int) {
	/**
	模拟前三顺位
	*/
	allHitResult := make([]int, 14)
	allHitResultSecondRound := make([]int, 14)
	allHitResultThirdRound := make([]int, 14)
	for xx := 0; xx < 1000000; xx++ {

		hitMap := make(map[int]bool)
		for round := 0; round < 3; round++ {

			teamId := pickTeamId()

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

/** 抽一次 */
func TestFirstPick2019(t *testing.T) {
	fmt.Println("2019模拟抽签---")

	d := draft.NewDraft([14]int{140, 140, 140, 125, 105, 90, 60, 60, 60, 30, 20, 10, 10, 10})

	firstPick(func() int {
		return d.PickTeamId()
	})

	/**
	1 -> 13747
	2 -> 13941
	3 -> 13861
	4 -> 12543
	5 -> 10598
	6 -> 9166
	7 -> 6128
	8 -> 5977
	9 -> 6083
	10 -> 2919
	11 -> 2079
	12 -> 925
	13 -> 1004
	14 -> 1029
	*/
}

func TestThreePick2019(t *testing.T) {

	d := draft.NewDraft([14]int{140, 140, 140, 125, 105, 90, 60, 60, 60, 30, 20, 10, 10, 10})

	threePick(func() int {
		return d.PickTeamId()
	})

	/**

	第一轮
	1 -> 139476
	2 -> 140638
	3 -> 140090
	4 -> 125071
	5 -> 104909
	6 -> 90422
	7 -> 59770
	8 -> 60083
	9 -> 60171
	10 -> 29727
	11 -> 19905
	12 -> 9889
	13 -> 9912
	14 -> 9937
	第二轮
	1 -> 134218
	2 -> 133909
	3 -> 133771
	4 -> 121867
	5 -> 105850
	6 -> 92016
	7 -> 63365
	8 -> 63418
	9 -> 63582
	10 -> 32888
	11 -> 21774
	12 -> 11174
	13 -> 10981
	14 -> 11187
	第三轮
	1 -> 127275
	2 -> 127924
	3 -> 126827
	4 -> 119122
	5 -> 104915
	6 -> 93607
	7 -> 67503
	8 -> 67665
	9 -> 67660
	10 -> 35704
	11 -> 24191
	12 -> 12510
	13 -> 12479
	14 -> 12618
	*/
}

func TestNewRoundResult(t *testing.T) {
	d := draft.NewDraft([14]int{250, 199, 156, 119, 88, 63, 43, 28, 17, 11, 8, 7, 6, 5})
	r := d.NewRoundResult()
	fmt.Println(r)

	d2019 := draft.NewDraft([14]int{140, 140, 140, 125, 105, 90, 60, 60, 60, 30, 20, 10, 10, 10})
	r2019 := d2019.NewRoundResult()
	fmt.Println(r2019)
}
