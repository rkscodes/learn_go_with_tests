package array_sum

func Sum(x []int) (sum int) {
	for _, v := range x {
		sum += v
	}
	return
}

// func SumAll(x ...[]int) (sum []int) {
// 	for i := range x {
// 		sum = append(sum, Sum(x[i]))
// 	}
// 	return
// }

func SumAll(x ...[]int) []int {
	res := make([]int, 0, len(x))
	for i := range x {
		res = append(res, Sum(x[i]))
	}
	return res
}

func SumAllTails(x ...[]int) []int {
	res := make([]int, 0, len(x))
	for _, v := range x {
		if len(v) == 0 {
			res = append(res, 0)
			continue
		}
		temp := v[1:]
		res = append(res, Sum(temp))
	}
	return res
}
