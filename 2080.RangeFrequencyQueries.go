type RangeFreqQuery struct {
	m map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	m := make(map[int][]int)
	for i, a := range arr {
		m[a] = append(m[a], i)
	}
	return RangeFreqQuery{m: m}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	//i := binSearchLeft(this.m[value], left)
	//j := binSearchLeft(this.m[value], right+1)
	i := binSearchRight(this.m[value], left-1)
	j := binSearchRight(this.m[value], right)
	return j - i
}

// return index if target insert into nums, insert into left if same
func binSearchLeft(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// return index if target insert into nums, insert into right if same
func binSearchRight(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */