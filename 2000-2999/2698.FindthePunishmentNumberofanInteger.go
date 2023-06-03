func punishmentNumber(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		sq := i * i
		if isPunishNumber(sq, i) {
			ans += sq
		}
	}
	return ans
}

func isPunishNumber(num, target int) bool {
	if target < 0 || num < target {
		return false
	}
	if num == target {
		return true
	}
	return isPunishNumber(num/10, target-num%10) || isPunishNumber(num/100, target-num%100) || isPunishNumber(num/1000, target-num%1000)
}