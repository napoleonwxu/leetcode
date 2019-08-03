type ExamRoom struct {
    seats []int
    n int
}


func Constructor(N int) ExamRoom {
    return ExamRoom{seats: []int{}, n: N}
}


func (this *ExamRoom) Seat() int {
    m := len(this.seats)
    if m == 0 {
        this.seats = append(this.seats, 0)
        return 0
    }
    dis := max(this.seats[0], this.n-1-this.seats[m-1])
    for i := 1; i < m; i++ {
        dis = max(dis, (this.seats[i]-this.seats[i-1])/2)
    }
    if dis == this.seats[0] {
        this.seats = append([]int{0}, this.seats...)
        return 0
    }
    for i := 1; i < m; i++ {
        if dis == (this.seats[i]-this.seats[i-1])/2 {
            tmp := make([]int, i+1)
            copy(tmp, this.seats[:i])
            tmp[i] = this.seats[i-1] + dis
            this.seats = append(tmp, this.seats[i:]...)
            return this.seats[i]
        }
    }
    this.seats = append(this.seats, this.n-1)
    return this.n-1
}


func (this *ExamRoom) Leave(p int)  {
    for i, seat := range this.seats {
        if seat == p {
            this.seats = append(this.seats[:i], this.seats[i+1:]...)
            return
        }
    }
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(N);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */