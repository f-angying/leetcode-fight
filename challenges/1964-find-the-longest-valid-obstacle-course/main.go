package main

import "fmt"

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	n := len(obstacles)
	dp := make([]int, n)
	tails := []int{}

	for i := 0; i < n; i++ {
		idx := search(tails, obstacles[i])
		if idx == len(tails) {
			tails = append(tails, obstacles[i])
		} else {
			tails[idx] = obstacles[i]
		}
		dp[i] = idx + 1
	}

	return dp
}

func search(tails []int, target int) int {
	left, right := 0, len(tails)
	for left < right {
		mid := left + (right-left)/2
		if tails[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	obstacles := []int{1, 2, 3, 2}
	rs := longestObstacleCourseAtEachPosition(obstacles)
	fmt.Println("Output: ", rs)

	obstacles2 := []int{2, 2, 1}
	rs = longestObstacleCourseAtEachPosition(obstacles2)
	fmt.Println("Output: ", rs)

	obstacles3 := []int{3, 1, 5, 6, 4, 2}
	rs = longestObstacleCourseAtEachPosition(obstacles3)
	fmt.Println("Output: ", rs)
}
