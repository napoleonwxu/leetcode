func wateringPlants(plants []int, capacity int) int {
	n, i := len(plants), -1
	water, steps := 0, 0
	for i < n-1 {
		for i < n-1 && water >= plants[i+1] {
			water -= plants[i+1]
			i++
		}
		if i == n-1 {
			steps += i + 1
		} else {
			steps += 2 * (i + 1)
		}
		water = capacity
	}
	return steps
}