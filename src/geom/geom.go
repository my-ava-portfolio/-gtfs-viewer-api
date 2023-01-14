package geom

func GetBounds(x, y []float32) [4]float32 {
	minX, maxX := FindMaxAndMin(x)
	minY, maxY := FindMaxAndMin(y)
	return [4]float32{minX, minY, maxX, maxY}
}


func FindMaxAndMin(elements []float32) (float32, float32) {
	min, max := elements[0], elements[0]
	for _, element := range elements {
		if element < min {
			min = element
		}
		if element > max {
			max = element
		}
	}
	return min, max
}

func FindMaxAndMinInt(elements []uint32) (uint32, uint32) {
	min, max := elements[0], elements[0]
	for _, element := range elements {
		if element < min {
			min = element
		}
		if element > max {
			max = element
		}
	}
	return min, max
}