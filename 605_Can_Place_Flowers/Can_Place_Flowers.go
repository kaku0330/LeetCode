package canplaceflowers

func CanPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if i == 0 {
			if len(flowerbed) == 1 && flowerbed[i] == 0 {
				return true
			}
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		} else if i == len(flowerbed)-1 {
			if flowerbed[i] == 0 && flowerbed[i-1] == 0 {
				flowerbed[i-1] = 1
				n--
			}
		} else {
			if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		}
	}

	if n > 0 {
		return false
	} else {
		return true
	}
}
