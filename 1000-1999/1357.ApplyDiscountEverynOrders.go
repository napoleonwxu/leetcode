type Cashier struct {
    n, c, discount int
    price map[int]int
}


func Constructor(n int, discount int, products []int, prices []int) Cashier {
    Map := make(map[int]int, len(products))
    for i, p := range products {
        Map[p] = prices[i]
    }
    return Cashier{n: n, c: 0, discount: discount, price: Map}
}


func (this *Cashier) GetBill(product []int, amount []int) float64 {
    sum := 0
    for i, p := range product {
        sum += amount[i] * this.price[p]
    }
    this.c++
    if this.c == this.n {
        this.c = 0
        return float64((100-this.discount)*sum)/100
    }
    return float64(sum)
}


/**
 * Your Cashier object will be instantiated and called as such:
 * obj := Constructor(n, discount, products, prices);
 * param_1 := obj.GetBill(product,amount);
 */