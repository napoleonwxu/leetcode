type TopVotedCandidate struct {
	times []int
	leads []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	cnt := make(map[int]int)
	leads := make([]int, len(persons))
	lead := -1
	for i, p := range persons {
		cnt[p]++
		if i == 0 || cnt[p] >= cnt[lead] {
			lead = p
		}
		leads[i] = lead
	}
	return TopVotedCandidate{times: times, leads: leads}
}

func (this *TopVotedCandidate) Q(t int) int {
	left, right := 0, len(this.times)
	for left < right {
		mid := (left + right) / 2
		if this.times[mid] <= t {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return this.leads[left-1]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */