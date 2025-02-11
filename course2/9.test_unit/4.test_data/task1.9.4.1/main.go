package main

func average(xs []float64) float64 {
	total := 0.0
	i := 0
	for _, x := range xs {
		total += x
		i++
	}
	return total / float64(i)
}
func generateSlice(size int) []float64 {
	var s []float64
	for i := 0; i < size; i++ {
		s = append(s, float64(i))
	}
	return s
}
