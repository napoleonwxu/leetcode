func subarrayLCM(nums []int, k int) int {
	cnt, n := 0, len(nums)
	for i := 0; i < n; i++ {
		cur := 1
		for j := i; j < n; j++ {
			cur = lcm(cur, nums[j])
			if cur == k {
				cnt++
			} else if cur > k {
				break
			}
		}
	}
	return cnt
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
