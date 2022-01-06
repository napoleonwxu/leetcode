func findEvenNumbers(digits []int) []int {
	// O(n)
	m := make(map[int]int)
	for _, d := range digits {
		m[d]++
	}
	nums := []int{}
	for num := 100; num < 999; num += 2 {
		cnts := make([]int, 10)
		tmp := num
		for tmp > 0 {
			cnts[tmp%10]++
			tmp /= 10
		}
		flag := true
		for i, cnt := range cnts {
			if cnt > m[i] {
				flag = false
				break
			}
		}
		if flag {
			nums = append(nums, num)
		}
	}
	/* o(n^3)
	   n := len(digits)
	   if n < 3 {
	       return nil
	   }
	   m := make(map[int]bool)
	   for i, a := range digits {
	       if a == 0 {
	           continue
	       }
	       for j, b := range digits {
	           for k, c := range digits {
	               if j != i && k != i && k != j && c&1 == 0 {
	                   m[100*a+10*b+c] = true
	               }
	           }
	       }
	   }
	   nums := []int{}
	   for num := range m {
	       nums = append(nums, num)
	   }
	   sort.Ints(nums)
	*/
	return nums
}