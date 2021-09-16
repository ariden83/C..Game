package main

import (
    "fmt" // Use fmt.Println(...) to debug your code
)

func ClosestToZero(slice []float32) float32 {
     fmt.Println(slice)
    if len(slice) > 0 {
        returnValue := slice[0]
        for i := 1; i < len(slice); i++ {
            if IntAbs(slice[i]) < IntAbs(returnValue) {
                returnValue = slice[i]
            } else if IntAbs(slice[i]) == IntAbs(returnValue) && returnValue < slice[i] {
                returnValue = slice[i]
            }
        }
        return returnValue
    } 
    return 0
}

func IntAbs(x float32) float32 {
    if x < 0 {
        return -x
    }
    return x
}
