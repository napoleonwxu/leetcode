func maxBottlesDrunk(numBottles int, numExchange int) int {
    drunk, empty := 0, 0
    for numBottles > 0 {
        drunk += numBottles
        empty += numBottles
        numBottles = 0
        for empty >= numExchange {
            empty -= numExchange
            numBottles++
            numExchange++
        }
    }
    return drunk
}