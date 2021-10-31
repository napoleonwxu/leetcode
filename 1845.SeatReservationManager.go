type SeatManager struct {
	reserved []bool
	idx      int
}

func Constructor(n int) SeatManager {
	return SeatManager{reserved: make([]bool, n+1), idx: 1}
}

func (this *SeatManager) Reserve() int {
	seat := this.idx
	this.reserved[this.idx] = true
	for ; this.idx < len(this.reserved) && this.reserved[this.idx]; this.idx++ {
	}
	return seat
}

func (this *SeatManager) Unreserve(seatNumber int) {
	this.reserved[seatNumber] = false
	if seatNumber < this.idx {
		this.idx = seatNumber
	}
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */