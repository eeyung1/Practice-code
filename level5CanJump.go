package piscine

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
