type FrequencyTracker struct {
	cnts map[int]int
	freq map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		cnts: make(map[int]int),
		freq: make(map[int]int),
	}
}

func (this *FrequencyTracker) Add(number int) {
	cur := this.cnts[number]
	this.freq[cur]--
	this.cnts[number]++
	this.freq[cur+1]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if cur := this.cnts[number]; cur > 0 {
		this.freq[cur]--
		this.cnts[number]--
		this.freq[cur-1]++
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	return this.freq[frequency] > 0
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */