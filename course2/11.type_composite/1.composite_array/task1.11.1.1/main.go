package main

func sum(xs [8]int) int {
	var sumAll int
	for _, x := range xs {
		sumAll += x
	}
	return sumAll
}
func average(xs [8]int) float64 {

	var sumAll float64
	for _, x := range xs {
		sumAll += float64(x)
	}
	return sumAll / float64(len(xs))
}
func averageFloat(xs [8]float64) float64 {
	var sumAll float64
	for _, x := range xs {
		sumAll += x
	}
	return sumAll / float64(len(xs))
}
func reverse(xs [8]int) [8]int {
	var xsR [8]int
	i := len(xs) - 1
	j := 0
	for i >= 0 {
		xsR[j] = xs[i]
		i--
		j++
	}
	return xsR

}
