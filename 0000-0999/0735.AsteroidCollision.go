func asteroidCollision(asteroids []int) []int {
    stack := make([]int, 0, len(asteroids))
    for _, asteroid := range asteroids {
        needAdd := true
        for len(stack) > 0 && stack[len(stack)-1] > 0 && asteroid < 0 {
            if stack[len(stack)-1] < -asteroid {
                stack = stack[:len(stack)-1]
            } else if stack[len(stack)-1] == -asteroid {
                stack = stack[:len(stack)-1]
                needAdd = false
                break
            } else {
                needAdd = false
                break
            }
        }
        if needAdd {
            stack = append(stack, asteroid)
        }
    }
    return stack
}