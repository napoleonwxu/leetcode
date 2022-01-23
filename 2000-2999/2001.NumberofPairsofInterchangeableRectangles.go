func interchangeableRectangles(rectangles [][]int) int64 {
	cnt := int64(0)
	ratio := make(map[float64]int64)
	for _, rect := range rectangles {
		r := float64(rect[0]) / float64(rect[1])
		cnt += ratio[r]
		ratio[r]++
	}
	return cnt
}