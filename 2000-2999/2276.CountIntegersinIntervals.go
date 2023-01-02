type CountIntervals struct {
	intervals [][]int
	cnt       int
}

func Constructor() CountIntervals {
	return CountIntervals{intervals: [][]int{}, cnt: 0}
}

func (this *CountIntervals) Add(left int, right int) {
	l := binSearchLeft(this.intervals, left, 1)
	r := binSearchRight(this.intervals, right, 0)
	n := len(this.intervals)
	if l < n {
		left = min(left, this.intervals[l][0])
	}
	if r > 0 {
		right = max(right, this.intervals[r-1][1])
	}
	for i := l; i < r; i++ {
		this.cnt -= this.intervals[i][1] - this.intervals[i][0] + 1
	}
	this.cnt += right - left + 1
	if l >= n {
		this.intervals = append(this.intervals, []int{left, right})
	} else {
		this.intervals = append(this.intervals[:l+1], this.intervals[r:]...)
		this.intervals[l] = []int{left, right}
	}
	/* TLE
	   tmp := make([][]int, 0, n-(r-l)+1)
	   tmp = append(tmp, this.intervals[:l]...)
	   tmp = append(tmp, []int{left, right})
	   tmp = append(tmp, this.intervals[r:]...)
	   this.intervals = tmp
	*/
}

func (this *CountIntervals) Count() int {
	return this.cnt
}

func binSearchLeft(intervals [][]int, target, idx int) int {
	left, right := 0, len(intervals)
	for left < right {
		mid := (left + right) / 2
		if intervals[mid][idx] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func binSearchRight(intervals [][]int, target, idx int) int {
	left, right := 0, len(intervals)
	for left < right {
		mid := (left + right) / 2
		if intervals[mid][idx] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/**
 * Your CountIntervals object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(left,right);
 * param_2 := obj.Count();
 */