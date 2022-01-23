import (
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	Map := make(map[string][]int)
	for _, order := range orders {
		table, _ := strconv.Atoi(order[1])
		food := order[2]
		if _, ok := Map[food]; !ok {
			Map[food] = make([]int, 501)
		}
		Map[food][table]++
	}
	foods := []string{}
	for food := range Map {
		foods = append(foods, food)
	}
	sort.Strings(foods)
	header := []string{"Table"}
	header = append(header, foods...)
	ans := [][]string{header}
	for table := 1; table <= 500; table++ {
		cnts := make([]int, len(foods))
		sum := 0
		for i, food := range foods {
			cnts[i] += Map[food][table]
			sum += Map[food][table]
		}
		if sum != 0 {
			row := make([]string, len(cnts)+1)
			row[0] = strconv.Itoa(table)
			for i, cnt := range cnts {
				row[i+1] = strconv.Itoa(cnt)
			}
			ans = append(ans, row)
		}
	}
	return ans
}