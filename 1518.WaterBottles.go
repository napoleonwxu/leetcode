func numWaterBottles(numBottles int, numExchange int) int {
	// O(1)
	return numBottles + (numBottles-1)/(numExchange-1)
	// O(logN)
	cnt := numBottles
	for numBottles >= numExchange {
		cnt += numBottles / numExchange
		numBottles = numBottles/numExchange + numBottles%numExchange
	}
	return cnt
}