package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GetSmallestInteger(slice []int) int {
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
	} else {
		cgreader.Trace("passed slice has invalid size...\n")
	}
	return 0
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
