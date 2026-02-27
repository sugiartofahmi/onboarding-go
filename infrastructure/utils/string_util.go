package utils

import "strconv"

func StringToInt(value string) int {
    intValue, err := strconv.Atoi(value)
    if err != nil {
        return 0
    }
    return intValue
}