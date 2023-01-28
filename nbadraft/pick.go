package nbadraft

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
*
https://www.zhihu.com/question/52895544

1. 30支球队，没进季后赛的14支球队，即乐透球队，有资格抽状元签
2. 一共有1000个签，分配给未能进入季后赛的14支球队，每个队根据战绩倒序分配不同数量的签，分配的数字组合随机，倒数第一有250个签，4分之1就是这里来的，然后抽签的时候看那个数字组合跟谁手里的某一签一样就是谁的，只抽取123顺位，4-14顺位按照没抽到的那11支球队战绩倒序分配。
3. 14个球队签数量分配：250, 199, 156, 119, 88, 63, 43, 28, 17, 11, 8, 7, 6, 5
4. 怎么抽？14个乒乓球分别贴上1-14数字，随机滚出4个，加起来是1001可能，其中11、12、13、14这个组合不算，剩下1000种可能。

*/

/** 抽出四个球的组合结果对应的签编号 */
var ballsResultToSignNum map[string]int

/** 签编号对应的球队ID */
var signNumToTeamId map[int]int

func init() {

	fmt.Println("----------------数据初始化------------------")

	/** 14个球 */
	n := 14
	/** 抽出4个球 */
	m := 4

	nn := 1
	mm := 1
	nm := 1

	/** A14 */
	for i := 1; i <= n; i++ {
		nn = nn * i
	}
	/** A4 */
	for i := 1; i <= m; i++ {
		mm = mm * i
	}
	/** A14-5 */
	for i := 1; i <= (n - m); i++ {
		nm = nm * i
	}

	/** 即 C14-4 = 14 *13 *12 *11 /4 /3 /2 /1 */
	fmt.Printf("排列组合数量：%v\n", nn/mm/nm)

	/** 所有组合结果 */
	allResult := make([]string, 1000)
	ballsResultToSignNum = make(map[string]int)

	count := 0
	for i := 1; i <= 11; i++ {
		for j := i + 1; j <= 12; j++ {
			for k := j + 1; k <= 13; k++ {
				for l := k + 1; l <= 14; l++ {
					//fmt.Printf("%d -> %d -> %d -> %d\n", i, j, k, l)
					if i == 11 && j == 12 && k == 13 && l == 14 {
						continue
					}
					allResult[count] = fmt.Sprintf("%d-%d-%d-%d", i, j, k, l)
					count++
					ballsResultToSignNum[fmt.Sprintf("%d-%d-%d-%d", i, j, k, l)] = count
				}
			}
		}
	}
	fmt.Printf("全部组合结果数量：%v\n", count)
	fmt.Printf("第1000个组合：%v\n", allResult[999])
	fmt.Printf("第1000个组合的签号：%v\n", ballsResultToSignNum[allResult[999]])

	weightArr := [14]int{250, 199, 156, 119, 88, 63, 43, 28, 17, 11, 8, 7, 6, 5}
	allWeight := 0
	signNumToTeamId = make(map[int]int)

	//TODO 随机签号分配给球队
	for num, weight := range weightArr {
		for i := 0; i < weight; i++ {
			signNumToTeamId[i+1+allWeight] = num + 1
		}
		allWeight = allWeight + weight
	}
	fmt.Printf("全部权重总量：%v\n", allWeight)
	fmt.Printf("第1000签号对应的球队ID：%v\n", signNumToTeamId[1000])
	fmt.Printf("第1签号对应的球队ID：%v\n", signNumToTeamId[1])
	fmt.Printf("第250签号对应的球队ID：%v\n", signNumToTeamId[250])
	fmt.Printf("第251签号对应的球队ID：%v\n", signNumToTeamId[251])
	fmt.Printf("第995签号对应的球队ID：%v\n", signNumToTeamId[995])
	fmt.Printf("第996签号对应的球队ID：%v\n", signNumToTeamId[996])

	fmt.Println("----------------数据初始化------------------")

}

var rr = rand.New(rand.NewSource(time.Now().Unix()))

func Pick() string {
	/** 全部球14个 */
	allBalls := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	chosenBalls := make([]int, 4)
	for i := 0; i < 4; i++ {
		hit := rr.Intn(14 - i)
		chosenBalls[i] = allBalls[hit]
		index := hit
		/** 删除选择的球 */
		allBalls = append(allBalls[:index], allBalls[index+1:]...)

	}

	sort.Ints(chosenBalls)
	result := fmt.Sprintf("%d-%d-%d-%d", chosenBalls[0], chosenBalls[1], chosenBalls[2], chosenBalls[3])

	if result == "11-12-13-14" {
		/** 递归重抽 */
		return Pick()
	}

	return result
}

func TeamId(pickResult string) int {
	return signNumToTeamId[ballsResultToSignNum[pickResult]]
}

func PickTeamId() int {
	return TeamId(Pick())
}
