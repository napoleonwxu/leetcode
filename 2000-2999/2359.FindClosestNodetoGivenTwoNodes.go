func closestMeetingNode(edges []int, node1 int, node2 int) int {
    if node1 == node2 {
        return node1
    }
    node1Steps := calculateSteps(edges, node1)
    node2Steps := calculateSteps(edges, node2)
    ans, minStep := -1, len(edges)
    for node, step1 := range node1Steps {
        if step2, ok := node2Steps[node]; ok {
            step := max(step1, step2)
            if step < minStep || (step == minStep && node < ans) {
                minStep = step
                ans = node
            }
        }
    }
    return ans
}

func calculateSteps(edges []int, node int) map[int]int {
    steps := make(map[int]int, len(edges))
    step := 0
    for node != -1 {
        if _, ok := steps[node]; ok {
            break
        }
        steps[node] = step
        step++
        node = edges[node]
    }
    return steps
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
