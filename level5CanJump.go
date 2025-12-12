package piscine

func CanJump(nums []uint) bool {
    if len(nums) == 0 {
        return false // No elements in the array, can't jump to the last index
    }
    
    maxReach := 0 // This will track the furthest position we can reach
    
    for i := 0; i < len(nums); i++ {
        if i > maxReach {
            return false // If we reach an index that is beyond the maximum reachable index, return false
        }
        // Update maxReach to the maximum of current position + steps we can take at current position
        maxReach = max(maxReach, i+int(nums[i]))
        
        // If we can already reach the last index, no need to continue
        if maxReach >= len(nums)-1 {
            return true
        }
    }
    return false
}

// Helper function to get the maximum of two values
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
