package arrayslice

func Sum(numbers []int) (total int) {
	for _, n := range numbers {
		total += n
	}
	return
}

func SumAll(numbers ...[]int) (sumed []int) {
	for _, n := range numbers {
		sumed = append(sumed, Sum(n))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
