func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
    eng, exp := initialEnergy, initialExperience
    hour := 0
    for i := 0; i < len(energy); i++ {
        if eng <= energy[i] {
            hour += energy[i] - eng + 1
            eng += energy[i] - eng + 1
        }
        if exp <= experience[i] {
            hour += experience[i] - exp + 1
            exp += experience[i] - exp + 1
        }
        eng -= energy[i]
        exp += experience[i]
    }
    return hour
}
