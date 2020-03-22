func angleClock(hour int, minutes int) float64 {
    angle := float64(30*hour) + 0.5*float64(minutes) - float64(6*minutes)
    if angle < 0 {
        angle += 360
    }
    if angle > 180 {
        angle = 360 - angle
    }
    return angle
}