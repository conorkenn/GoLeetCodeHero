package leetcode

func convertTemperature(celsius float64) []float64 {
	k := celsius + 273.15
	f := celsius*1.80 + 32.00
	return []float64{k, f}
}
