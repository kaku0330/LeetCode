package partitioningintominimumnumberofdecibinarynumbers

func MinPartitions(n string) int {
	var num rune
	for _, val := range n {
		if val > num {
			num = val
		}
	}

	return int(num - 48)
}
