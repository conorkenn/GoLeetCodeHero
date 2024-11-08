package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {

	count := 0
	flowerbed = append(flowerbed, 0)
	t := []int{0}
	flowerbed = append(t, flowerbed...)
	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			count += 1
		}
		if count >= n {
			return true
		}
	}
	return false
}
