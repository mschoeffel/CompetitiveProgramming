package X2463_MinimumTotalDistanceTraveled

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func Main() {
	robot := []int{0, 4, 6}
	factory := [][]int{{2, 2}, {6, 2}}

	result := minimumTotalDistance(robot, factory)
	fmt.Println("Result Minimum Total Distance Traveled: " + strconv.Itoa(int(result)))
}

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})

	lengthRobot, lengthFactory := len(robot), len(factory)

	dp := make([][]int64, lengthRobot+1)
	for i := range dp {
		dp[i] = make([]int64, lengthFactory+1)
		if i < lengthRobot {
			dp[i][lengthFactory] = math.MaxInt64
		}
	}

	for j := lengthFactory - 1; j >= 0; j-- {
		var prefix int64
		qq := make([][2]int64, 0)
		qq = append(qq, [2]int64{int64(lengthRobot), 0})

		for i := lengthRobot - 1; i >= 0; i-- {
			prefix += int64(abs(robot[i] - factory[j][0]))

			if qq[0][0] > int64(i+factory[j][1]) {
				qq = qq[1:]
			}

			curr := dp[i][j+1] - prefix
			for len(qq) > 0 && qq[len(qq)-1][1] >= curr {
				qq = qq[:len(qq)-1]
			}

			qq = append(qq, [2]int64{int64(i), curr})
			dp[i][j] = qq[0][1] + prefix
		}
	}

	return dp[0][0]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
