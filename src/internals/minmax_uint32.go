package internals


type minMaxUint32Model struct {
    Min uint32
    Max uint32
}

func GetMinmax_uint32Array(input []uint32) minMaxUint32Model {

    mins := make(chan uint32) // channel for mins
    maxs := make(chan uint32) // channel for maxs

    go calculateMinMaxWithChannelUint32(input[:len(input)/2], mins, maxs)
    go calculateMinMaxWithChannelUint32(input[len(input)/2:], mins, maxs)

    min1, min2, max1, max2 := <-mins, <-mins, <-maxs, <-maxs //take value from channel

    var minVal uint32
    var maxVal uint32
    if min1 <= min2 {
        minVal = min1
    } else {
        minVal = min2
    }

    if max1 >= max2 {
        maxVal = max1
    } else {
        maxVal = max2
    }

    return minMaxUint32Model{Min: minVal, Max: maxVal}
}

func calculateMinMaxWithChannelUint32(input []uint32, mins chan uint32, maxs chan uint32) {

    var min uint32 = input[0]
    var max uint32 = input[1]

    if min > max {
        min = input[1]
        max = input[0]
    }

    for i := 0; i < len(input); i++ {
        if input[i] < min {
            min = input[i]
        }
        if input[i] > max {
            max = input[i]
        }
    }

    mins <- min // add mins to mins channel
    maxs <- max // add maxs to max channel
}