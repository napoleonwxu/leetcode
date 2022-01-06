import "math/rand"

type Solution struct {
	m  int
	n  int
	mp map[int]bool
}

func Constructor(m int, n int) Solution {
	return Solution{m: m, n: n, mp: make(map[int]bool)}
}

func (this *Solution) Flip() []int {
	for {
		r := rand.Intn(this.m * this.n)
		if !this.mp[r] {
			this.mp[r] = true
			return []int{r / this.n, r % this.n}
		}
	}
}

func (this *Solution) Reset() {
	this.mp = make(map[int]bool)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */