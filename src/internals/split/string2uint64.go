package split

import "strconv"


func StringToUint64(inputValue string) uint64 {
	value, err := strconv.ParseUint(inputValue, 10, 64)
	if err != nil {
		panic(err)
	}	
	return value
}
