package bounds


type minMaxFloat32Model struct {
    Min float32
    Max float32
}

func GetMinmax_float32Array(input []float32) minMaxFloat32Model {

    mins := make(chan float32) // channel for mins
    maxs := make(chan float32) // channel for maxs

    go calculateMinMaxWithChannelFloat32(input[:len(input)/2], mins, maxs)
    go calculateMinMaxWithChannelFloat32(input[len(input)/2:], mins, maxs)

    min1, min2, max1, max2 := <-mins, <-mins, <-maxs, <-maxs //take value from channel

    var minVal float32
    var maxVal float32
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

    return minMaxFloat32Model{Min: minVal, Max: maxVal}
}

func calculateMinMaxWithChannelFloat32(input []float32, mins chan float32, maxs chan float32) {

    var min float32 = input[0]
    var max float32 = input[1]

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