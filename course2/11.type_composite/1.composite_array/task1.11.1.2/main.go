package main

func sortDescInt(xs [8]int) [8]int {
	var a []int
	b := [8]int{}
	a = append(a, xs[:]...)
	i := 0
	for i != 8 {
		c := a[0]
		for _, x := range a {
			if x > c {
				c = x

			}
		}
		for i, v := range a {
			if v == c {
				a = append(a[:i], a[i+1:]...)
			}
		}
		b[i] = c
		i++
	}
	return b
}
func sortAscInt(xs [8]int) [8]int {
	var a []int
	resSort := sortDescInt(xs)
	a = append(a, resSort[:]...)
	var resAsc [8]int
	for i, v := range a {
		resAsc[7-i] = v
	}
	return resAsc
}
func sortDescFloat(xs [8]float64) [8]float64 {
	var a []float64
	b := [8]float64{}
	a = append(a, xs[:]...)
	i := 0
	for i != 8 {
		c := a[0]
		for _, x := range a {
			if x > c {
				c = x

			}
		}
		for i, v := range a {
			if v == c {
				a = append(a[:i], a[i+1:]...)
			}
		}
		b[i] = c
		i++
	}
	return b
}
func sortAscFloat(xs [8]float64) [8]float64 {
	var a []float64
	resSort := sortDescFloat(xs)
	a = append(a, resSort[:]...)
	var resAsc [8]float64
	for i, v := range a {
		resAsc[7-i] = v
	}
	return resAsc
}
