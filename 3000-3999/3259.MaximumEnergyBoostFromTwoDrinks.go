func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
    curA, curB := int64(energyDrinkA[0]), int64(energyDrinkB[0])
    for i := 1; i < len(energyDrinkA); i++ {
        tmpA := curA
        curA = int64(energyDrinkA[i]) + max(curA, curB-int64(energyDrinkB[i-1]))
        curB = int64(energyDrinkB[i]) + max(curB, tmpA-int64(energyDrinkA[i-1]))
    }
    return max(curA, curB)
}