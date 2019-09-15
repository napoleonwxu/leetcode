func findOrder(numCourses int, prerequisites [][]int) []int {
    graph := make([][]int, numCourses)
    indegree := make([]int, numCourses)
    for _, p := range prerequisites {
        graph[p[1]] = append(graph[p[1]], p[0])
        indegree[p[0]]++
    }
    course := make([]int, 0, numCourses)
    for i, d := range indegree {
        if d == 0 {
            course = append(course, i)
        }
    }
    for i := 0; i < len(course); i++ {
        for _, j := range graph[course[i]] {
            indegree[j]--
            if indegree[j] == 0 {
                course = append(course, j)
            }
        }
    }
    if len(course) == numCourses {
        return course
    }
    return nil
}