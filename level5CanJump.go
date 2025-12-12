package main

import (
	"fmt"
)

func CanJump(nums []uint) bool {
    if len(nums) == 0 {
        return false 
    }
    
    maxReach := 0 
    
    for i := 0; i < len(nums); i++ {
        if i > maxReach {
            return false 
        }
        maxReach = max(maxReach, i+int(nums[i]))
        
        if maxReach >= len(nums)-1 {
            return true
        }
    }
    return false
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}