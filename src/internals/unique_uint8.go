package internals


func GetUniquesUint8(values []uint8) []uint8 {
	uniques := make([]uint8, 0, len(values))
	m := make(map[uint8]bool)

	for _, val := range values {
		if _, ok := m[val]; !ok {
			m[val] = true
			uniques = append(uniques, val)
		}
	}

	return uniques
}
