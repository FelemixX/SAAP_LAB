package lab_5

// FindMinMax - параллельное вычисление максимального и минимального элементов одномерного массива
func FindMinMax(arr []int) (int, int) {
	n := len(arr)
	if n == 0 {
		return 0, 0
	}
	if n == 1 {
		return arr[0], arr[0]
	}

	chMin := make(chan int)
	chMax := make(chan int)

	go func() {
		min := arr[0]
		for _, v := range arr {
			if v < min {
				min = v
			}
		}
		chMin <- int(min)
	}()

	go func() {
		max := arr[0]
		for _, v := range arr {
			if v > max {
				max = v
			}
		}
		chMax <- int(max)
	}()

	return <-chMin, <-chMax
}
